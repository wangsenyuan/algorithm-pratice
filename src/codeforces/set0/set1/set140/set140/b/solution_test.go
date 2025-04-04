package main

import (
	"bufio"
	"slices"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect []int) {
	res, friends, alex := process(bufio.NewReader(strings.NewReader(s)))

	n := len(res)

	pos := make([]int, n)
	for i, v := range alex {
		pos[v-1] = i
	}

	type pair struct {
		first  int
		second int
	}

	ps := make([]pair, n)

	check := func(arr []int) int {
		for i := 0; i < n; i++ {
			ps[i] = pair{arr[i] - 1, i}
		}
		slices.SortFunc(ps, func(a, b pair) int {
			return a.first - b.first
		})

		var sum int
		first, second := -1, -1
		for i, j := 0, 0; i < n; i++ {
			if first == -1 || pos[i] < pos[first] {
				second = first
				first = i
			} else if second == -1 || pos[i] < pos[second] {
				second = i
			}
			for j < n && ps[j].first == i {
				k := ps[j].second
				card := first
				if k == card {
					card = second
				}
				if card < 0 {
					t.Fatalf("Sample result %v, not correct, alex send no card to friends %d", arr, k)
				}

				for u, v := range friends[k] {
					v--
					if v == card {
						sum += u
						break
					}
				}

				j++
			}
		}

		return sum
	}

	if check(expect) != check(res) {
		t.Fatalf("Sample result %v, not correct, expect %v", res, expect)
	}
}

func TestSample1(t *testing.T) {
	s := `4
1 2 3 4
4 1 3 2
4 3 1 2
3 4 2 1
3 1 2 4
`
	expect := []int{2, 1, 1, 4}
	runSample(t, s, expect)
}
