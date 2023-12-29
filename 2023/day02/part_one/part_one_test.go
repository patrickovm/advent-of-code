package partone

import (
	"log"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	result, err := SumOfPossibleIDs("input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := 8

	if result != expected {
		t.Errorf("Sum of IDs of valid games returned %d, expected %d", result, expected)
	}
}

func TestPartOneInput(t *testing.T) {
	result, err := SumOfPossibleIDs("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := 1734

	if result != expected {
		t.Errorf("Sum of IDs of valid games returned %d, expected %d", result, expected)
	}
}
