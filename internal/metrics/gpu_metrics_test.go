package gpumetrics

import (
	"testing"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/assert"
)

func TestGetuUID(t *testing.T) {
	mockDevice := new(MockNvmlMetricsManager)

	// Set up the mock to return specific values
	mockDevice.On("GetUUID").Return("mock-uuid", nvml.SUCCESS)

	// Call the function under test
	uuid, ret := mockDevice.GetUUID()

	// Assert the results
	assert.Equal(t, "mock-uuid", uuid)
	assert.Equal(t, nvml.SUCCESS, ret)

	// Verify that the method was called
	mockDevice.AssertExpectations(t)
}

func TestFetchDeviceMetrics(t *testing.T) {
	mockDevice := new(MockNvmlMetricsManager)

	mockDevice.On("GetUUID").Return("mock-uuid", nvml.SUCCESS)
	mockDevice.On("GetName").Return("mock-name", nvml.SUCCESS)
	mockDevice.On("GetTemperature", nvml.TEMPERATURE_GPU).Return(uint32(70), nvml.SUCCESS)
	mockDevice.On("GetPowerUsage").Return(uint32(150), nvml.SUCCESS)
	mockDevice.On("GetMemoryInfo").Return(nvml.Memory{Total: 8192, Free: 4096, Used: 4096}, nvml.SUCCESS)
	mockDevice.On("GetUtilizationRates").Return(nvml.Utilization{Gpu: 50, Memory: 30}, nvml.SUCCESS)

	device, ret := FetchDeviceMetrics(mockDevice)
	assert.Equal(t, nvml.SUCCESS, ret)
	assert.NotNil(t, device)
	assert.Equal(t, "mock-uuid", device.UUID)
	assert.Equal(t, "mock-name", device.Name)
	assert.Equal(t, uint32(70), device.Temperature)
	assert.Equal(t, uint32(150), device.Power)
	assert.Equal(t, uint64(8192), device.MemoryTotal)
	assert.Equal(t, uint64(4096), device.MemoryFree)
	assert.Equal(t, uint64(4096), device.MemoryUsed)
	assert.Equal(t, uint32(50), device.UtilizationGpu)
	assert.Equal(t, uint32(30), device.UtilizationMemory)

	mockDevice.AssertExpectations(t)
}
