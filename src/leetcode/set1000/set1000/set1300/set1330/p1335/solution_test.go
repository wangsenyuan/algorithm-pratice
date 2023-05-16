package p1335

import "testing"

func runSample(t *testing.T, job []int, d int, expect int) {
	res := minDifficulty(job, d)

	if res != expect {
		t.Fatalf("Sample expect %d, but got %d", expect, res)
	}
}

func TestSample1(t *testing.T) {
	job := []int{6, 5, 4, 3, 2, 1}
	d := 2
	expect := 7
	runSample(t, job, d, expect)
}

func TestSample2(t *testing.T) {
	job := []int{9, 9, 9}
	d := 4
	expect := -1
	runSample(t, job, d, expect)
}

func TestSample3(t *testing.T) {
	job := []int{7, 1, 7, 1, 7, 1}
	d := 3
	expect := 15
	runSample(t, job, d, expect)
}

func TestSample4(t *testing.T) {
	job := []int{11, 111, 22, 222, 33, 333, 44, 444}
	d := 6
	expect := 843
	runSample(t, job, d, expect)
}
