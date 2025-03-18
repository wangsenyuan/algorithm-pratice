package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _, _ := process(reader)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func readInt(bytes []byte, from int, val *int) int {
	i := from
	sign := 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	tmp := 0
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(reader *bufio.Reader) (a int) {
	bs, _ := reader.ReadBytes('\n')
	readInt(bs, 0, &a)
	return
}

func readTwoNums(reader *bufio.Reader) (a int, b int) {
	res := readNNums(reader, 2)
	a, b = res[0], res[1]
	return
}

func readThreeNums(reader *bufio.Reader) (a int, b int, c int) {
	res := readNNums(reader, 3)
	a, b, c = res[0], res[1], res[2]
	return
}

func readNNums(reader *bufio.Reader, n int) []int {
	res := make([]int, n)
	x := 0
	bs, _ := reader.ReadBytes('\n')
	for i := 0; i < n; i++ {
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') && bs[x] != '-' {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func process(reader *bufio.Reader) (res string, name string, p string) {
	n, k := readTwoNums(reader)
	name = readString(reader)
	p = readString(reader)
	res = solve(n, k, name, p)
	return
}

const fail = "No solution"

func solve(n int, k int, name string, p string) string {
	lps := kmp(name)
	m := len(name)
	good := make([]bool, m+1)
	for i := m; i > 0; i = lps[i-1] {
		good[i] = true
	}

	res := make([]byte, n)
	for i := range n {
		res[i] = '-'
	}
	prev := 0
	for i := range len(p) {
		if p[i] == '1' {
			if prev > i && !good[prev-i] {
				return fail
			}
			copy(res[i:], name)
			prev = i + m
		}
	}
	// 已经满足了第一个要求
	vis := make([][]bool, n)
	for i := range n {
		vis[i] = make([]bool, m+1)
	}

	var dfs func(i int, j int) bool

	dfs = func(i int, j int) bool {
		if i == n {
			return true
		}
		if vis[i][j] {
			return false
		}
		vis[i][j] = true
		if res[i] != '-' {
			for j > 0 && res[i] != name[j] {
				j = lps[j-1]
			}
			if res[i] == name[j] {
				j++
			}
			if j == m {
				if p[i-j+1] == '0' {
					return false
				}
				j = lps[j-1]
			}
			return dfs(i+1, j)
		}

		// res[i] == '-'
		for x := range k {
			res[i] = byte('a' + x)
			nj := j
			for nj > 0 && res[i] != name[nj] {
				nj = lps[nj-1]
			}
			if res[i] == name[nj] {
				nj++
			}
			if nj == m && p[i-nj+1] == '0' {
				// bad
				continue
			}
			if nj == m {
				nj = lps[nj-1]
			}
			if dfs(i+1, nj) {
				return true
			}
		}
		return false
	}

	if !dfs(0, 0) {
		return fail
	}
	return string(res)
}

func kmp(s string) []int {
	n := len(s)
	lps := make([]int, n)
	for i := 1; i < n; i++ {
		j := lps[i-1]
		for j > 0 && s[i] != s[j] {
			j = lps[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		lps[i] = j
	}
	return lps
}
