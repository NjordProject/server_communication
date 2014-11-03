package main

import (
	"math/rand"
)

type Driver interface {
	Send(Drone, MessageFromServer) bool
	Receive(Drone) MessageFromDrone
}

type FakeDriver struct {}

func (fd FakeDriver) Send(drone Drone, msg MessageFromServer) bool {
	return true
}

func (fd FakeDriver) Receive(drone Drone) MessageFromDrone {
	message_from_drone := MessageFromDrone{Coord{X: rand.Int(), Y: rand.Int(), Z: rand.Int()},
								 		   Sensor{S1: rand.Int(), S2: rand.Int(), S3: rand.Int(), S4: rand.Int(), S5: rand.Int(), S6: rand.Int()},
								 		    0}
	return message_from_drone
}