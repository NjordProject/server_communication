package main

import (
	"sync"
)

type Obstacle Coord

var wg_obstacle sync.WaitGroup

func init() {
	wg_obstacle.Add(1)
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

func ObstacleHandler(obs chan Obstacle) {

}