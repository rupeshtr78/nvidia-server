package gpumetrics

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

var _ = Describe("FetchGpuMetrics", func() {
	var (
		fakeGpuDeviceManager *MockGpuDeviceManager
		index                int
		expectedGpuResult    *GpuResult
	)

	BeforeEach(func() {
		fakeGpuDeviceManager = &MockGpuDeviceManager{}
		index = 0
		expectedGpuResult = &GpuResult{
			index: index,
			gpu:   &NvidiaDevice{},
		}
	})

	Context("when device handle gets fetched successfully", func() {
		BeforeEach(func() {
			fakeGpuDeviceManager.On("DeviceGetHandleByIndex", index).Return(expectedGpuResult.gpu, nvml.SUCCESS)
		})

		It("returns the fetched GPU result and no error", func() {
			gpuResult, err := fetchGpuMetrics(fakeGpuDeviceManager, index)

			Expect(err).NotTo(HaveOccurred())
			Expect(gpuResult).To(Equal(expectedGpuResult))
		})
	})

	Context("when device handle fetch fails", func() {
		var expectedError nvml.Return

		BeforeEach(func() {
			expectedError = nvml.ERROR_UNKNOWN
			fakeGpuDeviceManager.On("DeviceGetHandleByIndex", index).Return(nil, expectedError)
		})

		It("returns an error", func() {
			_, err := fetchGpuMetrics(fakeGpuDeviceManager, index)

			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError(fmt.Sprintf("failed to get device handle: %v", expectedError)))
		})
	})
})
