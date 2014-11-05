package main

import (
	"strconv"
	"fmt"
)

type Coord struct {
	X int
	Y int
	Z int
}

type Obstacle Coord

type Sensor struct {
	S1 int
	S2 int
	S3 int
	S4 int
	S5 int
	S6 int
}

type MessageFromDrone struct {
	Position Coord
	Sensor
	CodeMsg int
}

func (msg MessageFromDrone) ToString() (str string){
	str = "Position : \n" +
	      "	X : " + strconv.Itoa(msg.Position.X) + "\n" +
	      "	Y : " + strconv.Itoa(msg.Position.Y) + "\n" +
	      "	Z : " + strconv.Itoa(msg.Position.Z) + "\n"
	return str
}

type MessageFromServer struct {
	Direction Coord
	CodeMsg int
}

func ExtractObstacle(msg MessageFromDrone) [6]Obstacle {
	var obs [6]Obstacle
	obs[0] = Obstacle{msg.Position.X, msg.Position.Y, msg.Position.Z - msg.Sensor.S1}
	obs[1] = Obstacle{msg.Position.X, msg.Position.Y, msg.Position.Z + msg.Sensor.S2}
	obs[2] = Obstacle{msg.Position.X, msg.Position.Y - msg.Sensor.S3, msg.Position.Z}
	obs[3] = Obstacle{msg.Position.X, msg.Position.Y + msg.Sensor.S4, msg.Position.Z}
	obs[4] = Obstacle{msg.Position.X - msg.Sensor.S5, msg.Position.Y, msg.Position.Z}
	obs[5] = Obstacle{msg.Position.X + msg.Sensor.S6, msg.Position.Y, msg.Position.Z}
	return obs
}

func MessageFromDroneHandler(message_from_drone_channel chan MessageFromDrone) {
	for {
		message_frome_drone, ok := <- message_from_drone_channel
		if ok {
			fmt.Println(message_frome_drone.ToString())
			//Log du message
			//Enregistrement des obstacles
			//obstacles := ExtractObstacle(message_frome_drone)
		} else {
			Wg_message_from_drone.Done()
			return
		}
	}
}