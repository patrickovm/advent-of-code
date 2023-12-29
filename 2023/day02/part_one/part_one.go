package partone

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var maxCubeCounts = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func isGamePossible(game string) bool {
	sets := strings.Split(game, "; ")

	for _, subset := range sets {
		cubes := strings.Split(subset, ", ")

		for _, cube := range cubes {
			cubeParts := strings.Split(cube, " ")
			count, err := strconv.Atoi(cubeParts[0])
			if err != nil {
				log.Fatal(err)
			}

			color := cubeParts[1]

			if count > maxCubeCounts[color] {
				return false
			}
		}
	}

	return true
}

func SumOfPossibleIDs(filePath string) (int, error) {
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

	possibleGamesIdSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		gameParts := strings.Split(scanner.Text(), ": ")
		info, game := gameParts[0], gameParts[1]

		if isGamePossible(game) {
			gameId, _ := strconv.Atoi(strings.Split(info, " ")[1])
			possibleGamesIdSum += gameId
		}

	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}

	return possibleGamesIdSum, nil
}
