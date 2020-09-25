package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func main() {
	config := getConfig()
	registerMetrics()

	go func(config configuration) {
		fetchDevices(config)
		for range time.Tick(config.ScrapeInterval) {
			fetchDevices(config)
		}
	}(config)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(fmt.Sprintf("localhost:%d", config.Port), nil)
}
