package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 117505
	result := PartOne()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 1254
	result := PartTwo()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}
