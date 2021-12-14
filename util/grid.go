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

func MakeFullGrid(x int, y int, value interface{}) (Grid) {

	grid := Grid{}

	for i := 0; i <= x;i++ {
		for j := 0; j <= y;j++ {
			grid.SetValue(i, j, value)
		}
	}

	return grid
}

func MergeGrids(a Grid, b Grid) Grid {

	newGrid := Grid{}

	aMaxY := a.getMaxY()
	aMaxX := a.getMaxX()
	bMaxY := b.getMaxY()
	bMaxX := b.getMaxX()

	for i := 0; i <= aMaxY;i++ {
		for j := 0; j <= aMaxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			if a[key].Value != nil {
				newGrid.SetValue(j, i, a[key].Value)
			} else {
				newGrid.SetValue(j, i, nil)
			}
		}
	}

	for i := 0; i <= bMaxY;i++ {
		for j := 0; j <= bMaxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			if b[key].Value != nil {
				newGrid.SetValue(j, i, b[key].Value)
			}
		}
	}

	return newGrid
}

func (g Grid) FillGrid(value interface{}) {

	maxY := g.getMaxY()
	maxX := g.getMaxX()

	for i := 0; i <= maxY;i++ {
		for j := 0; j <= maxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			if g[key].Value == nil {
				g.SetValue(j, i, value)
			}
		}
	}
}

func (g Grid) FlipVertically() {

	newGrid := g.Clone()
	g.Clear()

	maxY := newGrid.getMaxY()
	maxX := newGrid.getMaxX()

	for i := maxY; i >= 0;i-- {
		for j := 0; j <= maxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			g.SetValue(j, maxY - i, newGrid[key].Value)
		}
	}
}

func (g Grid) FlipHorzontially() {

	newGrid := g.Clone()
	g.Clear()

	maxY := newGrid.getMaxY()
	maxX := newGrid.getMaxX()

	for i := 0; i <= maxY;i++ {
		for j := maxX; j >= 0;j-- {
			key := fmt.Sprintf("%d,%d", j, i)
			g.SetValue(maxX - j, i, newGrid[key].Value)
		}
	}
}

func (g Grid) Subset(minX int, maxX int, minY int, maxY int) Grid {

	newGrid := Grid{}

	for i := minY; i <= maxY;i++ {
		for j := minX; j <= maxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			newGrid.SetValue(Abs(minX - j), Abs(minY - i), g[key].Value)
		}
	}

	return newGrid
}

func (g Grid) Clear() {

	maxY := g.getMaxY()
	maxX := g.getMaxX()

	for i := 0; i <= maxY;i++ {
		for j := 0; j <= maxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			delete(g, key)
		}
	}
}

func (g Grid) Clone() Grid {

	newGrid := Grid{}

	maxY := g.getMaxY()
	maxX := g.getMaxX()

	for i := 0; i <= maxY;i++ {
		for j := 0; j <= maxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			newGrid.SetValue(j, i, g[key].Value)
		}
	}

	return newGrid
}

func (g Grid) PrintGrid(padding int) {

	maxY := g.getMaxY()
	maxX := g.getMaxX()

	for i := 0; i <= maxY;i++ {
		fmt.Println("")
		for j := 0; j <= maxX;j++ {
			key := fmt.Sprintf("%d,%d", j, i)
			if g[key].Value != nil {
				paddingStr := strconv.Itoa(padding)
				fmt.Printf("%" + paddingStr + "v", (g[key].Value))
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

// Returns all coordinates between and inclusive of the given start and end
func (g Grid) GetPointsBetween(start Coordinate, end Coordinate) []Coordinate {

	coordinates := []Coordinate{start}

	if end.X == start.X && end.Y == start.Y {
		return coordinates
	}

	slopeX := end.X - start.X
	slopeY := end.Y - start.Y

	gcd := Gcd(slopeX, slopeY)

	slopeX = Abs(slopeX / gcd)
	slopeY = Abs(slopeY / gcd)

	if end.X < start.X {
		slopeX = -slopeX
	}

	if end.Y < start.Y {
		slopeY = -slopeY
	}

	// No slope given
	if slopeX == 0 && slopeY == 0 {
		return coordinates;
	}

	atX := start.X + slopeX
	atY := start.Y + slopeY

	for {
		newCoordinate := g.GetCoordinate(atX, atY)

		coordinates = append(coordinates, newCoordinate)
		atX += slopeX
		atY += slopeY

		if newCoordinate.String() == end.String() {
			break
		}
	}

	return coordinates
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

func (g Grid) GetAdjacent(coor Coordinate) []Coordinate {
	adjacent := []Coordinate{}

	// Above
	if coor.Y > 0 {
		adjacent = append(adjacent, g.GetCoordinate(coor.X, coor.Y - 1))
	}

	// Right
	if coor.X < g.getMaxX() {
		adjacent = append(adjacent, g.GetCoordinate(coor.X + 1, coor.Y))
	}

	// Below
	if coor.Y < g.getMaxY() {
		adjacent = append(adjacent, g.GetCoordinate(coor.X, coor.Y + 1))
	}

	// Left
	if coor.X > 0 {
		adjacent = append(adjacent, g.GetCoordinate(coor.X - 1, coor.Y))
	}

	return adjacent
}

func (g Grid) GetSurrounding(coor Coordinate) []Coordinate {
	adjacent := []Coordinate{}

	// Above
	if coor.Y > 0 {
		adjacent = append(adjacent, g.GetCoordinate(coor.X, coor.Y - 1))
	}

	// Above left
	if coor.Y > 0 && coor.X > 0 {
		adjacent = append(adjacent, g.GetCoordinate(coor.X - 1, coor.Y - 1))
	}

	// Above right
	if coor.Y > 0 && coor.X < g.getMaxX() {
		adjacent = append(adjacent, g.GetCoordinate(coor.X + 1, coor.Y - 1))
	}

	// Right
	if coor.X < g.getMaxX() {
		adjacent = append(adjacent, g.GetCoordinate(coor.X + 1, coor.Y))
	}

	// Below
	if coor.Y < g.getMaxY() {
		adjacent = append(adjacent, g.GetCoordinate(coor.X, coor.Y + 1))
	}

	// Below left
	if coor.Y < g.getMaxY() && coor.X > 0 {
		adjacent = append(adjacent, g.GetCoordinate(coor.X - 1, coor.Y + 1))
	}

	// Below Right
	if coor.Y < g.getMaxY() &&  coor.X < g.getMaxX() {
		adjacent = append(adjacent, g.GetCoordinate(coor.X + 1, coor.Y + 1))
	}

	// Left
	if coor.X > 0 {
		adjacent = append(adjacent, g.GetCoordinate(coor.X - 1, coor.Y))
	}

	return adjacent
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

func (g Grid) getMinY() int {

	min := 99999999999999 // not great but lazy

	for _,coor := range g {
		if coor.Y < min {
			min = coor.Y
		}
	}

	return min
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

func (g Grid) GetMaxX() int {
	return g.getMaxX()
}

func (g Grid) GetMinX() int {
	return g.getMinX()
}

func (g Grid) GetMinY() int {
	return g.getMinY()
}

func (g Grid) GetMaxY() int {
	return g.getMaxY()
}



type Coordinate struct {
	X int
	Y int
	Value interface{}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}
