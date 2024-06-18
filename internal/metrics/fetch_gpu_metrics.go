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

	gpuResChan := make(chan *GpuResult, count) // Buffered channel for results
	defer close(gpuResChan)

	errChan := make(chan error, 1) // only need to store one error
	defer close(errChan)

	wg := new(sync.WaitGroup)
	wg.Add(count)

	for i := 0; i < count; i++ {

		index := i
		errGroup.Go(func() error {
			defer wg.Done()

			res, err := fetchGpuMetrics(gpu, index)
			if err != nil {
				errChan <- err
			}

			if ctx.Err() != nil {
				errChan <- ctx.Err()
			}

			mu.Lock()
			defer mu.Unlock()

			gpuResChan <- res

			return nil
		})
	}

	wg.Wait()

	// main goroutine reads from gpuResChan and updates the gpuMap
	for res := range gpuResChan {
		gpuMap[res.index] = res.gpu
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

	if len(errChan) > 0 {
		return nil, <-errChan
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
