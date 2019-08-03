package main

import (
	"strconv"
	"testing"
)

func runReminder(t *testing.T, n int, K int) {
	S := strconv.Itoa(K)
	res := reminder(n, S)

	if res != K%n {
		t.Errorf("Sample %d %d, expect %d, but got %d", n, K, K%n, res)
	} else {
		t.Logf("Sample %d %d, result got %d", n, K, res)
	}
}

func TestReminder(t *testing.T) {
	runReminder(t, 10, 11)
	runReminder(t, 11, 1)
	runReminder(t, 13, 12324)

	runReminder(t, 103, 12399824)

	runReminder(t, 100, 1000000)
}
