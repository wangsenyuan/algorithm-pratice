package p1349

import "testing"

func runSample(t *testing.T, seats []string, expect int) {
	ss := make([][]byte, len(seats))

	for i := 0; i < len(seats); i++ {
		ss[i] = []byte(seats[i])
	}

	res := maxStudents(ss)

	if res != expect {
		t.Errorf("Sample %v, expect %d, but got %d", seats, expect, res)
	}
}

func TestSample1(t *testing.T) {
	ss := []string{
		"#.##.#",
		".####.",
		"#.##.#",
	}
	expect := 4
	runSample(t, ss, expect)
}

func TestSample2(t *testing.T) {
	ss := []string{
		".#",
		"##",
		"#.",
		"##",
		".#",
	}
	expect := 3
	runSample(t, ss, expect)
}

func TestSample3(t *testing.T) {
	ss := []string{
		"#...#",
		".#.#.",
		"..#..",
		".#.#.",
		"#...#",
	}
	expect := 10
	runSample(t, ss, expect)
}

func TestSample4(t *testing.T) {
	ss := []string{
		"####",
		".###",
	}
	expect := 1
	runSample(t, ss, expect)
}
