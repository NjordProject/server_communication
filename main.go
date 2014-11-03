package main

import (
	"fmt"
	"time"
	"sync"
)

var wgObs sync.WaitGroup
var wgMessageDrone sync.WaitGroup

func ObsHandler(obs chan Coord) {
	for {
		o, ok := <- obs
		if ok {
			o.Print()
		} else {
			fmt.Println("No more obstacle, close")
			wgObs.Done()
			return
		}
	}
}

func MsgDroneHandler(msg_drone chan MessageDrone) {
	for {
		md, ok := <- msg_drone
		if ok {
			md.Print()
		} else {
			fmt.Println("No more message from drone, close")
			wgMessageDrone.Done()
			return
		}
	}
}

func main() {
	obs := make(chan Coord, 100)
	msgDrone := make(chan MessageDrone, 100)
	messages_server := make([]chan MessageServer, len(drones))
	for i := range messages_server {
		messages_server[i] = make(chan MessageServer, 100)
	}

	wgObs.Add(1)
	wgMessageDrone.Add(1)

	go ObsHandler(obs)
	go MsgDroneHandler(msgDrone)

	LaunchDrones(obs, msgDrone, messages_server)

	time.Sleep(5 * time.Second)

	for i := 0; i < 5; i++ {
		msg := MessageServer{Coord{i, i, i}, i}
		for _, ms := range messages_server {
			ms <- msg
		}
	}

	for _, ms := range messages_server {
		close(ms)
	}

	wg_drone.Wait()
	close(obs)
	close(msgDrone)
	wgObs.Wait()
	wgMessageDrone.Wait()
}