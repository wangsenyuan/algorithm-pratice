package main

import "testing"

func runSample(t *testing.T, n int, x int, expect bool) {
	res := solve(n, x)

	if len(res) > 0 != expect {
		t.Fatalf("Sample expect %t, but got %s", expect, res)
	}

	if expect {
		i, j := 0, n-1
		mem := make(map[byte]int)

		for i <= j {
			if res[i] != res[j] {
				t.Fatalf("result %s, not a valid palindrom", res)
			}
			mem[res[i]]++
			mem[res[j]]++
			i++
			j--
		}

		if len(mem) != x {
			t.Fatalf("Sample result %s, has %d distinct letters other than %d", res, len(mem), x)
		}
	}
}

func TestSample1(t *testing.T) {
	runSample(t, 3, 3, false)
}

func TestSample2(t *testing.T) {
	runSample(t, 2, 1, true)
}

func TestSample3(t *testing.T) {
	runSample(t, 4, 2, true)
}

func TestSample4(t *testing.T) {
	runSample(t, 5, 2, true)
}
func TestSample5(t *testing.T) {
	runSample(t, 1, 1, true)
}
