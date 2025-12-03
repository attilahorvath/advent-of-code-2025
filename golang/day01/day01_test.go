package main

import (
	"strconv"
	"testing"
)

var rotations []string = []string{"L68",
	"L30",
	"R48",
	"L5",
	"R60",
	"L55",
	"L1",
	"L99",
	"R14",
	"L82"}

func TestZeros(t *testing.T) {
	dial := newDial()

	for _, rot := range rotations {
		dist, _ := strconv.Atoi(rot[1:])
		dial.rotate(rot[:1], dist)
	}

	if dial.zeros != 3 {
		t.Errorf("zeros = %d, want 3", dial.zeros)
	}
}

func TestZerosWithClicks(t *testing.T) {
	dial := newDial()

	for _, rot := range rotations {
		dist, _ := strconv.Atoi(rot[1:])
		dial.rotateClicks(rot[:1], dist)
	}

	if dial.zeros != 6 {
		t.Errorf("zeros with clicks = %d, want 6", dial.zeros)
	}
}
