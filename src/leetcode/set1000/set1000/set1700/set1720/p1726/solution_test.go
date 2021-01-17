package p1726

import "testing"

func runSample(t *testing.T, grid []string, catJump, mouseJump int, expect bool) {
	res := canMouseWin(grid, catJump, mouseJump)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{"####F", "#C...", "M...."}
	cat := 1
	mouse := 2
	expect := true
	runSample(t, grid, cat, mouse, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{"M.C...F"}
	cat := 1
	mouse := 4
	expect := true
	runSample(t, grid, cat, mouse, expect)
}

func TestSample3(t *testing.T) {
	grid := []string{"M.C...F"}
	cat := 1
	mouse := 3
	expect := false
	runSample(t, grid, cat, mouse, expect)
}

func TestSample4(t *testing.T) {
	grid := []string{"C...#", "...#F", "....#", "M...."}
	cat := 2
	mouse := 5
	expect := false
	runSample(t, grid, cat, mouse, expect)
}

func TestSample5(t *testing.T) {
	grid := []string{".M...", "..#..", "#..#.", "C#.#.", "...#F"}
	cat := 3
	mouse := 1
	expect := true
	runSample(t, grid, cat, mouse, expect)
}

func TestSample6(t *testing.T) {
	grid := []string{"..#C", "...#", "..M.", "#F..", "...."}
	cat := 2
	mouse := 1
	expect := true
	runSample(t, grid, cat, mouse, expect)
}
