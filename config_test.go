package main

import (
	"bytes"
	"github.com/spf13/viper"
	"reflect"
	"testing"
	"time"
)

func TestReadConfig(t *testing.T) {
	var yamlContent = []byte(`
port: 9123
requestTimeout: 5s
scrapeInterval: 60s
devices:
  - macAddress: "ABC12345"
    displayName: "livingRoomShutter"
    type: "switch25"
    username: "some-user"
    password: "pass123"
  - macAddress: "123DEF"
    displayName: "kitchenShutter"
    type: "switch25"
    username: "another-user"
    password: "secure"
`)

	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(yamlContent))

	var givenConfig configuration
	err := viper.Unmarshal(&givenConfig)
	if err != nil {
		t.Errorf("could not unmarshal status endpoint content: %v", err)
	}

	expectedConfig := configuration{
		Port:           9123,
		RequestTimeout: time.Second * time.Duration(5),
		ScrapeInterval: time.Second * time.Duration(60),
		Devices: []device{
			{
				MACAddress:  "ABC12345",
				DisplayName: "livingRoomShutter",
				Type:        "switch25",
				Username:    "some-user",
				Password:    "pass123",
			},
			{
				MACAddress:  "123DEF",
				DisplayName: "kitchenShutter",
				Type:        "switch25",
				Username:    "another-user",
				Password:    "secure",
			},
		},
	}

	if !reflect.DeepEqual(givenConfig, expectedConfig) {
		t.Error("given config file does not match expected config")
	}

	if givenConfig.Devices[0].getStatusURL() != "http://shellyswitch25-ABC12345/status" {
		t.Errorf("getStatusURL() = %s; want http://shellyswitch25-ABC12345/status", givenConfig.Devices[0].getStatusURL())
	}

	if givenConfig.Devices[1].getStatusURL() != "http://shellyswitch25-123DEF/status" {
		t.Errorf("getStatusURL() = %s; want http://shellyswitch25-123DEF/status", givenConfig.Devices[1].getStatusURL())
	}
}
