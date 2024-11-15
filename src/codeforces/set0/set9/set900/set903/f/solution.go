package main

import (
	"bufio"
	"fmt"
	"os"
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

func process(r *bufio.Reader) int {
	n := readNum(r)
	a := readNNums(r, 4)
	mat := make([]string, 4)
	for i := range 4 {
		mat[i] = readString(r)[:n]
	}
	return solve(a, mat)
}

const inf = 1 << 50

func solve(a []int, s []string) int {
	n := len(s[0])

	memo := make([][4][1 << 4][1 << 4][1 << 4]int, n)
	var dfs func(int, int, int, int, int) int
	dfs = func(j, i, cur, pre, pre2 int) int {
		if j < 0 {
			if pre > 0 || pre2 > 0 {
				return a[3]
			}
			return 0
		}
		if i > 3 {
			if pre2 > 0 {
				return dfs(j-2, 0, 0, 0, 0) + a[3]
			}
			return dfs(j-1, 0, 0, cur, pre)
		}
		p := &memo[j][i][cur][pre][pre2]
		if *p > 0 {
			return *p - 1
		}
		res := min(
			dfs(j, i+1, cur<<1|int(s[i][j]>>2&1^1), pre, pre2),
			dfs(j, i+1, cur<<1, pre, pre2)+a[0],
			dfs(j, i+2, cur<<min(4-i, 2), pre&^(3<<max(2-i, 0)), pre2)+a[1],
			dfs(j, i+3, cur<<min(4-i, 3), pre&^(7<<max(1-i, 0)), pre2&^(7<<max(1-i, 0)))+a[2],
		)
		*p = res + 1
		return res
	}

	return dfs(n-1, 0, 0, 0, 0)
}
