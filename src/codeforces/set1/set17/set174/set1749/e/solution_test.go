package main

import "testing"

func runSample(t *testing.T, land []string, expect []string) {
	ok, res := solve(land)

	if len(expect) > 0 != ok {
		t.Fatalf("Sample expect %v, but got %t, %v", expect, ok, res)
	}
	if len(expect) == 0 {
		return
	}

	a := countCacti(expect)
	b := countCacti(res)

	if a != b {
		t.Fatalf("Sample expect %v, but got  %v", expect, res)
	}

	for r := 0; r < len(res); r++ {
		for c := 0; c < len(res[r]); c++ {
			if res[r][c] == '#' {
				for i := 0; i < 4; i++ {
					u, v := r+dd[i], c+dd[i+1]
					if u >= 0 && u < len(res) && v >= 0 && v < len(res[r]) && res[u][v] == '#' {
						t.Fatalf("Sample expect %v, but got  %v", expect, res)
					}
				}
			}
		}
	}

}

func countCacti(land []string) int {
	var res int
	for _, row := range land {
		for i := 0; i < len(row); i++ {
			if row[i] == '#' {
				res++
			}
		}
	}
	return res
}

func TestSample1(t *testing.T) {
	land := []string{
		".#..",
		"..#.",
	}
	expect := []string{
		".#.#",
		"#.#.",
	}
	runSample(t, land, expect)
}

func TestSample2(t *testing.T) {
	land := []string{
		"#.#",
		"...",
		".#.",
	}

	runSample(t, land, nil)
}

func TestSample3(t *testing.T) {
	land := []string{
		".....",
		".....",
		".....",
		".....",
		".....",
	}
	expect := []string{
		"....#",
		"...#.",
		"..#..",
		".#...",
		"#....",
	}
	runSample(t, land, expect)
}

func TestSample4(t *testing.T) {
	land := []string{
		"#..",
		".#.",
		"#.#",
		"...",
	}
	expect := []string{
		"#..",
		".#.",
		"#.#",
		"...",
	}
	runSample(t, land, expect)
}
