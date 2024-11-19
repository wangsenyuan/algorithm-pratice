package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect string) {
	res, pages, k := process(bufio.NewReader(strings.NewReader(s)))

	if res == expect {
		return
	}
	if res == "IMPOSSIBLE" {
		t.Fatalf("Sample %s expect %s, but got IMPOSSIBLE", s, expect)
	}
	ord := make([]int, 26)
	for i := 0; i < 26; i++ {
		ord[i] = -1
	}
	for i, x := range []byte(res) {
		u := int(x - 'a')
		if ord[u] != -1 {
			t.Fatalf("Sample result %v, has duplicates", res)
		}
		ord[u] = i
	}
	book := readAsBook(pages, k)

	var arr []string
	for _, cur := range book {
		arr = append(arr, cur.lines...)
	}

	id := func(x byte) int {
		return int(x - 'a')
	}
	for i := 0; i+1 < len(arr); i++ {
		var j int
		for j < len(arr[i]) && j < len(arr[i+1]) && arr[i][j] == arr[i+1][j] {
			j++
		}
		if j == len(arr[i]) {
			continue
		}
		if j == len(arr[i+1]) || ord[id(arr[i][j])] < 0 || ord[id(arr[i][j])] >= ord[id(arr[i+1][j])] {
			t.Fatalf("Sample result %v, not correct", res)
		}
	}
}

func TestSample1(t *testing.T) {
	s := `3 3
2
b
b
bbac
0
a
aca
acba
1
ab
c
ccb`
	expect := "acb"
	runSample(t, s, expect)
}
