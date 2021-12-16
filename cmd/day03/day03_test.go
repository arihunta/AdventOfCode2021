package advent2021

import (
	"testing"
)

func TestDay03_01_00(t *testing.T) {
	face := Day03_01("03-00")
	if face != 198 {
		t.Errorf("Expected 198, received %v", face)
	}
}

func TestDay03_01(t *testing.T) {
	face := Day03_01("03")
	if face != 3429254 {
		t.Errorf("Expected 3429254, received %v", face)
	}
}

func TestDay03_02_00(t *testing.T) {
	face := Day03_02("03-00")
	if face != 230 {
		t.Errorf("Expected 230, received %v", face)
	}
}

func TestDay03_02(t *testing.T) {
	face := Day03_02("03")
	if face != 5410338 {
		t.Errorf("Expected 5410338, received %v", face)
	}
}
