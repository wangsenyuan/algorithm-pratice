package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	return strings.TrimSpace(s)
}

func process(reader *bufio.Reader) int {
	s := make([]string, 3)
	for i := range 3 {
		s[i] = readString(reader)
	}
	return solve(s)
}

func solve(s []string) int {
	doIt := func(arr []int) int {
		// 0, 1, 2
		res := s[arr[0]]
		for i := 1; i < 3; i++ {
			x := s[arr[i]]
			lps := kmp(x)
			var k int
			for j := 0; j < len(res); j++ {
				for k > 0 && x[k] != res[j] {
					k = lps[k-1]
				}
				if x[k] == res[j] {
					k++
				}
				if k == len(x) {
					// x already found
					break
				}
			}
			// k is for n
			res += x[k:]
		}
		return len(res)
	}

	ans := len(s[0]) + len(s[1]) + len(s[2])
	arr := make([]int, 3)
	var dfs func(i int, flag int)
	dfs = func(i int, flag int) {
		if i == 3 {
			ans = min(ans, doIt(arr))
			return
		}
		for j := 0; j < 3; j++ {
			if flag&(1<<j) == 0 {
				arr[i] = j
				dfs(i+1, flag|(1<<j))
			}
		}

	}

	dfs(0, 0)

	return ans
}

func kmp(s string) []int {
	n := len(s)
	next := make([]int, n)
	for i := 1; i < n; i++ {
		j := next[i-1]
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
	return next
}
