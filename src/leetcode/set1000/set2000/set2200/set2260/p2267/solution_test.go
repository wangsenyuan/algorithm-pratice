package p2267

import "testing"

func runSample(t *testing.T, grid []string, expect bool) {
	g := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		g[i] = []byte(grid[i])
	}
	res := hasValidPath(g)

	if res != expect {
		t.Errorf("Sample expect %t, but got %t", expect, res)
	}
}

func TestSample1(t *testing.T) {
	grid := []string{
		"(((",
		")((",
		"(()",
		"(()",
	}
	expect := true

	runSample(t, grid, expect)
}

func TestSample2(t *testing.T) {
	grid := []string{
		"))",
		"((",
	}
	expect := false

	runSample(t, grid, expect)
}


func TestSample3(t *testing.T) {
	grid := []string{
		"()",
		"()",
	}
	expect := false

	runSample(t, grid, expect)
}
