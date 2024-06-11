package main

import (
	"context"
	"log"
	gpumetrics "nvidia-server/internal/metrics"
	"nvidia-server/server"

	"github.com/NVIDIA/go-nvml/pkg/nvml"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	device := gpumetrics.NewNvidiaDevice()

	intErr := device.Init()
	if intErr != nvml.SUCCESS {
		log.Fatal(intErr)
	}
	defer device.Shutdown()

	count, ret := device.DeviceGetCount()
	if ret != nvml.SUCCESS || count == 0 {
		log.Fatalf("failed to get device count: %v", ret)
	}

	address := ":8080"

	err := server.GpuHttpServer(ctx, device, address, count)
	if err != nil {
		log.Fatal(err)
	}
}
