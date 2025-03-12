package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
}
func process(reader *bufio.Reader) string {
	box := make([]string, 5)
	for i := range 5 {
		box[i] = readString(reader)
	}
	return solve(box)
}

func solve(box []string) string {
	dp := make([]int, 1<<19)

	for i := range len(dp) {
		dp[i] = -1
	}
	dp[0] = 0

	// 位置进行编号
	// 012
	// 3456
	// 789AB
	// CDEF
	// HIJ
	var lines [][]int
	// 怎么把这个lines得到呢？直接看出来当然可以，但是太笨了
	lines = append(lines, []int{0, 1, 2})
	lines = append(lines, []int{3, 4, 5, 6})
	lines = append(lines, []int{7, 8, 9, 10, 11})
	lines = append(lines, []int{12, 13, 14, 15})
	lines = append(lines, []int{16, 17, 18})
	lines = append(lines, []int{0, 3, 7})
	lines = append(lines, []int{1, 4, 8, 12})
	lines = append(lines, []int{2, 5, 9, 13, 16})
	lines = append(lines, []int{6, 10, 14, 17})
	lines = append(lines, []int{11, 15, 18})
	lines = append(lines, []int{2, 6, 11})
	lines = append(lines, []int{1, 5, 10, 15})
	lines = append(lines, []int{0, 4, 9, 14, 18})
	lines = append(lines, []int{3, 8, 13, 17})
	lines = append(lines, []int{7, 12, 16})

	get := func(state int) []int {
		var res []int
		for _, line := range lines {
			for i := 0; i < len(line); i++ {
				if (state>>line[i])&1 == 0 {
					continue
				}
				tmp := state
				for j := i; j < len(line); j++ {
					if (state>>line[j])&1 == 0 {
						break
					}
					// 这一段都是有值的
					tmp ^= 1 << line[j]
					res = append(res, tmp)
				}
			}
		}
		return sortAndUnique(res)
	}

	var dfs func(state int) (ans bool)

	dfs = func(state int) (ans bool) {
		if dp[state] >= 0 {
			return dp[state] > 0
		}

		defer func() {
			if ans {
				dp[state] = 1
			} else {
				dp[state] = 0
			}
		}()

		sub := get(state)
		ans = false
		for _, cur := range sub {
			if !dfs(cur) {
				// 当前状态下，只要有一个失败的子状态，那么当前状态就是获胜的状态
				ans = true
				break
			}
		}
		return
	}
	var cur int
	for i, row := range box {
		row = strings.ReplaceAll(row, " ", "")
		for k, x := range []byte(row) {
			if x == 'O' {
				cur ^= 1 << lines[i][k]
			}
		}
	}
	if dfs(cur) {
		return "Karlsson"
	}
	return "Lillebror"
}

func sortAndUnique(arr []int) []int {
	sort.Ints(arr)
	var n int
	for i := 1; i <= len(arr); i++ {
		if i == len(arr) || arr[i-1] < arr[i] {
			arr[n] = arr[i-1]
			n++
		}
	}
	return arr[:n]
}
