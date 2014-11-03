package main

import (
	"sync"
	"os"
	"encoding/json"
	"time"
)

const drones_config_file = "config/drones.json"

var wg_drone sync.WaitGroup
var drones []Drone

type droneConfig struct {
	Number int `json:"number"`
	Driver string `json:"driver"`
}

type Drone struct {
	Number int
	Driver
}

func (d Drone) Run(obstacle_channel chan Coord,
				   message_from_drone_channel chan MessageFromDrone,
				   message_from_server_channel chan MessageFromServer) {
	for {
		msg_from_drone := d.Driver.Receive(d)
		message_from_server_channel <- msg_from_drone
		for _, o := range msg_from_drone.GetObstacles() {
			obstacle_channel <- o
		}

		select {
		case msg, ok := <- message_from_server_channel:
			if ok {
				d.Driver.Send(d, msg)
			} else {
				wg_drone.Done()
				return
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func init() {
	file, err := os.Open(drones_config_file)
	if err != nil {
	}
	defer file.Close()
	var drones_config []droneConfig
	err = json.NewDecoder(file).Decode(&drones_config)
	drones = make([]Drone, len(drones_config))
	for i, dc := range drones_config {
		drones[i].Number = dc.Number
		switch dc.Driver {
		case "fake":
			drones[i].Driver = FakeDriver{}
		}
	}
	wg_drone.Add(len(drones_config))
}

func LaunchDrones(obstacle_channel chan Coord,
				  message_from_drone_channel chan MessageFromDrone,
				  message_from_server_channel []chan MessageFromServer) {
	for i, d := range drones {
		d.Run(obstacle_channel, message_from_drone_channel, message_from_server_channel[i])
	}
}