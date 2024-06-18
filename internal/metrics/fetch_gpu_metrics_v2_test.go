package gpumetrics

import (
	"context"
	"testing"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/assert"
)

func TestFetchAllGpuInfoV2(t *testing.T) {
	mockDeviceManager := new(MockGpuDeviceManager)

	// Happy path test
	t.Run("HappyPath", func(t *testing.T) {
		mockDeviceManager.On("DeviceGetHandleByIndex").Return(&MockNvidiaDevice{}, nvml.SUCCESS)
		mockDeviceManager.On("GetUUID").Return("mock-uuid", nvml.SUCCESS)
		mockDeviceManager.On("GetName").Return("mock-name", nvml.SUCCESS)
		mockDeviceManager.On("GetTemperature").Return(42, nvml.SUCCESS)
		mockDeviceManager.On("GetPowerUsage").Return(100, nvml.SUCCESS)
		mockDeviceManager.On("GetMemoryInfo").Return(nvml.Memory{Total: 1024}, nvml.SUCCESS)

		gpuCount := 1
		nvml.Init()
		defer nvml.Shutdown()
		_, err := FetchAllGpuInfo(context.Background(), mockDeviceManager, gpuCount)
		assert.NoError(t, err)
		// assert.Len(t, info, gpuCount)
	})

	// Error handling test
	t.Run("ErrorHandling", func(t *testing.T) {
		mockDeviceManager.On("DeviceGetHandleByIndex").Return(nil, nvml.ERROR_UNKNOWN)
		gpuCount := 1
		nvml.Init()
		defer nvml.Shutdown()
		info, err := FetchAllGpuInfo(context.Background(), mockDeviceManager, gpuCount)
		assert.Error(t, err)
		assert.Len(t, info, 0)
	})

	// Edge case test: empty list of GPUs
	t.Run("EmptyList", func(t *testing.T) {
		mockDeviceManager.On("DeviceGetHandleByIndex").Return(nil, nvml.ERROR_UNKNOWN)
		gpuCount := 0
		nvml.Init()
		defer nvml.Shutdown()
		info, err := FetchAllGpuInfo(context.Background(), mockDeviceManager, gpuCount)
		assert.NoError(t, err)
		assert.Len(t, info, 0)
	})
}
