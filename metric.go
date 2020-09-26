package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	labels = []string{"name", "mac", "type"}

	errorCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "shelly_error_count",
			Help: "Shows number of failed requests for device",
		},
		labels,
	)

	temperatureGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "shelly_temperature",
			Help: "Shows current temperature",
		},
		labels,
	)

	isOvertemperatureGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "shelly_overtemperature",
			Help: "Shows wether device is over normal temperature",
		},
		labels,
	)

	voltageGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "shelly_voltage",
			Help: "Shows current voltage"},
		labels,
	)

	uptimeGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "shelly_uptime",
			Help: "Shows current uptime"},
		labels,
	)

	isUpdateAvailableGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "shelly_update_available",
			Help: "Shows whether an update is available"},
		labels,
	)
)

func registerMetrics() {
	prometheus.Register(errorCounter)
	prometheus.Register(temperatureGauge)
	prometheus.Register(isOvertemperatureGauge)
	prometheus.Register(voltageGauge)
	prometheus.Register(uptimeGauge)
	prometheus.Register(isUpdateAvailableGauge)
}
