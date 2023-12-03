package main

import (
	"fmt"
	partone "github.com/patrickovm/advent-of-code/2023/day01/part_one"
	parttwo "github.com/patrickovm/advent-of-code/2023/day01/part_two"
)

func main() {
	fmt.Printf("Solution to part one: %d\n", partone.Calibration("input.txt"))

	fmt.Printf("Solution to part two: %d\n", parttwo.Calibration("input.txt"))
}
