package partone

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Calibration(inputPath string) int {
	lines, err := ReadLines(inputPath)
	if err != nil {
		log.Fatal(err)
	}

	var num int
	var sum int

	for _, v := range lines {
		num, err = findNumber(v)
		if err != nil {
			log.Fatal(err)
		}
		sum += num
	}

	return sum
}

func ReadLines(filePath string) ([]string, error) {
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

	var inputLines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputLines = append(inputLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return inputLines, nil
}

func findNumber(s string) (int, error) {
	re := regexp.MustCompile(`\d`)

	matches := re.FindAllString(s, -1)

	if len(matches) == 0 {
		return -1, fmt.Errorf("no number found in the string")
	}

	firstNum := matches[0]
	lastNum := matches[len(matches)-1]

	numberString := fmt.Sprintf("%s%s", firstNum, lastNum)

	result, err := strconv.Atoi(numberString)
	if err != nil {
		return -1, fmt.Errorf("failed to convert the matched string to a number: %v", err)
	}

	return result, nil
}
