package server

import (
	"context"
	"encoding/json"
	gpumetrics "nvidia-server/internal/metrics"

	"log"
	"net/http"
	"time"
)

// GpuHttpServer creates a http server to expose the GPU metrics curl http://localhost:8080/gpuinfo | jq .
func GpuHttpServer(ctx context.Context, device gpumetrics.GpuDeviceManager, address string, gpuCount int) error {
	http.HandleFunc("/gpuinfo", func(w http.ResponseWriter, r *http.Request) {

		info, err := gpumetrics.FetchAllGpuInfo(ctx, device, gpuCount)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json, err := json.Marshal(info)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(json)
		log.Default().Printf("Collected GPU Info at %s\n", time.Now().Format("2006-01-02 15:04:05"))

	})

	// start the server
	server := &http.Server{
		Addr:           address,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    15 * time.Second,
		MaxHeaderBytes: 1 << 20,
		Handler:        nil,
	}

	log.Default().Printf("Starting GPU Metrics HTTP Server at %s\n", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil

}
