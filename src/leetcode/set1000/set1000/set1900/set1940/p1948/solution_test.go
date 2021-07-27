package p1948

import (
	"testing"
)

func runSample(t *testing.T, paths [][]string, expect [][]string) {
	deleteDuplicateFolder(paths)
}

func TestSample1(t *testing.T) {
	paths := [][]string{
		{"c"}, {"b"}, {"p"}, {"c", "a", "b", "a"}, {"c", "a", "b"}, {"c", "a"}, {"b", "a"}, {"b", "a", "a"}, {"p", "a"}, {"p", "a", "a"},
	}
	expect := [][]string{
		{"c"}, {"c", "a"},
	}

	runSample(t, paths, expect)
}
