package main

import "testing"

func TestCase1(t *testing.T) {
	// arrange
	pressures := []float32{1000, 850, 1000}
	expected := 1

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase2(t *testing.T) {
	// arrange
	pressures := []float32{850, 960}
	expected := -1

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase3(t *testing.T) {
	// arrange
	pressures := []float32{850, 960}
	expected := -1

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase4(t *testing.T) {
	// arrange
	pressures := []float32{850, 960}
	expected := -1

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase5(t *testing.T) {
	// arrange
	pressures := []float32{30, 960, 343, 2345, 2342}
	expected := 0

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase6(t *testing.T) {
	// arrange
	pressures := []float32{600, 600, 600, 600, 30}
	expected := 4

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase7(t *testing.T) {
	// arrange
	pressures := []float32{600, 600, 400, 600, 30}
	expected := 2

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}

func TestCase8(t *testing.T) {
	// arrange
	pressures := []float32{600}
	expected := 0

	// act
	result := stormEye(pressures)

	// assert
	if result != expected {
		t.Error("expected ", expected, " got ", result)
	}
}
