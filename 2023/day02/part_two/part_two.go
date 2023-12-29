package parttwo

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func minimumSetOfCubes(game string) (int, int, int) {
	var minSet = map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	subsets := strings.Split(game, "; ")

	for _, subset := range subsets {
		cubes := strings.Split(subset, ", ")

		for _, cube := range cubes {
			cubeParts := strings.Split(cube, " ")
			count, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				log.Fatal(err)
			}

			color := cubeParts[1]

			if minSet[color] < count {
				minSet[color] = count
			}
		}
	}

	return minSet["red"], minSet["green"], minSet["blue"]
}

func SumOfPower(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	minimumSetPowerSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		game := strings.Split(scanner.Text(), ": ")[1]
		minRed, minGreen, minBlue := minimumSetOfCubes(game)

		minimumSetPowerSum += minRed * minGreen * minBlue

	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return minimumSetPowerSum, nil
}
