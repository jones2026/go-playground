package main

import "testing"

func TestSum(t *testing.T) {
	total := sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
	expectedSum := 10
	if total != expectedSum {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, expectedSum)
	}
}

func TestSubtraction(t *testing.T) {
	expectedDifference := 5
	actualDifference := subtraction(10, 5)
	if actualDifference != expectedDifference {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", actualDifference, expectedDifference)
	}
}
