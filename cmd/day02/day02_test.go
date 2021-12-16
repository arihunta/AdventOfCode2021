package advent2021

import (
	"testing"
)

func TestDay02_01_00(t *testing.T) {
	face := Day02_01("02-00")
	if face != 150 {
		t.Errorf("Expected 150, received %v", face)
	}
}

func TestDay02_01(t *testing.T) {
	face := Day02_01("02")
	if face != 2215080 {
		t.Errorf("Expected 2215080, received %v", face)
	}
}

func TestDay02_02_00(t *testing.T) {
	face := Day02_02("02-00")
	if face != 900 {
		t.Errorf("Expected 900, received %v", face)
	}
}

func TestDay02_02(t *testing.T) {
	face := Day02_02("02")
	if face != 1864715580 {
		t.Errorf("Expected 1864715580, received %v", face)
	}
}
