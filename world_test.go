package main

import (
	"testing"
)

func TestNewWorld(t *testing.T) {
	wl := NewWorld(20, 10)
	if wl.width != 20 {
		t.Error(t, "Width did not set correctly")
	}
	if wl.height != 10 {
		t.Error(t, "Height did not set correctly")
	}

	if len(wl.cells) != 200 {
		t.Error(t, "Cell did not set correctly")
	}
}

func TestWorld_isInside(t *testing.T) {
	world := NewWorld(10, 10)
	isInBounds := world.isInside(5, 9)
	if !isInBounds {
		t.Error(t, "Given vector should be inside world boundaries")
	}
}

func TestWorld_isInsideNot(t *testing.T) {
	world := NewWorld(10, 10)
	isInBounds := world.isInside(10, 10)
	if isInBounds {
		t.Error(t, "Given vector should not be inside world boundaries")
	}
}

func TestWorld_LookNotFound(t *testing.T) {
	world := NewWorld(10, 10)
	//x 0 and y 0 is not valid
	if found := world.Look(1, 1, Directions["n"]); found {
		t.Error(t, "No Cell should be found, since direction North out of range")
	}

	if found := world.Look(1, 1, Directions["ne"]); found {
		t.Error(t, "No Cell should be found, since direction North East is empty.")
	}

	if found := world.Look(1, 1, Directions["e"]); found {
		t.Error(t, "No Cell should be found, since direction East is empty.")
	}

	if found := world.Look(1, 1, Directions["se"]); found {
		t.Error(t, "Cell should be found, at x 1 y 1, we set before")
	}
	if found := world.Look(1, 1, Directions["s"]); found {
		t.Error(t, "Cell should be found")
	}
	if found := world.Look(1, 1, Directions["w"]); found {
		t.Error(t, "No Cell should be found, since direction West is empty.")
	}
	if found := world.Look(1, 1, Directions["sw"]); found {
		t.Error(t, "No Cell should be found, since direction South West is empty.")
	}
}
