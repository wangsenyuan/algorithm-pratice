package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		color := readNNums(reader, n)
		res := solve(m, color)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(k int, c []int) int {
	n := len(c)

	pos := make([][]int, k+1)

	for i := 0; i < n; i++ {
		x := c[i]
		if pos[x] == nil {
			pos[x] = make([]int, 0, 1)
		}
		pos[x] = append(pos[x], i)
	}

	checkColor := func(j int, md int) bool {
		ps := pos[j]

		prev := -1
		var change bool

		for _, i := range ps {
			gap := i - prev - 1
			if gap > md {
				if change || gap/2 > md {
					return false
				}
				change = true
			}
			prev = i
		}
		return (n-prev-1 <= md) || (n-prev-1)/2 <= md && !change
	}

	check := func(x int) bool {
		// 是否能够最大间距，且最多改变一次某个plank的颜色，从1到n
		// dp[i][0] 表示从i开始，且没有修改时，是否能够到达n
		// dp[i][1] 表示从i开始，且已经修改过一次后，是否能够到达n
		// dp[i][1] = dp[next[i][1] || some j, dp[j][0], 且 j - i - 1 <= x
		for i := 1; i <= k; i++ {
			if checkColor(i, x) {
				return true
			}
		}
		return false
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
