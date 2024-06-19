package gpumetrics

import (
	"testing"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/assert"
)

func TestFetchDeviceMetrics(t *testing.T) {
	m := &MockNvmlMetricsManager{}
	deviceUUID := "device-uuid"
	deviceName := "device-name"
	deviceTemp := uint32(70)
	devicePower := uint32(100)
	deviceMemoryTotal := uint64(1024)
	deviceMemoryFree := uint64(512)
	deviceMemoryUsed := uint64(256)
	deviceUtilizationGpu := uint32(50)
	deviceUtilizationMemory := uint32(20)

	m.On("GetUUID").Return(deviceUUID, nvml.SUCCESS)
	m.On("GetName").Return(deviceName, nvml.SUCCESS)
	m.On("GetTemperature", nvml.TEMPERATURE_GPU).Return(deviceTemp, nvml.SUCCESS)
	m.On("GetPowerUsage").Return(devicePower, nvml.SUCCESS)
	m.On("GetMemoryInfo").Return(nvml.Memory{Total: deviceMemoryTotal, Free: deviceMemoryFree, Used: deviceMemoryUsed}, nvml.SUCCESS)
	m.On("GetUtilizationRates").Return(nvml.Utilization{Gpu: deviceUtilizationGpu, Memory: deviceUtilizationMemory}, nvml.SUCCESS)

	result, err := FetchDeviceMetrics(m)
	assert.Equal(t, err, nvml.SUCCESS)
	assert.Equal(t, deviceUUID, result.UUID)
	assert.Equal(t, deviceName, result.Name)
	assert.Equal(t, deviceTemp, result.Temperature)
	assert.Equal(t, devicePower, result.Power)
	assert.Equal(t, deviceMemoryTotal, result.MemoryTotal)
	assert.Equal(t, deviceMemoryFree, result.MemoryFree)
	assert.Equal(t, deviceMemoryUsed, result.MemoryUsed)
	assert.Equal(t, deviceUtilizationGpu, result.UtilizationGpu)
	assert.Equal(t, deviceUtilizationMemory, result.UtilizationMemory)
}

func TestFetchDeviceMetrics_Error(t *testing.T) {
	m := &MockNvmlMetricsManager{}
	m.On("GetUUID").Return("", nvml.ERROR_INVALID_ARGUMENT)
	m.On("GetTemperature", nvml.TEMPERATURE_GPU).Return("", nvml.ERROR_INVALID_ARGUMENT)
	m.On("GetName").Return("", nvml.ERROR_INVALID_ARGUMENT)
	m.On("GetPowerUsage").Return("", nvml.ERROR_INVALID_ARGUMENT)
	m.On("GetMemoryInfo").Return(nvml.Memory{}, nvml.ERROR_INVALID_ARGUMENT)
	m.On("GetUtilizationRates").Return(nvml.Utilization{}, nvml.ERROR_INVALID_ARGUMENT)

	result, err := FetchDeviceMetrics(m)
	assert.Error(t, err)
	assert.Nil(t, result)

}
