package main

import (
	"time"
	"sync"
)

var number_of_message_from_drone_handler int
var number_of_message_from_server_handler int
var number_of_drones int

var Wg_message_from_drone sync.WaitGroup
var Wg_message_from_server sync.WaitGroup
var Wg_drones sync.WaitGroup

var drones_slice []Drone
var message_from_drone_channel chan MessageFromDrone
var message_from_server_channel_slice []chan MessageFromServer

func init() {
	JsonConfig := ReadConfig("config/fake.json")
	number_of_message_from_drone_handler = JsonConfig.Number_of_message_from_drone_handler
	Wg_message_from_drone.Add(number_of_message_from_drone_handler)
	number_of_message_from_server_handler = JsonConfig.Number_of_message_from_server_handler
	Wg_message_from_server.Add(number_of_message_from_server_handler)
	number_of_drones = len(JsonConfig.Drones)
	Wg_drones.Add(number_of_drones)
	drones_slice = make([]Drone, number_of_drones)
	for i, d := range JsonConfig.Drones {
		drones_slice[i].Number = d.Number
		switch d.Driver {
		case "fake":
			drones_slice[i].Driver = FakeDriver{}
		}
	}
	message_from_drone_channel = make(chan MessageFromDrone, 100)
	message_from_server_channel_slice = make([]chan MessageFromServer, len(drones_slice))
	for i := range message_from_server_channel_slice {
		message_from_server_channel_slice[i] = make(chan MessageFromServer, 100)
	}
}

func main() {
	for i := 0; i < number_of_message_from_drone_handler; i++ {
		go MessageFromDroneHandler(message_from_drone_channel)
	}

	for i := 0; i < number_of_message_from_server_handler; i++ {
	}

	for i, d := range drones_slice {
		go d.Run(message_from_drone_channel, message_from_server_channel_slice[i])
	}

	time.Sleep(5 * time.Second)

	Wg_message_from_server.Wait()
	for _, m := range message_from_server_channel_slice {
		close(m)
	}
	Wg_drones.Wait()
	close(message_from_drone_channel)
	Wg_message_from_drone.Wait()
}