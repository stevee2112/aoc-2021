package util

import (
	"fmt"
	"strconv"
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

func (dg *DirectedGraph) At() Coordinate {
	return dg.at
}

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

func (g Grid) SetValue(x int, y int, value interface{}) {
	coordinate := Coordinate{x, y, value}
	g[coordinate.String()] = coordinate
}

func (g Grid) GetCoordinate(x int, y int) Coordinate {
	return g[fmt.Sprintf("%d,%d", x, y)]
}

func (g Grid) PrintGrid(padding int) {

	for i := 0; i <= g.getMaxY();i++ {
		fmt.Println("")
		for j := 0; j <= g.getMaxX();j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			if g[key].Value != nil {
				fmt.Printf("%" + strconv.Itoa(padding) + "s ", (g[key].Value.(string)))
			}
		}
	}
	fmt.Println("");
	fmt.Println("");
}

func (g Grid) Traverse(action func(coor Coordinate) bool) {
	for _,coordinate := range g {
		if !action(coordinate) { // stop if false
			return
		}
	}
}

func (g Grid) GetRows() (rows [][]Coordinate) {
	for i := 0; i <= g.getMaxY();i++ {
		row := []Coordinate{}
		for j := 0; j <= g.getMaxX();j++ {
			row = append(row, g.GetCoordinate(j, i))
		}
		rows = append(rows,row)
	}

	return rows;
}

func (g Grid) GetCols() (cols [][]Coordinate) {
	for i := 0; i <= g.getMaxX();i++ {
		col := []Coordinate{}
		for j := 0; j <= g.getMaxY();j++ {
			col = append(col, g.GetCoordinate(i, j))
		}
		cols = append(cols,col)
	}

	return cols;
}

func (g Grid) getMinX() int {

	min := 99999999999999 // not great but lazy

	for _,coor := range g {
		if coor.X < min {
			min = coor.X
		}
	}

	return min
}

func (g Grid) getMaxX() int {

	max := -99999999999999 // not great but lazy

	for _,coor := range g {
		if coor.X > max {
			max = coor.X
		}
	}

	return max
}

func (g Grid) getMaxY() int {

	max := -999999999 // not great but lazy

	for _,coor := range g {
		if coor.Y > max {
			max = coor.Y
		}
	}

	return max
}


type Coordinate struct {
	X int
	Y int
	Value interface{}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
