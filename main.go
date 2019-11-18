package main

import (
	"fmt"
	tm "github.com/buger/goterm"
	"strings"
	"time"
)

func main() {
	//Cell
	var initial_grid = [][]int{
		{0, 0, 1, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 1, 0, 0, 0, 0, 0},
		{1, 0, 0, 1, 1, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
	}

	//create world
	var world = make(World, 10)
	for y, row := range initial_grid {
		world[y] = make([]Cell, len(row))
		for x, space := range row {
			var alive bool
			if space == 1 {
				alive = true
			}
			var cell = Cell{Pos: Vector{x, y}, Alive: alive}
			world[y][x] = cell
		}
	}

	for i := 0; i < 100; i++ {
		tm.Clear()
		box := tm.NewBox(40|tm.PCT, 10, 0)
		var printThis []string
		//update world
		for y, row := range world {
			for x, cell := range row {
				var neighbours = findNeighbours(world, cell)
				cell.NextState(neighbours)
				world[y][x] = cell
			}
			printThis = append(printThis, strings.Join(buildRow(row)[:], " "))
		}
		time.Sleep(1 * time.Second)
		fmt.Fprint(box, strings.Join(printThis[:], "\n"))
		// Move Box to approx center of the screen
		tm.Print(tm.MoveTo(box.String(), 40|tm.PCT, 40|tm.PCT))
		tm.Flush()
	}
}

func getChar(cell Cell) string {
	if cell.Alive {
		return "*"
	}
	return " "
}

func buildRow(row []Cell) []string {
	var strip []string
	for _, cell := range row {
		strip = append(strip, getChar(cell))
	}
	return strip
}

func findNeighbours(world World, cell Cell) (count int) {
	view := &View{world, cell}
	for _, direction := range DirectionNames {
		if view.Look(Directions[direction]) {
			count += 1
		}
	}
	return
}
