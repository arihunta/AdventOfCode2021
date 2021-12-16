package advent2021

import (
	"testing"
)

func TestDay01_01_00(t *testing.T) {
	face := Day01_01("01-00")
	if face != 7 {
		t.Errorf("Expected 7, received %v", face)
	}
}

func TestDay01_01(t *testing.T) {
	face := Day01_01("01")
	if face != 1722 {
		t.Errorf("Expected 1722, received %v", face)
	}
}

func TestDay01_02_00(t *testing.T) {
	face := Day01_02("01-00")
	if face != 5 {
		t.Errorf("Expected 5, received %v", face)
	}
}

func TestDay01_02(t *testing.T) {
	face := Day01_02("01")
	if face != 1748 {
		t.Errorf("Expected 1748, received %v", face)
	}
}
