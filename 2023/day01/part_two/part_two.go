package parttwo

import (
	"fmt"
	partone "github.com/patrickovm/advent-of-code/2023/day01/part_one"
	"log"
	"regexp"
	"strconv"
)

func Calibration(inputPath string) int {
	input, err := partone.ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	sum := 0
	re := regexp.MustCompile(`^(\d|one|two|three|four|five|six|seven|eight|nine)`)

	for _, line := range input {
		var currentLine []string

		for i := range line {
			found := re.FindString(line[i:])
			if found != "" {
				currentLine = append(currentLine, found)
			}
		}

		for index, word := range currentLine {
			if !isNumber(word) && word != "" {
				currentLine[index] = wordToNumber(word)
			}
		}

		firstDigit := currentLine[0]
		lastDigit := currentLine[len(currentLine)-1]

		doubleDigit := fmt.Sprintf("%s%s", firstDigit, lastDigit)
		num, err := strconv.Atoi(doubleDigit)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	return sum
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)

	return err == nil
}

func wordToNumber(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"

	default:
		return ""
	}
}
