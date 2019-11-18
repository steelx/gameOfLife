package main

import "strings"

var DirectionNames = strings.Split("n ne e se s sw w nw", " ")
var Directions = map[string]Vector{
	"n":  {0, -1},
	"ne": {1, -1},
	"e":  {1, 0},
	"se": {1, 1},
	"s":  {0, 1},
	"sw": {1, 1},
	"w":  {1, 0},
	"nw": {1, -1},
}

type Vector struct {
	x int
	y int
}

type Cell struct {
	Pos   Vector
	Alive bool
}

func NewCell(x, y int, alive bool) *Cell {
	c := new(Cell)
	c.Alive = alive
	c.Pos.x = x
	c.Pos.y = y
	return c
}

func (c *Cell) NextState(neighbours int) {
	if neighbours < 2 || neighbours > 3 || neighbours == 0 {
		c.Alive = false
	}

	if neighbours == 2 || neighbours == 3 {
		c.Alive = true
	}

	if !c.Alive && neighbours == 2 {
		c.Alive = true
	}
}

func (c Cell) Plus(vec Vector) Vector {
	return NewCell(c.Pos.x+vec.x, c.Pos.y+vec.y, c.Alive).Pos
}
