package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k, d := readThreeNums(reader)
	s := readString(reader)[:n]
	res := solve(n, k, d, s)
	fmt.Println(res)
}

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' {
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

const Alice = "Alice"
const Bob = "Bob"
const Draw = "Draw"

func solve(n int, k int, d int, s string) string {
	if k == 1 {
		var cnt int
		for i := 0; i < n; i++ {
			if s[i] == '0' {
				cnt++
			}
		}
		// 始终有机会进行操作
		if cnt >= d {
			return Draw
		}
		if cnt&2 == 0 {
			return Bob
		}
		return Alice
	}

	dp := make(map[int][]int)
	N := 1 << n
	K := (1 << k) - 1

	var dfs func(mask int) []int

	dfs = func(mask int) []int {
		if v, ok := dp[mask]; ok {
			return v
		}
		res := []int{0, 0}
		var mn = math.MinInt32
		var mx int
		add := K
		for add <= N {
			if mask&add == 0 {
				get := dfs(mask | add)
				res[0] |= get[0] ^ 1
				if get[0] != 0 {
					mx = max(mx, get[1]+1)
				} else {
					mn = min(mn, get[1]+1)
				}
			}
			add *= 2
		}
		if res[0] != 0 {
			res[1] = mn
		} else {
			res[1] = mx
		}
		dp[mask] = res
		return res
	}
	var base int
	for i := 0; i < n; i++ {
		x := int(s[i] - '0')
		base |= x << i
	}
	ans := dfs(base)
	if ans[0] == 0 {
		if ans[1] < d {
			return Bob
		}
		return Draw
	}
	if ans[1] < d {
		return Alice
	}
	return Draw
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
