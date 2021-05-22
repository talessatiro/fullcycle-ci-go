package main

import "testing"

func TestSum(test *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		test.Errorf("Invalid Test Result: Expected value = %d | Actual value = %d", 30, total)
	}
}

func TestSub(test *testing.T) {
	total := sub(10, 5)

	if total != 5 {
		test.Errorf("Invalid Test Result: Expected value = %d | Actual value = %d", 5, total)
	}
}

func TestTimes(test *testing.T) {
	total := times(5, 5)

	if total != 25 {
		test.Errorf("Invalid Test Result: Expected value = %d | Actual value = %d", 25, total)
	}
}

func TestDiv(test *testing.T) {
	total := div(10, 5)

	if total != 2 {
		test.Errorf("Invalid Test Result: Expected value = %d | Actual value = %d", 2, total)
	}
}
