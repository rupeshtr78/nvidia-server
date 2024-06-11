package gpumetrics

import (
	"log"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// FetchDeviceMetrics fetches the metrics for a given device nvml.Device implements NvmlMetricsManager
func FetchDeviceMetrics(device NvmlMetricsManager) (*NvidiaDevice, nvml.Return) {

	// Get the device UUID
	uuid, err := device.GetUUID()
	if err != nvml.SUCCESS {
		log.Default().Printf("Failed to get device UUID: %v\n", err)
		return nil, err
	}

	// Get the device name and UUID
	name, err := device.GetName()
	if err != nvml.SUCCESS {
		log.Default().Printf("Failed to get device name: %v\n", err)
		return nil, err
	}

	// Get the device temperature
	temp, err := device.GetTemperature(nvml.TEMPERATURE_GPU)
	if err != nvml.SUCCESS {
		log.Default().Printf("Failed to get device temperature: %v\n", err)
		return nil, err
	}

	power, err := device.GetPowerUsage()
	if err != nvml.SUCCESS {
		log.Default().Printf("Failed to get device power usage: %v\n", err)
		return nil, err
	}

	// Get the device memory info Total Free Used  uint64
	memory, err := device.GetMemoryInfo()
	if err != nvml.SUCCESS {
		log.Default().Printf("Failed to get device memory info: %v\n", err)
		return nil, err
	}

	// Get the device utilization Gpu and Memory uint
	utilization, err := device.GetUtilizationRates()
	if err != nvml.SUCCESS {
		log.Default().Printf("Failed to get device utilization: %v\n", err)
		return nil, err
	}

	g := NewNvidiaDevice()

	g.Name = name
	g.UUID = uuid
	g.Temperature = temp
	g.Power = power
	g.MemoryTotal = memory.Total
	g.MemoryFree = memory.Free
	g.MemoryUsed = memory.Used
	g.UtilizationGpu = utilization.Gpu
	g.UtilizationMemory = utilization.Memory

	return g, nvml.SUCCESS
}
