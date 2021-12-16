package advent2021

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func Day01_01(data string) int {

	file, _ := os.Open(data)
	defer file.Close()
	count := -1
	lastVal := math.MinInt32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newVal, _ := strconv.Atoi(scanner.Text())
		// fmt.Printf("%d > %d = %t -> %d\n", newVal, lastVal, (newVal > lastVal), count)
		if newVal > lastVal {
			count = count + 1
		}
		lastVal = newVal
	}

	return count

}

func Day01_02(data string) int {

	file, _ := os.Open(data)
	defer file.Close()
	count := 0
	lastVal := 0

	var buff [3]int
	buffIdx := 0

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	lastVal = lastVal - buff[buffIdx]
	buff[buffIdx], _ = strconv.Atoi(scanner.Text())
	lastVal = lastVal + buff[buffIdx]
	buffIdx = (buffIdx + 1) % 3

	scanner.Scan()
	lastVal = lastVal - buff[buffIdx]
	buff[buffIdx], _ = strconv.Atoi(scanner.Text())
	lastVal = lastVal + buff[buffIdx]
	buffIdx = (buffIdx + 1) % 3

	scanner.Scan()
	lastVal = lastVal - buff[buffIdx]
	buff[buffIdx], _ = strconv.Atoi(scanner.Text())
	lastVal = lastVal + buff[buffIdx]
	buffIdx = (buffIdx + 1) % 3

	fmt.Println(buff)
	fmt.Println(lastVal)

	for scanner.Scan() {
		newVal := lastVal - buff[buffIdx]
		buff[buffIdx], _ = strconv.Atoi(scanner.Text())
		newVal = newVal + buff[buffIdx]
		// fmt.Printf("%d > %d = %t -> %d\n", newVal, lastVal, (newVal > lastVal), count)
		if newVal > lastVal {
			count = count + 1
		}
		buffIdx = (buffIdx + 1) % 3
		lastVal = newVal
	}

	return count

}
