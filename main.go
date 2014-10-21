package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var wgDrone sync.WaitGroup
var wgObs sync.WaitGroup
var wgMessageDrone sync.WaitGroup

type Coord struct {
	X int
	Y int
	Z int
}

func (c Coord) Print() {
	fmt.Println("X:", c.X, "Y:", c.Y, "Z:", c.Z)
}

type Sensor struct {
	S1 int
	S2 int
	S3 int
	S4 int
	S5 int
	S6 int
}

type MessageDrone struct {
	Position Coord
	Sensor Sensor
	CodeMsg int
}

func (md MessageDrone) Print() {
	fmt.Println("Position :")
	fmt.Println("  x :", md.Position.X)
	fmt.Println("  y :", md.Position.Y)
	fmt.Println("  z :", md.Position.Z)
	fmt.Println("Sensor :")
	fmt.Println("  s1 :", md.Sensor.S1)
	fmt.Println("  s2 :", md.Sensor.S2)
	fmt.Println("  s3 :", md.Sensor.S3)
	fmt.Println("  s4 :", md.Sensor.S4)
	fmt.Println("  s5 :", md.Sensor.S5)
	fmt.Println("  s6 :", md.Sensor.S6)
	fmt.Println("Message :", md.CodeMsg)
}

func (md MessageDrone) GetObs() [6]Coord {
	var obs [6]Coord
	obs[0] = Coord{md.Position.X, md.Position.Y, md.Position.Z - md.Sensor.S1}
	obs[1] = Coord{md.Position.X, md.Position.Y, md.Position.Z + md.Sensor.S2}
	obs[2] = Coord{md.Position.X, md.Position.Y + md.Sensor.S3, md.Position.Z}
	obs[3] = Coord{md.Position.X, md.Position.Y - md.Sensor.S4, md.Position.Z}
	obs[4] = Coord{md.Position.X - md.Sensor.S5, md.Position.Y, md.Position.Z}
	obs[5] = Coord{md.Position.X + md.Sensor.S6, md.Position.Y, md.Position.Z}
	return obs
}

type MessageServer struct {
	Direction Coord
	CodeMsg int
}

func (ms MessageServer) Print() {
	fmt.Println("Direction :")
	fmt.Println("  x :", ms.Direction.X)
	fmt.Println("  y :", ms.Direction.Y)
	fmt.Println("  z :", ms.Direction.Z)
	fmt.Println("Message :", ms.CodeMsg)
}

type Driver interface {
	Send(Drone, MessageServer) bool
	Receive(Drone) MessageDrone
}

type FakeDriver struct {
}

func (fd FakeDriver) Send(drone Drone, msg MessageServer) bool {
	return true
}

func (fd FakeDriver) Receive(drone Drone) MessageDrone {
	messageDrone := MessageDrone{Coord{X: rand.Int(), Y: rand.Int(), Z: rand.Int()},
								 Sensor{S1: rand.Int(), S2: rand.Int(), S3: rand.Int(), S4: rand.Int(), S5: rand.Int(), S6: rand.Int()},
								 0}
	return messageDrone
}

type Drone struct {
	Num int
	Driver Driver
}

func (d Drone) Run(obs chan Coord, msgDrone chan MessageDrone, msgServer chan MessageServer) {
	for {
		// Receive message from the drone
		msg_from_drone := d.Driver.Receive(d)
		// Send it for log
		msgDrone <- msg_from_drone
		// Send obstacle for store
		for _, o := range msg_from_drone.GetObs() {
			obs <- o
		}
		// Receive message from server or not
		select {
		case msg, ok := <- msgServer:
			if ok {
				d.Driver.Send(d, msg)
			} else {
				fmt.Println("Goroutine n°", d.Num, "closed")
				wgDrone.Done()
				return
			}
		default:
		}

		time.Sleep(1 * time.Second)
	}
}

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
	nbr_drone := 3
	obs := make(chan Coord, 100)
	msgDrone := make(chan MessageDrone, 100)
	messages_server := make([]chan MessageServer, nbr_drone)
	for _, ms := range messages_server {
		ms = make(chan MessageServer, 100)
	}

	wgDrone.Add(nbr_drone)
	wgObs.Add(1)
	wgMessageDrone.Add(1)

	drone_slice := make([]Drone, nbr_drone)

	go ObsHandler(obs)
	go MsgDroneHandler(msgDrone)

	for i, d := range drone_slice {
		d.Num = i
		d.Driver = FakeDriver{}
		go d.Run(obs, msgDrone, messages_server[i])
	}

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

	wgDrone.Wait()
	close(obs)
	close(msgDrone)
	wgObs.Wait()
	wgMessageDrone.Wait()
}