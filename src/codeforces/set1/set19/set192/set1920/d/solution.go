package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		ops := make([][]int, n)
		for i := 0; i < n; i++ {
			ops[i] = readNNums(reader, 2)
		}
		queries := readNNums(reader, m)

		res := solve(ops, queries)
		for i, x := range res {
			if i+1 < m {
				buf.WriteString(fmt.Sprintf("%d ", x))
			} else {
				buf.WriteString(fmt.Sprintf("%d\n", x))
			}
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

const inf = 2e18

func solve(ops [][]int, queries []int) []int {
	n := len(ops)
	dp := make([]int, n+1)
	last := make([]int, n+1)
	add := true
	var pos []int
	for i, op := range ops {
		x := op[1]
		if op[0] == 1 {
			last[i+1] = x
			dp[i+1] = dp[i] + 1
		} else {
			last[i+1] = last[i]
			x++
			var tmp int = inf
			if x <= inf/dp[i] {
				tmp = x * dp[i]
			}
			dp[i+1] = tmp
			if add {
				pos = append(pos, i+1)
			}
		}
		if dp[i+1] == inf {
			add = false
		}
	}
	m := len(queries)

	ans := make([]int, m)

	for i, k := range queries {
		for j := len(pos) - 1; j >= 0; j-- {
			idx := pos[j]
			if dp[idx] > k && k > dp[idx-1] {
				if k%dp[idx-1] == 0 {
					k = dp[idx-1]
					break
				}
				k %= dp[idx-1]
			}
		}

		j := sort.Search(len(dp), func(j int) bool {
			return dp[j] >= k
		})
		ans[i] = last[j]
	}

	return ans
}
