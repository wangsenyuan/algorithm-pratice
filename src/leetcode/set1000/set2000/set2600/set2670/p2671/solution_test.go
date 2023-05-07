package p2671

import "testing"

func TestSample1(t *testing.T) {
	tracker := Constructor()
	tracker.Add(3)
	tracker.Add(3)
	res := tracker.HasFrequency(2)
	if !res {
		t.Fatalf("There should be 2 3's , but got false")
	}
}

func TestSample2(t *testing.T) {
	tracker := Constructor()
	tracker.Add(3)
	tracker.Add(3)
	res := tracker.HasFrequency(1)
	if res {
		t.Fatalf("There should no 1 3's , but got true")
	}
}

func TestSample3(t *testing.T) {
	tracker := Constructor()
	tracker.Add(2)
	tracker.Add(3)
	res := tracker.HasFrequency(1)
	if !res {
		t.Fatalf("There should no 1 3's , but got true")
	}
}
