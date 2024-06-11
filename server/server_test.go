package server

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"github.com/stretchr/testify/mock"
)

func TestGpuHttpServer(t *testing.T) {
	// Create a mock GpuDeviceManager
	mockDevice := &MockGpuDeviceManager{}

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create a test server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GpuHttpServer(ctx, mockDevice, ":8080", 2)
	}))
	defer ts.Close()

	// Make a GET request to the test server
	resp, err := http.Get(ts.URL + "/gpuinfo")
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	// TODO: Add assertions for the response body
	bodyStr := string(body)
	if bodyStr != "" {
		t.Errorf("Expected response body 'OK', got %s", bodyStr)
	}

}

type MockGpuDeviceManager struct {
	mock.Mock
}

func (m *MockGpuDeviceManager) DeviceGetCount() (int, nvml.Return) {
	args := m.Called()
	return args.Int(0), args.Get(1).(nvml.Return)
}

func (m *MockGpuDeviceManager) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
	args := m.Called(index)
	return args.Get(0).(nvml.Device), args.Get(1).(nvml.Return)
}

func (m *MockGpuDeviceManager) GetUUID() (string, nvml.Return) {
	args := m.Called()
	return args.String(0), args.Get(1).(nvml.Return)
}

func (m *MockGpuDeviceManager) GetDevice() nvml.Device {
	args := m.Called()
	return args.Get(0).(nvml.Device)
}

// mock init
func (m *MockGpuDeviceManager) Init() nvml.Return {
	args := m.Called()
	return args.Get(0).(nvml.Return)
}

func (m *MockGpuDeviceManager) Shutdown() nvml.Return {
	args := m.Called()
	return args.Get(0).(nvml.Return)
}
