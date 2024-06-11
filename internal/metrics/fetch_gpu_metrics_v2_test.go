package gpumetrics

// func TestFetchAllGpuInfo_HappyPath(t *testing.T) {
// 	// Create a mock GpuDeviceManager
// 	gpu := &mockGpuDeviceManager{}

// 	// Set up the mock to return 2 devices
// 	gpu.On("DeviceGetCount").Return(2, nvml.SUCCESS)

// 	// Set up the mock to return device handles for each index
// 	gpu.On("DeviceGetHandleByIndex", 0).Return(&nvml.Device{}, nvml.SUCCESS)
// 	gpu.On("DeviceGetHandleByIndex", 1).Return(&nvml.Device{}, nvml.SUCCESS)

// 	// Set up the mock to return metrics for each device
// 	gpu.On("FetchDeviceMetrics", &nvml.Device{}).Return(&NvidiaDevice{}, nvml.SUCCESS)
// 	gpu.On("FetchDeviceMetrics", &nvml.Device{}).Return(&NvidiaDevice{}, nvml.SUCCESS)

// 	// Create a context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Call the function under test
// 	gpuMap, err := FetchAllGpuInfo(ctx, gpu, 2)

// 	// Assert the result
// 	if err != nil {
// 		t.Errorf("Expected no error, but got %v", err)
// 	}
// 	if len(gpuMap) != 2 {
// 		t.Errorf("Expected 2 GPUs in the map, but got %d", len(gpuMap))
// 	}
// }

// func TestFetchAllGpuInfo_NoDevices(t *testing.T) {
// 	// Create a mock GpuDeviceManager
// 	gpu := &mockGpuDeviceManager{}

// 	// Set up the mock to return 0 devices
// 	gpu.On("DeviceGetCount").Return(0, nvml.SUCCESS)

// 	// Create a context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Call the function under test
// 	gpuMap, err := FetchAllGpuInfo(ctx, gpu, 0)

// 	// Assert the result
// 	if err != nil {
// 		t.Errorf("Expected no error, but got %v", err)
// 	}
// 	if len(gpuMap) != 0 {
// 		t.Errorf("Expected 0 GPUs in the map, but got %d", len(gpuMap))
// 	}
// }

// func TestFetchAllGpuInfo_ContextCancellationV2(t *testing.T) {
// 	// Create a mock GpuDeviceManager
// 	gpu := &mockGpuDeviceManager{}

// 	// Set up the mock to return 2 devices
// 	gpu.On("DeviceGetCount").Return(2, nvml.SUCCESS)

// 	// Create a context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Cancel the context immediately
// 	cancel()

// 	// Call the function under test
// 	gpuMap, err := FetchAllGpuInfo(ctx, gpu, 2)

// 	// Assert the result
// 	if err == nil {
// 		t.Errorf("Expected an error due to context cancellation, but got none")
// 	}
// }

// func TestFetchAllGpuInfo_ContextCancellation(t *testing.T) {
// 	// Create a mock GpuDeviceManager
// 	gpu := &mockGpuDeviceManager{}

// 	// Set up the mock to return 2 devices
// 	gpu.On("DeviceGetCount").Return(2, nvml.SUCCESS)

// 	// Create a context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Cancel the context immediately
// 	cancel()

// 	// Call the function under test
// 	gpuMap, err := FetchAllGpuInfo(ctx, gpu, 2)

// 	// Assert the result
// 	if err == nil {
// 		t.Errorf("Expected an error due to context cancellation, but got none")
// 	}
// }

// func TestFetchAllGpuInfo_FetchDeviceMetricsError(t *testing.T) {
// 	// Create a mock GpuDeviceManager
// 	gpu := &mockGpuDeviceManager{}

// 	// Set up the mock to return 2 devices
// 	gpu.On("DeviceGetCount").Return(2, nvml.SUCCESS)

// 	// Set up the mock to return device handles for each index
// 	gpu.On("DeviceGetHandleByIndex", 0).Return(&nvml.Device{}, nvml.SUCCESS)
// 	gpu.On("DeviceGetHandleByIndex", 1).Return(&nvml.Device{}, nvml.SUCCESS)

// 	// Set up the mock to return an error for metrics retrieval
// 	gpu.On("FetchDeviceMetrics", &nvml.Device{}).Return(nil, nvml.ERROR_UNKNOWN)

// 	// Create a context
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// Call the function under test
// 	gpuMap, err := FetchAllGpuInfo(ctx, gpu, 2)

// 	// Assert the result
// 	if err == nil {
// 		t.Errorf("Expected an error due to metrics retrieval failure, but got none")
// 	}
// }

// type mockGpuDeviceManager struct {
// 	mock.Mock
// }

// func (m *mockGpuDeviceManager) Init() nvml.Return {
// 	return nvml.SUCCESS
// }

// func (m *mockGpuDeviceManager) Shutdown() nvml.Return {
// 	return nvml.SUCCESS
// }

// func (m *mockGpuDeviceManager) GetDevice() nvml.Device {
// 	// Return a mock device
// 	return m.GetDevice()
// }

// func (m *mockGpuDeviceManager) DeviceGetCount() (int, nvml.Return) {
// 	return 1, nvml.SUCCESS
// }

// func (m *mockGpuDeviceManager) DeviceGetHandleByIndex(index int) (nvml.Device, nvml.Return) {
// 	if index == 0 {
// 		// Return a mock device
// 		return m.GetDevice(), nvml.SUCCESS
// 	} else {
// 		return nil, nvml.ERROR_UNKNOWN
// 	}
// }
