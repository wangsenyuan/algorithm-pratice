package p797

import (
	"testing"
	"bytes"
	"strconv"
)

func runSample(t *testing.T, graph [][]int, expect [][]int) {
	res := allPathsSourceTarget(graph)

	if len(res) != len(expect) {
		t.Errorf("sample %v, expect %d result, but got %d result", graph, len(graph), len(res))
		return
	}

	set := make(map[string]bool)

	for _, path := range expect {
		set[toString(path)] = true
	}

	for _, path := range res {
		s := toString(path)
		if set[s] == false {
			t.Errorf("sample %v, result %s, is not expected", graph, s)
		}
	}
}

func toString(slice []int) string {
	var buf bytes.Buffer

	buf.WriteString(strconv.Itoa(slice[0]))
	for i := 1; i < len(slice); i++ {
		buf.WriteString("->")
		buf.WriteString(strconv.Itoa(slice[i]))
	}
	return buf.String()
}

func TestSample1(t *testing.T) {
	graph := [][]int{
		{1, 2}, {3}, {3}, {},
	}
	expect := [][]int{{0, 1, 3}, {0, 2, 3}}
	runSample(t, graph, expect)
}
