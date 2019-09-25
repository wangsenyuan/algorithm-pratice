package problem3

import "testing"

func runSample(t *testing.T, command string, obstacles [][]int, x int, y int, expect bool) {
	res := robot(command, obstacles, x, y)
	if res != expect {
		t.Errorf("Sample %v %v %d %d, expect %t, but got %t", command, obstacles, x, y, expect, res)
	}
}

func TestSample1(t *testing.T) {
	command := "URR"
	obstacles := [][]int{}
	x := 3
	y := 2
	expect := true
	runSample(t, command, obstacles, x, y, expect)
}

func TestSample2(t *testing.T) {
	command := "URR"
	obstacles := [][]int{{2, 2}}
	x := 3
	y := 2
	expect := false
	runSample(t, command, obstacles, x, y, expect)
}

func TestSample3(t *testing.T) {
	command := "URR"
	obstacles := [][]int{{4, 2}}
	x := 3
	y := 2
	expect := true
	runSample(t, command, obstacles, x, y, expect)
}
