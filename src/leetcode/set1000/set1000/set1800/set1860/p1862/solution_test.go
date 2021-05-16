package p1862

import (
	"testing"
)

func runSample(t *testing.T, box []string, expect []string) {
	b := make([][]byte, len(box))
	for i, r := range box {
		b[i] = []byte(r)
	}
	res := rotateTheBox(b)

	for i := 0; i < len(expect); i++ {
		if expect[i] != string(res[i]) {
			t.Errorf("Sample expect %v, but got %s at %d", expect, string(res[i]), i)
		}
	}
}

func TestSample1(t *testing.T) {
	box := []string{
		"#.#",
	}
	expect := []string{
		".", "#", "#",
	}
	runSample(t, box, expect)
}

func TestSample2(t *testing.T) {
	box := []string{
		"#.*.",
		"##*.",
	}
	expect := []string{
		"#.",
		"##",
		"**",
		"..",
	}
	runSample(t, box, expect)
}

func TestSample3(t *testing.T) {
	box := []string{
		"##*.*.",
		"###*..",
		"###.#.",
	}
	expect := []string{
		".##",
		".##",
		"##*",
		"#*.",
		"#.*",
		"#..",
	}
	runSample(t, box, expect)
}
