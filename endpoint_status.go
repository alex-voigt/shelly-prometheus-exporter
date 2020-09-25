package main

type StatusResponse struct {
	Device          device
	WiFiStatus      WiFi    `json:"wifi_sta"`
	Cloud           Cloud   `json:"cloud"`
	MQTT            MQTT    `json:"mqtt"`
	Serial          int     `json:"serial"`
	HasUpdate       bool    `json:"has_update"`
	MACAddress      string  `json:"mac"`
	Temperature     float32 `json:"temperature"`
	Overtemperature bool    `json:"overtemperature"`
	MemoryTotal     int     `json:"ram_total"`
	MemoryFree      int     `json:"ram_free"`
	FilesystemSize  int     `json:"fs_size"`
	FilesystemFree  int     `json:"fs_free"`
	Voltage         float32 `json:"voltage"`
	Uptime          int     `json:"uptime"`
}

type WiFi struct {
	Connected bool    `json:"connected"`
	SSID      string  `json:"ssid"`
	IP        string  `json:"ip"`
	RSSI      float32 `json:"rssi"`
}

type Cloud struct {
	Enabled   bool `json:"enabled"`
	Connected bool `json:"connected"`
}

type MQTT struct {
	Connected bool `json:"connected"`
}
