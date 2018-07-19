package main

import "testing"

func runSample(t *testing.T, M, N int, room []string, expect int) {
	rm := make([][]byte, M)
	for i := 0; i < M; i++ {
		rm[i] = []byte(room[i])
	}
	res := solve(M, N, rm)
	if res != expect {
		t.Errorf("Sample %d %d %v, expect %d, but got %d", M, N, room, expect, res)
	}

}
func TestSample1(t *testing.T) {
	M, N := 2, 3
	room := []string{"...", "..."}
	expect := 4
	runSample(t, M, N, room, expect)
}

func TestSample2(t *testing.T) {
	M, N := 2, 3
	room := []string{"x.x", "xxx"}
	expect := 1
	runSample(t, M, N, room, expect)
}

func TestSample3(t *testing.T) {
	M, N := 2, 3
	room := []string{"x.x", "x.x"}
	expect := 2
	runSample(t, M, N, room, expect)
}

func TestSample4(t *testing.T) {
	M, N := 10, 10
	room := []string{
		"....x.....",
		"..........",
		"..........",
		"..x.......",
		"..........",
		"x...x.x...",
		".........x",
		"...x......",
		"........x.",
		".x...x....",
	}
	expect := 46
	runSample(t, M, N, room, expect)
}
