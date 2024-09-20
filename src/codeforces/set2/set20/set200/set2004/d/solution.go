package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		s := readString(reader)
		cities := strings.Split(s, " ")
		queries := make([][]int, m)
		for i := 0; i < m; i++ {
			queries[i] = readNNums(reader, 2)
		}

		res := solve(n, cities, queries)

		for _, x := range res {
			buf.WriteString(fmt.Sprintf("%d\n", x))
		}
	}

	fmt.Print(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

const inf = 1 << 50

func solve(n int, cities []string, queries [][]int) []int {

	get := func(x byte) int {
		if x == 'R' {
			return 0
		}
		if x == 'G' {
			return 1
		}
		if x == 'B' {
			return 2
		}
		return 3
	}

	getPortal := func(i int) (int, int) {
		x := get(cities[i][0])
		y := get(cities[i][1])
		return min(x, y), max(x, y)
	}

	// 4 * 3 / 2
	prev := make([]int, 16)
	next := make([]int, 16)
	for i := 0; i < 16; i++ {
		prev[i] = -1
		next[i] = -1
	}

	// 还必须知道每个city的最短到高速的最短距离
	dp := make([][]int, n)
	fp := make([][]int, n)
	for i := range cities {
		dp[i] = make([]int, 16)
		for j := 0; j < 16; j++ {
			dp[i][j] = -1
		}
		x, y := getPortal(i)

		for j := 0; j < 16; j++ {
			if prev[j] >= 0 {
				u, v := j/4, j%4
				if u == x || u == y || v == x || v == y {
					dp[i][j] = prev[j]
				}
			}
		}
		prev[x*4+y] = i
		prev[y*4+x] = i
	}

	for i := n - 1; i >= 0; i-- {
		fp[i] = make([]int, 16)
		for j := 0; j < 16; j++ {
			fp[i][j] = n
		}
		x, y := getPortal(i)
		for j := 0; j < 16; j++ {
			if next[j] >= 0 {
				u, v := j/4, j%4
				if u == x || u == y || v == x || v == y {
					fp[i][j] = next[j]
				}
			}
		}
		next[x*4+y] = i
		next[y*4+x] = i
	}

	ans := make([]int, len(queries))

	for i, cur := range queries {
		x, y := cur[0]-1, cur[1]-1
		if x == y {
			ans[i] = 0
			continue
		}
		x, y = min(x, y), max(x, y)
		a, b := getPortal(x)
		c, d := getPortal(y)

		if a == c || a == d || b == c || b == d {
			ans[i] = y - x
		} else {
			// 必须到(b, c)
			ans[i] = inf
			for _, u := range []int{a, b} {
				for _, v := range []int{c, d} {
					j := u*4 + v
					if dp[y][j] > x || fp[x][j] < y {
						ans[i] = y - x
						// 可以在中间有一个踏板
						break
					}
					if dp[y][j] >= 0 {
						ans[i] = min(ans[i], y-dp[y][j]+x-dp[y][j])
					}
					if fp[x][j] < n {
						ans[i] = min(ans[i], fp[x][j]-y+fp[x][j]-x)
					}
				}
			}
			if ans[i] == inf {
				ans[i] = -1
			}
		}
	}

	return ans
}

func abs(num int) int {
	return max(num, -num)
}
