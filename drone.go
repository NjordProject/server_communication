package main

import (
	"time"
)

type Drone struct {
	Number int
	Driver
}

func (d Drone) Run(message_from_drone_channel chan MessageFromDrone,
				   message_from_server_channel chan MessageFromServer) {
	for {
		message_from_drone := d.Driver.Receive(d)
		message_from_drone_channel <- message_from_drone
		select {
		case msg, ok := <- message_from_server_channel:
			if ok {
				d.Driver.Send(d, msg)
			} else {
				Wg_drones.Done()
				return
			}
		default:
		}

		time.Sleep(1 * time.Second)
	}
}