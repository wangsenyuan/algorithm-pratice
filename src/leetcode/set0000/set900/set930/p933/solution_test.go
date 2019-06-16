package p933

import "testing"

func runSample(t *testing.T, pings []int, expect []int) {
	counter := Constructor()
	for i := 0; i < len(pings); i++ {
		res := counter.Ping(pings[i])
		if res != expect[i] {
			t.Errorf("Sample %v, expect %v, but fail at step %d, with %d", pings, expect, i, res)
			return
		}
	}
}

func TestSample1(t *testing.T) {
	pings := []int{1, 100, 3001, 3002}
	expect := []int{1, 2, 3, 3}
	runSample(t, pings, expect)
}
