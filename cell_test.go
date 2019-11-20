package main

import "testing"

func TestCell_NextState_withZeroNeighbours(t *testing.T) {
	cell := Cell{true}
	cell.NextState(0)

	if cell.Alive {
		t.Error(t, "Cell should die due to loneliness")
	}
}

func TestCell_NextState_withTwoNeighbours(t *testing.T) {
	cell := Cell{true}
	cell.NextState(2)

	if !cell.Alive {
		t.Error(t, "Cell should stay alive")
	}
}

func TestCell_NextState_withFourNeighbours(t *testing.T) {
	cell := Cell{true}
	cell.NextState(4)

	if cell.Alive {
		t.Error(t, "Cell should die due to over-population")
	}
}

func TestDeadCell_NextState_withTwoNeighbours(t *testing.T) {
	cell := Cell{false}
	cell.NextState(2)

	if !cell.Alive {
		t.Error(t, "Cell should get alive with correct conditions")
	}
}
