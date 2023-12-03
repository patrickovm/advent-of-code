package partone

import (
	"testing"
)

func TestPartOneExample(t *testing.T) {
	calibration := Calibration("input_test.txt")
	expected := 142
	if calibration != expected {
		t.Errorf("Calibration returned %d, expected %d", calibration, expected)
	}
}

func TestPartOneInput(t *testing.T) {
	calibration := Calibration("../input.txt")
	expected := 55712
	if calibration != expected {
		t.Errorf("Calibration returned %d, expected %d", calibration, expected)
	}
}
