package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 146622
	result := PartOne()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 31848
	result := PartTwo()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}
