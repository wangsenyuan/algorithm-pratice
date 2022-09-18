package p2408

import "testing"

func runSample(t *testing.T, arriveAlice string, leaveAlice string, arriveBob string, leaveBob string, expect int) {
	res := countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "08-15", "08-18", "08-16", "08-19", 3)
}

func TestSample2(t *testing.T) {
	runSample(t, "10-01", "10-31", "11-01", "12-31", 0)
}
