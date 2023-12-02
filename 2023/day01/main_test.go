package main

import "testing"

func TestPartOneCaseOne(t *testing.T) {
	calibration := partOne("input_test.txt")
	expected := 142
	if calibration != expected {
		t.Errorf("Calibration returned %d, expected %d", calibration, expected)
	}
}

func TestPartOneCasetwo(t *testing.T) {
	calibration := partOne("input.txt")
	expected := 55712
	if calibration != expected {
		t.Errorf("Calibration returned %d, expected %d", calibration, expected)
	}
}
