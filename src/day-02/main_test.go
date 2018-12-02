package main

import (
	"testing"
)

func TestPartOne(t *testing.T) {
	expected := 8118
	result := PartOne()
	if result != expected {
		t.Errorf("result: %d", result)
	}
}

func TestPartTwo(t *testing.T) {
	expected := "jbbenqtlaxhivmwyscjukztdp"
	result := PartTwo()
	if result != expected {
		t.Errorf("result: %s", result)
	}
}
