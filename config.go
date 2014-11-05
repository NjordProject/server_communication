package main

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
)

type JsonConfig struct {
	Number_of_message_from_drone_handler int
	Number_of_message_from_server_handler int
	Drones []DroneConfig
}

type DroneConfig struct {
	Number int
	Driver string
}

func ReadConfig(filePath string) (config JsonConfig) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Can't open file")
		os.Exit(1)
	}
	var jsonConfig JsonConfig
	json.Unmarshal(file, &jsonConfig)
	return jsonConfig
}