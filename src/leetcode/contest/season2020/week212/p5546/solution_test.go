package p5546

import "testing"

func runSample(t *testing.T, time []int, keys string, expect byte) {
	res := slowestKey(time, keys)
	if res != expect {
		t.Errorf("Smaple expect %b, but got %b", expect, res)
	}
}

func TestSample1(t *testing.T) {
	releaseTimes := []int{9, 29, 49, 50}
	keysPressed := "cbcd"
	var expect byte = 'c'
	runSample(t, releaseTimes, keysPressed, expect)
}

func TestSample2(t *testing.T) {
	releaseTimes := []int{12, 23, 36, 46, 62}
	keysPressed := "spuda"
	var expect byte = 'a'
	runSample(t, releaseTimes, keysPressed, expect)
}
