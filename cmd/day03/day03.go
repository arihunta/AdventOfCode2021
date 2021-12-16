package advent2021

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Day03_01(data string) int {

	file, _ := os.Open(data)
	defer file.Close()
	var tally []int = nil

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newVal := strings.Split(scanner.Text(), "")
		if tally == nil {
			tally = make([]int, len(newVal))
		}
		for i, s := range newVal {
			if s == "1" {
				tally[i]++
			} else {
				tally[i]--
			}
		}
	}

	gam := 0
	eps := 0

	for _, s := range tally {
		gam *= 2
		eps *= 2
		if s > 0 {
			gam++
		} else {
			eps++
		}
	}

	return gam * eps

}

type Node struct {
	Zero     *Node
	One      *Node
	Children int
}

func (me *Node) Populate(number []string, idx int) {
	if idx < len(number) {
		me.Children++
	}
	if idx >= len(number) {
		return
	}
	if number[idx] == "0" {
		if me.Zero == nil {
			newNode := Node{nil, nil, 0}
			me.Zero = &newNode
		}
		me.Zero.Populate(number, idx+1)
	} else {
		if me.One == nil {
			newNode := Node{nil, nil, 0}
			me.One = &newNode
		}
		me.One.Populate(number, idx+1)
	}
}

func (me *Node) Oxygen(currentVal int) int {
	if me.Children == 0 {
		fmt.Printf("Terminating with %d\n", currentVal)
		return currentVal
	}
	if me.One != nil && me.Zero != nil {
		if me.One.Children >= me.Zero.Children {
			fmt.Printf("Selecting one %d\n", currentVal)
			return me.One.Oxygen(currentVal*2 + 1)
		} else {
			fmt.Printf("Selecting zero %d\n", currentVal)
			return me.Zero.Oxygen(currentVal * 2)
		}
	} else if me.One != nil {
		fmt.Printf("defaulting one %d\n", currentVal)
		return me.One.Oxygen(currentVal*2 + 1)
	} else if me.Zero != nil {
		fmt.Printf("defaulting one %d\n", currentVal)
		return me.Zero.Oxygen(currentVal * 2)
	}
	return 0
}

func (me *Node) Co2(currentVal int) int {
	if me.Children == 0 {
		fmt.Printf("Terminating with %d\n", currentVal)
		return currentVal
	}
	if me.One != nil && me.Zero != nil {
		if me.One.Children < me.Zero.Children {
			fmt.Printf("Selecting one %d\n", currentVal)
			return me.One.Co2(currentVal*2 + 1)
		} else {
			fmt.Printf("Selecting zero %d\n", currentVal)
			return me.Zero.Co2(currentVal * 2)
		}
	} else if me.One != nil {
		fmt.Printf("defaulting one %d\n", currentVal)
		return me.One.Co2(currentVal*2 + 1)
	} else if me.Zero != nil {
		fmt.Printf("defaulting one %d\n", currentVal)
		return me.Zero.Co2(currentVal * 2)
	}
	return 0
}

func Day03_02(data string) int {

	file, _ := os.Open(data)
	defer file.Close()
	head := Node{nil, nil, 0}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newVal := strings.Split(scanner.Text(), "")
		// fmt.Println(newVal)
		head.Populate(newVal, 0)
		// fmt.Printf(">> %d\n", head.Children)
		// if head.One != nil {
		// 	fmt.Printf(">> one %d\n", head.One.Children)
		// }
		// if head.Zero != nil {
		// 	fmt.Printf(">> zero %d\n", head.Zero.Children)
		// }
	}

	oxygen := head.Oxygen(0)
	co2 := head.Co2(0)

	return oxygen * co2
}
