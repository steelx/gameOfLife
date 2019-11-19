package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type World struct {
	cells  []Cell
	width  int
	height int
}

/*
 x → →

 y
 ↓
 ↓
*/

func NewWorld(width, height int) *World {
	return &World{
		cells:  make([]Cell, width*height),
		width:  width,
		height: height,
	}
}

func (wl World) generateMap(percent int) {
	percentageAlive := percent * len(wl.cells) / 100

	// Fill Alive Cells as per percentage given
	for i := percentageAlive; i > 0; i-- {
		wl.cells[i] = Cell{true}
	}

	// Randomize Alive Cells
	cellsClone := wl.cells
	seed := rand.New(rand.NewSource(time.Now().Unix()))
	for i := len(wl.cells); i > 0; i-- {
		randomIndex := seed.Intn(i)
		wl.cells[i-1], cellsClone[randomIndex] = cellsClone[randomIndex], wl.cells[i-1]
	}
	wl.cells = cellsClone
}

//Print to screen
func (wl World) print() {
	for y := 0; y < wl.height; y++ {
		//creates new row
		fmt.Print("█")
		//columns
		for x := 0; x < wl.width; x++ {
			fmt.Print(getChar(wl.getCell(x, y)))
		}
		fmt.Println("█")
	}
}

func (wl World) isInside(x, y int) bool {
	return x >= 0 && x < wl.width && y >= 0 && y < wl.height
}

// Look checks a Cell at given direction
// and returns true if a Cell is found
func (wl World) Look(x, y int, dir Vector) bool {
	if !wl.isInside(x, y) {
		return false
	}
	return wl.Plus(x, y, dir)
}

func (wl World) getCell(x, y int) bool {
	return wl.cells[x+(y*wl.width)].Alive
}

func (wl *World) setCell(x, y int, alive bool) {
	if !wl.isInside(x, y) {
		fmt.Printf("Cordinates %v and %v are not in range! \n", x, y)
		os.Exit(1)
	}

	wl.cells[x+(y*wl.width)] = Cell{alive}
}

// Plus accepts x, y of Cell PLUS direction vector
// returns if Cell exists in the given direction
func (wl World) Plus(x, y int, vec Vector) bool {
	return wl.getCell(x+vec.x, y+vec.y)
}

// findNeighbours of given co-ordinates
func (wl World) findNeighbours(x, y int) (count int) {
	for _, direction := range DirectionNames {
		if wl.Look(x, y, Directions[direction]) {
			count++
		}
	}
	return
}

func getChar(isAlive bool) string {
	if isAlive {
		return "*"
	}
	return " "
}
