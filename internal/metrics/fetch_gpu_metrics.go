package gpumetrics

import (
	"context"
	"fmt"
	"sync"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"golang.org/x/sync/errgroup"
)

type GpuResult struct {
	index int
	gpu   *NvidiaDevice
}

// FetchGpuInfo fetches the metrics for all GPU devices
func FetchAllGpuInfo(ctx context.Context, gpu GpuDeviceManager, count int) (GpuMap, error) {

	if count == 0 {
		return nil, fmt.Errorf("no GPU devices found")
	}

	var mu sync.Mutex
	errGroup, ctx := errgroup.WithContext(ctx)
	gpuMap := make(GpuMap)

	for i := 0; i < count; i++ {

		index := i
		errGroup.Go(func() error {
			res, err := fetchGpuMetrics(gpu, index)
			if err != nil {
				return err
			}

			if ctx.Err() != nil {
				return ctx.Err()
			}

			mu.Lock()
			defer mu.Unlock()

			gpuMap[res.index] = res.gpu

			return nil
		})
	}

	if err := errGroup.Wait(); err != nil {
		return nil, err
	}

	if len(gpuMap) == 0 {
		return nil, fmt.Errorf("no GPU devices found")
	}

	if len(gpuMap) != count {
		return nil, fmt.Errorf("expected %d GPU devices, got %d", count, len(gpuMap))
	}

	return gpuMap, nil
}

func fetchGpuMetrics(gpuDeviceManager GpuDeviceManager, index int) (*GpuResult, error) {
	device, ret := gpuDeviceManager.DeviceGetHandleByIndex(index)
	if ret != nvml.SUCCESS {
		return nil, fmt.Errorf("failed to get device handle: %v", ret)
	}

	g, err := FetchDeviceMetrics(device)
	if err != nvml.SUCCESS {
		return nil, err
	}

	deviceInfo := &GpuResult{
		index: index,
		gpu:   g,
	}

	return deviceInfo, nil
}
