package advent2021

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Day02_01(data string) int {

	file, _ := os.Open(data)
	defer file.Close()

	x := 0
	y := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "forward":
			x = x + val
		case "up":
			y = y - val
		case "down":
			y = y + val
		}
	}

	return x * y

}

func Day02_02(data string) int {

	file, _ := os.Open(data)
	defer file.Close()

	x := 0
	y := 0
	pitch := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(instruction[1])
		switch instruction[0] {
		case "forward":
			x = x + val
			y = y + val*pitch
		case "up":
			pitch = pitch - val
		case "down":
			pitch = pitch + val
		}
	}

	return x * y

}
