package util

import (
	"fmt"
)

type DirectedGraph struct {
	Map Grid
	at Coordinate
	Visits map[string]int
}

type Direction string

const (
	North = Direction("N")
	South = Direction("S")
	East  = Direction("E")
	West  = Direction("W")
)

func (dg *DirectedGraph) SetCoordinate(coordinate Coordinate) *DirectedGraph {
	dg.Map[coordinate.String()] = coordinate
	dg.Visits[coordinate.String()]++

	return dg
}

func (dg *DirectedGraph) Move(direction Direction) *DirectedGraph {

	delete(dg.Map, dg.at.String())
	switch (direction) {
	case North:
		dg.at.Y++
	case South:
		dg.at.Y--
	case East:
		dg.at.X++
	case West:
		dg.at.X--
	}

	dg.SetCoordinate(dg.at)

	return dg
}

func NewDirectedGraph (value interface{}) (*DirectedGraph) {
	dg := DirectedGraph{
		Map: Grid{},
		at: Coordinate{0, 0, value},
		Visits: map[string]int{},
	}

	dg.SetCoordinate(dg.at)

	return &dg
}

type Grid map[string]Coordinate

type Coordinate struct {
	X int
	Y int
	Value interface{}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
