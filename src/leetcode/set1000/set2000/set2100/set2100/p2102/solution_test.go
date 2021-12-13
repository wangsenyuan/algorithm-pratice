package p2102

import "testing"

func TestSample1(t *testing.T) {
	tracker := Constructor()
	tracker.Add("bradford", 2)
	tracker.Add("branford", 3)

	res := tracker.Get()

	if res != "branford" {
		t.Fatalf("should get branford at 1th Get(), but got %s", res)
	}

	tracker.Add("alps", 2)

	res = tracker.Get()

	if res != "alps" {
		t.Fatalf("should get alps at 2nd Get() but got %s", res)
	}

	tracker.Add("orland", 2)

	res = tracker.Get()

	if res != "bradford" {
		t.Fatalf("should get bradford at 3rd Get(), but got %s", res)
	}

	tracker.Add("orlando", 3)

	res = tracker.Get()

	if res != "bradford" {
		t.Fatalf("should get bradford at 4rd Get(), but got %s", res)
	}

	tracker.Add("alpine", 2)

	res = tracker.Get()

	if res != "bradford" {
		t.Fatalf("should get bradford at 5th Get(), but got %s", res)
	}

	res = tracker.Get()

	if res != "orland" {
		t.Fatalf("should get orland at 6th Get(), bug got %s", res)
	}

}
