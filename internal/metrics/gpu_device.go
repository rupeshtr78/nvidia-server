package gpumetrics

import (
	"log"
	"sync"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// GpuDeviceManager is an interface for managing GPU devices
type GpuDeviceManager interface {
	Init() nvml.Return
	Shutdown() nvml.Return
	GetDevice() nvml.Device
	DeviceGetCount() (int, nvml.Return)
	DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return)
}

var once sync.Once

// NvidiaDevice is a struct to hold the metrics for an Nvidia GPU device
type NvidiaDevice struct {
	Device            nvml.Device
	Name              string
	UUID              string
	Temperature       uint32
	Power             uint32
	MemoryTotal       uint64
	MemoryFree        uint64
	MemoryUsed        uint64
	UtilizationGpu    uint32
	UtilizationMemory uint32
}

// NvmlMetricsManager is an interface for managing NVML metrics
type NvmlMetricsManager interface {
	GetUUID() (string, nvml.Return)
	GetName() (string, nvml.Return)
	GetTemperature(nvml.TemperatureSensors) (uint32, nvml.Return)
	GetPowerUsage() (uint32, nvml.Return)
	GetMemoryInfo() (nvml.Memory, nvml.Return)
	GetUtilizationRates() (nvml.Utilization, nvml.Return)
}

// GpuMap is a map of GPU index to GpuMetrics
type GpuMap map[int]*NvidiaDevice

// NewGpuMap creates a new GpuMap
func NewGpuMap() GpuMap {
	return make(GpuMap)
}

// verify MeticManager implements the nvml.Device interface
var _ GpuDeviceManager = &NvidiaDevice{}

func NewNvidiaDevice() *NvidiaDevice {
	return &NvidiaDevice{}
}

// Init initializes the NVML library
func (m *NvidiaDevice) Init() nvml.Return {
	var intErr = nvml.SUCCESS
	once.Do(func() {
		ret := nvml.Init()
		if ret != nvml.SUCCESS {
			intErr = ret
			return
		}
	})
	return intErr
}

// Shutdown shuts down the NVML library
func (m *NvidiaDevice) Shutdown() nvml.Return {
	var intErr = nvml.SUCCESS
	once.Do(func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			intErr = ret
			return
		}

	})
	return intErr
}

// DeviceGetCount returns the number of GPU devices
func (m *NvidiaDevice) DeviceGetCount() (int, nvml.Return) {
	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		log.Default().Printf("Failed to get device count: %v\n", ret)
	}
	return count, ret
}

// DeviceGetHandleByIndex returns the device handle for a given index
func (m *NvidiaDevice) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	device, ret := nvml.DeviceGetHandleByIndex(index)
	if ret != nvml.SUCCESS {
		log.Default().Printf("Failed to get device handle by index: %v\n", ret)
	}
	return device, ret
}

// GetDevice returns the nvml.Device
func (m *NvidiaDevice) GetDevice() nvml.Device {
	return m.Device
}

// func (m *NvidiaDevice) GetUUID() (string, nvml.Return) {
// 	ret := nvml.SUCCESS
// 	uuid, ret := m.Device.GetUUID()
// 	if ret != nvml.SUCCESS {
// 		log.Default().Printf("Failed to get UUID: %v\n", ret)
// 	}
// 	return uuid, ret
// }

// func (m *NvidiaDevice) GetName() (string, nvml.Return) {
// 	name, ret := m.Device.GetName()
// 	if ret != nvml.SUCCESS {
// 		log.Default().Printf("Failed to get name: %v\n", ret)
// 	}
// 	return name, ret
// }

// func (m *NvidiaDevice) GetTemperature(sensor nvml.TemperatureSensors) (uint32, nvml.Return) {
// 	temp, ret := m.Device.GetTemperature(sensor)
// 	if ret != nvml.SUCCESS {
// 		log.Default().Printf("Failed to get temperature: %v\n", ret)
// 		return 0, ret
// 	}
// 	return temp, ret
// }

// func (m *NvidiaDevice) GetPowerUsage() (uint32, nvml.Return) {
// 	power, ret := m.Device.GetPowerUsage()
// 	if ret != nvml.SUCCESS {
// 		log.Default().Printf("Failed to get fan speed: %v\n", ret)
// 		return 0, ret
// 	}
// 	return power, ret
// }

// func (m *NvidiaDevice) GetMemoryInfo() (nvml.Memory, nvml.Return) {
// 	memory, ret := m.Device.GetMemoryInfo()
// 	if ret != nvml.SUCCESS {
// 		log.Default().Printf("Failed to get memory info: %v\n", ret)
// 		return nvml.Memory{}, ret

// 	}
// 	return memory, ret
// }

// func (m *NvidiaDevice) GetUtilizationRates() (nvml.Utilization, nvml.Return) {
// 	utilization, ret := m.Device.GetUtilizationRates()
// 	if ret != nvml.SUCCESS {
// 		log.Default().Printf("Failed to get utilization rates: %v\n", ret)
// 		return nvml.Utilization{}, ret
// 	}
// 	return utilization, ret
// }
