package parttwo

import (
	"log"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	result, err := SumOfPower("../input_test.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := 2286

	if result != expected {
		t.Errorf("Sum of IDs of valid games returned %d, expected %d", result, expected)
	}

}

func TestPartOneInput(t *testing.T) {
	result, err := SumOfPower("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	expected := 70387

	if result != expected {
		t.Errorf("Sum of IDs of valid games returned %d, expected %d", result, expected)
	}
}
