package gpumetrics

import (
	"io"
	"log"
	"os"
	"testing"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

// TestMain exists to suppress log output during tests
func TestMain(m *testing.M) {
	log.SetOutput(io.Discard)
	os.Exit(m.Run())
}

func TestNewGpuMap(t *testing.T) {
	got := NewGpuMap()
	if got == nil {
		t.Errorf("NewGpuMap() = %v; want non-nil", got)
	}
}

func TestNewNvidiaDevice(t *testing.T) {
	got := NewNvidiaDevice()
	if got == nil {
		t.Errorf("NewNvidiaDevice() = %v; want non-nil", got)
	}
}

func TestNvidiaDevice_Init(t *testing.T) {
	tests := []struct {
		name string
		m    *NvidiaDevice
		want nvml.Return
	}{
		{
			name: "Test NvidiaDevice Init",
			m:    &NvidiaDevice{},
			want: nvml.SUCCESS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Init(); got != tt.want {
				t.Errorf("NvidiaDevice.Init() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNvidiaDevice_Shutdown(t *testing.T) {
	tests := []struct {
		name string
		m    *NvidiaDevice
		want nvml.Return
	}{
		{
			name: "test nvml shutdown",
			m:    &NvidiaDevice{},
			want: nvml.SUCCESS,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Init()
			if got := tt.m.Shutdown(); got != tt.want {
				t.Errorf("NvidiaDevice.Shutdown()= %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNvidiaDevice_DeviceGetCount(t *testing.T) {
	tests := []struct {
		name string
		m    *NvidiaDevice
		want nvml.Return
	}{
		{
			name: "test device count",
			m:    &NvidiaDevice{},
			want: nvml.SUCCESS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Init()
			defer tt.m.Shutdown()
			if _, ret := tt.m.DeviceGetCount(); ret != tt.want {
				t.Errorf("NvidiaDevice.DeviceGetCount() = %v, want %v", ret, tt.want)
			}
		})
	}
}

func TestNvidiaDevice_DeviceGetHandleByIndex(t *testing.T) {
	tests := []struct {
		name  string
		m     *NvidiaDevice
		index int
		want  nvml.Return
	}{
		{
			name:  "test device handle by index",
			m:     &NvidiaDevice{},
			index: 0,
			want:  nvml.SUCCESS,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Init()
			defer tt.m.Shutdown()
			if _, got := tt.m.DeviceGetHandleByIndex(tt.index); got != tt.want {
				t.Errorf("NvidiaDevice.DeviceGetHandleByIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNvidiaDevice_GetDevice(t *testing.T) {
	tests := []struct {
		name string
		m    *NvidiaDevice
		want nvml.Device
	}{
		{
			name: "test get device",
			m:    &NvidiaDevice{},
			want: (nvml.Device)(nil),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Init()
			defer tt.m.Shutdown()
			if got := tt.m.GetDevice(); got != tt.want {
				t.Errorf("NvidiaDevice.GetDevice() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestNvidiaDevice_Metrics(t *testing.T) {
// 	tests := []struct {
// 		name   string
// 		m      *NvidiaDevice
// 		sensor nvml.TemperatureSensors
// 		want   nvml.Return
// 	}{
// 		{
// 			name:   "test device memory info",
// 			m:      &NvidiaDevice{},
// 			sensor: nvml.TEMPERATURE_GPU,
// 			want:   nvml.SUCCESS,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.m.Init()
// 			defer tt.m.Shutdown()
// 			if _, got := tt.m.GetMemoryInfo(); got != tt.want {
// 				t.Errorf("NvidiaDevice.GetMemoryInfo() = %v, want %v", got, tt.want)
// 			}

// 			if _, got := tt.m.GetUUID(); got != tt.want {
// 				t.Errorf("NvidiaDevice.GetUUID() = %v, want %v", got, tt.want)
// 			}

// 			if _, got := tt.m.GetTemperature(tt.sensor); got != tt.want {
// 				t.Errorf("NvidiaDevice.GetTemperature() = %v, want %v", got, tt.want)
// 			}

// 			if _, got := tt.m.GetPowerUsage(); got != tt.want {
// 				t.Errorf("NvidiaDevice.GetPowerUsage() = %v, want %v", got, tt.want)
// 			}

// 			if _, got := tt.m.GetUtilizationRates(); got != tt.want {
// 				t.Errorf("NvidiaDevice.GetUtilizationRates() = %v, want %v", got, tt.want)
// 			}

// 			if _, got := tt.m.GetName(); got != tt.want {
// 				t.Errorf("NvidiaDevice.GetName() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}

// }
