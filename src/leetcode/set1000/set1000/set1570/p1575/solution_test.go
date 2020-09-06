package p1575

import "testing"

func runSample(t *testing.T, location []int, start, finish int, fuel int, expect int) {
	res := countRoutes(location, start, finish, fuel)

	if res != expect {
		t.Errorf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	locations := []int{2, 3, 6, 8, 4}
	start := 1
	finish := 3
	fuel := 5
	expect := 4
	runSample(t, locations, start, finish, fuel, expect)
}

func TestSample2(t *testing.T) {
	locations := []int{4, 3, 1}
	start := 1
	finish := 0
	fuel := 6
	expect := 5
	runSample(t, locations, start, finish, fuel, expect)
}

func TestSample3(t *testing.T) {
	locations := []int{1, 2, 3}
	start := 0
	finish := 2
	fuel := 40
	expect := 615088286
	runSample(t, locations, start, finish, fuel, expect)
}
