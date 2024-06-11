# NVIDIA Server Repository
The NVIDIA Server repository is a Go-based project that provides an HTTP server to expose GPU metrics. The server is designed to collect and serve information about NVIDIA GPU devices, making it easy to monitor and manage these devices.

## Key Features:

* GPU Metrics Collection: The server collects information about multiple NVIDIA GPU devices, including their names, UUIDs, temperatures, power usage, memory information, and utilization rates.
* HTTP Server: The server exposes a single endpoint, /gpuinfo, which returns the collected GPU metrics in JSON format.
* Error Handling: Robust error handling mechanisms are implemented to ensure that any errors encountered during metric collection or HTTP requests are properly handled and returned.

## Code Organization:

* internal/metrics: This package contains the gpumetrics module, which provides functions to collect GPU metrics using the go-nvml library.
* server: This package contains the GpuHttpServer function, which creates and starts an HTTP server to expose the collected GPU metrics.
Dependencies:

* go-nvml: The project relies on the go-nvml package from NVIDIA to interact with GPU devices.
Standard Go libraries: The project uses standard Go libraries, including context, encoding/json, log, net/http, and time.
Usage:

To use this repository, simply clone it and run the GpuHttpServer function, passing in a gpumetrics.GpuDeviceManager instance, an address string, and a GPU count integer. The server will start listening on the specified address and expose the /gpuinfo endpoint.

## Example:

```Go

ctx := context.Background()
deviceManager := gpumetrics.NewGpuDeviceManager()
err := GpuHttpServer(ctx, deviceManager, ":8080", 4)
if err != nil {
    log.Fatal(err)
}
```
This will start the server on port 8080 and collect information about 4 NVIDIA GPU devices. You can then use a tool like curl to retrieve the GPU metrics:

```Shell
go run cmd/main.go

curl http://localhost:8080/gpuinfo | jq . 
```
## License:

This repository is licensed under the Apache License 2.0. See LICENSE for details.

Contributing:

Contributions are welcome! If you'd like to contribute to this project, please fork the repository and submit a pull request.
