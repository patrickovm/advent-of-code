package main

import (
	"fmt"
	partone "github.com/patrickovm/advent-of-code/2023/day02/part_one"
	parttwo "github.com/patrickovm/advent-of-code/2023/day02/part_two"
	"log"
)

func main() {

	partOneResult, err := partone.SumOfPossibleIDs("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	partTwoResult, err := parttwo.SumOfPower("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution to part one: %d\n", partOneResult)
	fmt.Printf("Solution to part two: %d\n", partTwoResult)
}
