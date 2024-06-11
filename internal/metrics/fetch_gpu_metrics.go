package gpumetrics

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// FetchGpuInfo fetches the metrics for all GPU devices
func FetchAllGpuInfo(ctx context.Context, gpu GpuDeviceManager, count int) (GpuMap, error) {

	if count == 0 {
		return nil, fmt.Errorf("no GPU devices found")
	}

	gpuMap := make(GpuMap)

	// gpuChan := make(chan GpuMap, count)
	// defer close(gpuChan)

	errChan := make(chan error, 1) // only need to store one error
	defer close(errChan)

	wg := new(sync.WaitGroup)
	wg.Add(count)

	for i := 0; i < count; i++ {
		go func(i int) {
			defer wg.Done()

			// context done means there has been a cancellation signal
			select {
			case <-ctx.Done():
				errChan <- ctx.Err()
				return
			default:
			}

			device, ret := gpu.DeviceGetHandleByIndex(i)
			if ret != nvml.SUCCESS {
				errChan <- fmt.Errorf("failed to get device handle: %v", ret)
				return
			}

			g, err := FetchDeviceMetrics(device)
			if err != nvml.SUCCESS {
				errChan <- fmt.Errorf("failed to get device info: %v", err)
				return
			}

			gpuMap[i] = g
		}(i)
	}

	// main goroutine waits for all goroutines to finish
	wg.Wait()

	for i := 0; i < count; i++ {
		select {
		case <-ctx.Done(): // if context is done, ( cancelled or timeout ) return ctx.err
			return nil, ctx.Err()
		case err := <-errChan: // if there is an error, return err
			return nil, err
		default:
			log.Printf("Fetched Metrics for GPU %d\n", i)
		}
	}

	return gpuMap, nil
}
