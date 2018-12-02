package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 576
	result := PartOne()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 77674
	result := PartTwo()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}
