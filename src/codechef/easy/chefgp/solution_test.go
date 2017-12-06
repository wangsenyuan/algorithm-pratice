package main

import "testing"

func runSample(t *testing.T, x int, y int, s string, expectKiwi int) {
	res := solve(x, y, []byte(s))

	ax, by, kc := 0, 0, 0
	for i := 0; i < len(res); i++ {
		if res[i] == 'a' {
			ax++
			by = 0
		} else if res[i] == 'b' {
			by++
			ax = 0
		} else {
			ax = 0
			by = 0
			kc++
		}
		if ax > x || by > y {
			t.Errorf("sample %s should not have continuse a (more than %d) or b (more than %d)", s, x, y)
		}
	}
	if expectKiwi != kc {
		t.Errorf("sample %s should only have %d kiwis, but got %d, in %s", s, expectKiwi, kc, res)
	}
}
func TestSample1(t *testing.T) {
	runSample(t, 1, 1, "ab", 0)
}

func TestSample2(t *testing.T) {
	runSample(t, 1, 2, "aabb", 0)
}

func TestSample3(t *testing.T) {
	runSample(t, 2, 1, "aaaaab", 1)
}

func TestSample4(t *testing.T) {
	runSample(t, 1, 3, "aaaa", 3)
}

func TestSample5(t *testing.T) {
	runSample(t, 2, 2, "aaaaa", 2)
}
