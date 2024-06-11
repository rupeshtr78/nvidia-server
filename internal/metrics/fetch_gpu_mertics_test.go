package gpumetrics

import (
	"context"
	"testing"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/mock"
)

// Mock nvidia.Device
type Device struct {
	mock.Mock
}

func NewDevice() *NvidiaDevice {
	return &NvidiaDevice{}
}

func (d *Device) Init() nvml.Return {
	args := d.Called()
	return args.Get(0).(nvml.Return)
}

func (d *Device) Shutdown() nvml.Return {
	args := d.Called()
	return args.Get(0).(nvml.Return)
}

func (d *Device) GetDevice() nvml.Device {
	args := d.Called()
	return args.Get(0).(nvml.Device)
}

func (d *Device) DeviceGetCount() (int, nvml.Return) {
	args := d.Called()
	return args.Get(0).(int), args.Get(1).(nvml.Return)
}

func (d *Device) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	args := d.Called(index)
	return args.Get(0).(nvml.Device), args.Get(1).(nvml.Return)
}

var device = NewDevice()

func TestFetchAllGpuInfo(t *testing.T) {
	tests := []struct {
		name   string
		m      *NvidiaDevice
		device nvml.Device
		want   error
	}{
		{
			name:   "test fetch all gpu info",
			m:      &NvidiaDevice{},
			device: device.GetDevice(),
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Init()
			defer tt.m.Shutdown()
			ctx := context.Background()
			count := 2
			if _, ret := FetchAllGpuInfo(ctx, device, count); ret != tt.want {
				t.Errorf("FetchAllGpuInfo = %v, want %v", ret, tt.want)
			}
		})
	}
}
