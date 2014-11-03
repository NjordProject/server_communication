package main

type Coord struct {
	X int
	Y int
	Z int
}

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

type MessageFromServer struct {
	Direction Coord
	CodeMsg int
}