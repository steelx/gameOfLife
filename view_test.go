package main

import "testing"

func TestWorld_isInside(t *testing.T) {
	vec1 := Vector{0, 0}
	world := &World{}
	isInBounds := world.isInside(vec1)
	if !isInBounds {
		t.Error(t, "Given vector should be inside world boundaries")
	}
}

func TestWorld_isInsideNot(t *testing.T) {
	vec1 := Vector{10, 11}
	world := &World{}
	isInBounds := world.isInside(vec1)
	if isInBounds {
		t.Error(t, "Given vector should not be inside world boundaries")
	}
}
