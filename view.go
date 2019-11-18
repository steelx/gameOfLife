package main

type World [][]Cell

func (wl World) isInside(target Vector) bool {

	var (
		width  = len(wl[0])
		height = len(wl)
	)
	if target.x >= 0 && target.x < width &&
		target.y >= 0 && target.y < height {
		return true
	}
	return false
}

type View struct {
	world World
	cell  Cell
}

func (vw View) Look(dir Vector) bool {
	var target = vw.cell.Plus(dir)
	if vw.world.isInside(target) {
		var found = vw.world[target.y][target.x]
		return found.Alive
	}
	return false
}
