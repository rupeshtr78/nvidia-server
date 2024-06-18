package gpumetrics

import (
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/mock"
)

type MockNvidiaDevice struct {
	device            nvml.Device
	name              string
	uuid              string
	temperature       uint32
	power             uint32
	memoryTotal       uint64
	memoryFree        uint64
	memoryUsed        uint64
	utilizationGpu    uint32
	utilizationMemory uint32
	mock.Mock
}

// Mocked methods for MockNvidiaDevice
func (m *MockNvidiaDevice) Device() nvml.Device {
	return m.device
}

func (m *MockNvidiaDevice) Name() string {
	return m.name
}

func (m *MockNvidiaDevice) UUID() string {
	return m.uuid
}

func (m *MockNvidiaDevice) Temperature() uint32 {
	return m.temperature
}

func (m *MockNvidiaDevice) Power() uint32 {
	return m.power
}

func (m *MockNvidiaDevice) MemoryTotal() uint64 {
	return m.memoryTotal
}

func (m *MockNvidiaDevice) MemoryFree() uint64 {
	return m.memoryFree
}

func (m *MockNvidiaDevice) MemoryUsed() uint64 {
	return m.memoryUsed
}

func (m *MockNvidiaDevice) UtilizationGpu() uint32 {
	return m.utilizationGpu
}

func (m *MockNvidiaDevice) UtilizationMemory() uint32 {
	return m.utilizationMemory
}

type MockGpuDeviceManager struct {
	initError                nvml.Return
	shutdownError            nvml.Return
	device                   nvml.Device
	deviceCount              int
	deviceCountError         nvml.Return
	deviceHandleByIndexError nvml.Return
	mock.Mock
}

func (m *MockGpuDeviceManager) Init() nvml.Return {
	return m.initError
}

func (m *MockGpuDeviceManager) Shutdown() nvml.Return {
	return m.shutdownError
}

func (m *MockGpuDeviceManager) GetDevice() nvml.Device {
	return m.device
}

func (m *MockGpuDeviceManager) DeviceGetCount() (int, nvml.Return) {
	return m.deviceCount, m.deviceCountError
}

func (m *MockGpuDeviceManager) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	return m.device, m.deviceHandleByIndexError
}

// MockNvmlMetricsManager is a mock of NvmlMetricsManager interface
type MockNvmlMetricsManagerV2 struct {
	mock.Mock
}

func (m *MockNvmlMetricsManagerV2) GetUUID() (string, nvml.Return) {
	args := m.Called()
	return args.String(0), args.Get(1).(nvml.Return)
}

func (m *MockNvmlMetricsManagerV2) GetName() (string, nvml.Return) {
	args := m.Called()
	return args.String(0), args.Get(1).(nvml.Return)
}

func (m *MockNvmlMetricsManagerV2) GetTemperature(sensors nvml.TemperatureSensors) (uint32, nvml.Return) {
	args := m.Called(sensors)
	return uint32(args.Int(0)), args.Get(1).(nvml.Return)
}

func (m *MockNvmlMetricsManagerV2) GetPowerUsage() (uint32, nvml.Return) {
	args := m.Called()
	return uint32(args.Int(0)), args.Get(1).(nvml.Return)
}

func (m *MockNvmlMetricsManagerV2) GetMemoryInfo() (nvml.Memory, nvml.Return) {
	args := m.Called()
	return args.Get(0).(nvml.Memory), args.Get(1).(nvml.Return)
}

func (m *MockNvmlMetricsManagerV2) GetUtilizationRates() (nvml.Utilization, nvml.Return) {
	args := m.Called()
	return args.Get(0).(nvml.Utilization), args.Get(1).(nvml.Return)
}
