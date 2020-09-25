package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type configuration struct {
	Port           int
	ScrapeInterval time.Duration
	RequestTimeout time.Duration
	Devices        []device
}

type device struct {
	DisplayName string
	Username    string
	Password    string
	MACAddress  string
	Type        string
}

func (d device) getStatusURL() string {
	return fmt.Sprintf("http://shelly%s-%s/status", d.Type, d.MACAddress)
}

func getConfig() configuration {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal error in config file: %s \n", err))
	}

	var config configuration
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to unmarshal config into struct, %v", err)
	}

	return config
}
