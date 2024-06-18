package gpumetrics

import (
	"context"
	"testing"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchAllGpuInfoV2(t *testing.T) {
	mockDeviceManager := new(MockGpuDeviceManager)

	// Happy path test
	t.Run("HappyPath", func(t *testing.T) {
		mockDeviceManager.On("DeviceGetHandleByIndex").Return(&MockNvmlDevice{}, nvml.SUCCESS)
		mockDeviceManager.On("GetUUID").Return("mock-uuid", nvml.SUCCESS)
		mockDeviceManager.On("GetName").Return("mock-name", nvml.SUCCESS)
		mockDeviceManager.On("GetTemperature").Return(42, nvml.SUCCESS)
		mockDeviceManager.On("GetPowerUsage").Return(100, nvml.SUCCESS)
		mockDeviceManager.On("GetMemoryInfo").Return(nvml.Memory{Total: 1024}, nvml.SUCCESS)

		gpuCount := 1
		info, err := FetchAllGpuInfo(context.Background(), mockDeviceManager, gpuCount)
		assert.NoError(t, err)
		assert.Len(t, info, gpuCount)
	})

	// Error handling test
	t.Run("ErrorHandling", func(t *testing.T) {
		mockDeviceManager.On("DeviceGetHandleByIndex").Return(nil, nvml.ERROR_UNKNOWN)
		gpuCount := 1
		info, err := FetchAllGpuInfo(context.Background(), mockDeviceManager, gpuCount)
		assert.Error(t, err)
		assert.Len(t, info, 0)
	})

	// Edge case test: empty list of GPUs
	t.Run("EmptyList", func(t *testing.T) {
		mockDeviceManager.On("DeviceGetHandleByIndex").Return(nil, nvml.ERROR_UNKNOWN)
		gpuCount := 0
		info, err := FetchAllGpuInfo(context.Background(), mockDeviceManager, gpuCount)
		assert.NoError(t, err)
		assert.Len(t, info, 0)
	})
}

type MockGpuDeviceManager struct {
	mock.Mock
	// Mocked method return values
	DeviceGetCountReturn         int
	DeviceErr                    nvml.Return
	DeviceGetHandleByIndexReturn nvml.Device
	DeviceGetHandleByIndexErr    nvml.Return
	DeviceReturn                 nvml.Device
	// ... add more mocked method return values as needed

	// Call counters for testing purposes
	DeviceGetCountCallCount         int
	DeviceGetHandleByIndexCallCount int
	// ... add more call counters as needed
}

func (m *MockGpuDeviceManager) DeviceGetCount() (int, nvml.Return) {
	m.DeviceGetCountCallCount++
	return m.DeviceGetCountReturn, m.DeviceErr
}

func (m *MockGpuDeviceManager) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	m.DeviceGetHandleByIndexCallCount++
	return m.DeviceGetHandleByIndexReturn, m.DeviceGetHandleByIndexErr
}

func (m *MockGpuDeviceManager) GetDevice() nvml.Device {
	return m.DeviceReturn
}

func (m *MockGpuDeviceManager) Init() nvml.Return {
	return nvml.SUCCESS
}

func (m *MockGpuDeviceManager) Shutdown() nvml.Return {
	return nvml.SUCCESS
}

func (m *MockGpuDeviceManager) Reset() {
	// Reset call counters and return values
	m.DeviceGetCountCallCount = 0
	m.DeviceGetHandleByIndexCallCount = 0
	m.DeviceGetCountReturn = 0
	m.DeviceErr = nvml.SUCCESS
	m.DeviceGetHandleByIndexReturn = nil
	m.DeviceGetHandleByIndexErr = nvml.SUCCESS
}

type MockNvmlDevice struct {
	mock.Mock
}

func (m *MockNvmlDevice) DeviceGetCount() (int, nvml.Return) {
	args := m.Called()
	return args.Get(0).(int), args.Get(1).(nvml.Return)
}

func (m *MockNvmlDevice) DeviceGetHandleByIndex(index int) (*nvml.Device, nvml.Return) {
	args := m.Called(index)
	return args.Get(0).(*nvml.Device), args.Get(1).(nvml.Return)
}

func (m *MockNvmlDevice) FetchDeviceMetrics(device *nvml.Device) (*NvidiaDevice, nvml.Return) {
	args := m.Called(device)
	return args.Get(0).(*NvidiaDevice), args.Get(1).(nvml.Return)
}
