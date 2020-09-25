package main

import (
	"encoding/json"
	"testing"
)

func TestStatusResponseToStruct(t *testing.T) {
	statusEndpointResponse := []byte(`
{
  "wifi_sta": {
    "connected": true,
    "ssid": "my_wifi",
    "ip": "192.168.178.2",
    "rssi": -84
  },
  "cloud": {
    "enabled": false,
    "connected": false
  },
  "mqtt": {
    "connected": false
  },
  "time": "16:44",
  "unixtime": 1601052276,
  "serial": 12345,
  "has_update": true,
  "mac": "ABCDEF123456",
  "cfg_changed_cnt": 4,
  "actions_stats": {
    "skipped": 0
  },
  "rollers": [
    {
      "state": "stop",
      "power": 0.00,
      "is_valid": true,
      "safety_switch": false,
      "overtemperature": false,
      "stop_reason": "normal",
      "last_direction": "open",
      "current_pos": 94,
      "calibrating": false,
      "positioning": true
    }
  ],
  "meters": [
    {
      "power": 0.00,
      "overpower": 0.00,
      "is_valid": true,
      "timestamp": 1601052276,
      "counters": [
        0.000,
        0.000,
        0.000
      ],
      "total": 1773
    },
    {
      "power": 0.00,
      "overpower": 0.00,
      "is_valid": true,
      "timestamp": 1601052276,
      "counters": [
        0.000,
        0.000,
        0.000
      ],
      "total": 1910
    }
  ],
  "inputs": [
    {
      "input": 0,
      "event": "",
      "event_cnt": 0
    },
    {
      "input": 1,
      "event": "",
      "event_cnt": 0
    }
  ],
  "temperature": 60.49,
  "overtemperature": false,
  "tmp": {
    "tC": 60.49,
    "tF": 140.88,
    "is_valid": true
  },
  "update": {
    "status": "pending",
    "has_update": true,
    "new_version": "20200827-065456/v1.8.3@4a8bc427",
    "old_version": "20200812-091015/v1.8.0@8acf41b0"
  },
  "ram_total": 49504,
  "ram_free": 35496,
  "fs_size": 233681,
  "fs_free": 146082,
  "voltage": 234.84,
  "uptime": 3725427
}
`)
	var status StatusResponse

	err := json.Unmarshal(statusEndpointResponse, &status)
	if err != nil {
		t.Errorf("could not unmarshal status endpoint content: %v", err)
	}

	if status.Temperature != 60.49 {
		t.Errorf("Temperature = %f; want 60.49", status.Temperature)
	}

	if status.Overtemperature != false {
		t.Errorf("Overtemperature = %t; want false", status.Overtemperature)
	}

	if status.Voltage != 234.84 {
		t.Errorf("Voltage = %f; want 234.84", status.Voltage)
	}

	if status.Uptime != 3725427 {
		t.Errorf("Uptime = %d; want 3725427", status.Uptime)
	}

	if status.HasUpdate != true {
		t.Errorf("HasUpdate = %t; want true", status.HasUpdate)
	}

	if status.MACAddress != "ABCDEF123456" {
		t.Errorf("MACAddress = %s; want ABCDEF123456", status.MACAddress)
	}
}
