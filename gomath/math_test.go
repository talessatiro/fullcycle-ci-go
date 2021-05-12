package main

import "testing"

func TestSum(test *testing.T) {
	total := sum(15, 15)

	if total != 30 {
		test.Errorf("Invalid Test Result: Expected value = %d | Actual value = %d", 30, total)
	}
}
