package parttwo

import (
	"testing"
)

func TestPartTwoExample(t *testing.T) {
	calibration := Calibration("input_test.txt")
	expected := 281
	if calibration != expected {
		t.Errorf("Calibration returned %d, expected %d", calibration, expected)
	}
}

func TestPartTwoInput(t *testing.T) {
	calibration := Calibration("../input.txt")
	expected := 55413
	if calibration != expected {
		t.Errorf("Calibration returned %d, expected %d", calibration, expected)
	}
}
