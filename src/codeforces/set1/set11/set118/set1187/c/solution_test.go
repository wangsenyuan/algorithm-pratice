package main

import (
	"bufio"
	"strings"
	"testing"
)

func runSample(t *testing.T, s string, expect bool) {
	reader := bufio.NewReader(strings.NewReader(s))
	res, facts := process(reader)

	if len(res) > 0 != expect {
		t.Fatalf("Sample %s, expect %t, but got %v", s, expect, res)
	}
	if !expect {
		return
	}
	n := len(res)

	// dp[i]表示前面那个比i大的数的下标
	stack := make([]int, n)
	var top int
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for top > 0 && res[stack[top-1]] <= res[i] {
			top--
		}
		if top == 0 {
			dp[i] = -1
		} else {
			dp[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}

	for _, cur := range facts {
		l, r := cur[1]-1, cur[2]-1
		if cur[0] == 1 {
			if dp[r] >= l {
				t.Fatalf("result %v fail fact %v", res, cur)
			}
		} else {
			// cur[0] == 0
			if dp[r] < l {
				t.Fatalf("result %v fail fact %v", res, cur)
			}
		}
	}
}

func TestSample1(t *testing.T) {
	s := `7 4
1 1 3
1 2 5
0 5 6
1 6 7
`
	expect := true
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := `4 2
1 1 4
0 2 3
`
	expect := false
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := `7 2
1 1 4
0 4 5
`
	expect := true
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := `4 3
1 3 4
0 1 2
0 2 4
`
	expect := true
	runSample(t, s, expect)
}
