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
		arr := readNNums(reader, n*m)
		res := solve(n, m, arr)
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

func solve(n int, m int, arr []int) int {
	sort.Ints(arr)
	// 对于所有的子矩形，找出其中的最大值，- 最小值 求和
	// 只需要考虑最大值 + 两个最小值即可
	ans := solve1(n, m, arr)
	ans = max(ans, solve2(n, m, arr))

	return ans
}

func solve2(n int, m int, arr []int) int {
	// (1, 1) take the min value
	a, b, c := arr[0], arr[n*m-2], arr[n*m-1]

	res := -a * (n*m - 1)
	// b <= c
	if n > m {
		b, c = c, b
	}
	if b <= c {
		res += b * (n - 1)
		res += c * n * (m - 1)
	} else {
		res += c * (m - 1)
		res += b * (n - 1) * m
	}
	return res
}

func solve1(n int, m int, arr []int) int {
	a, b, c := arr[0], arr[1], arr[m*n-1]
	res := c * (n*m - 1)
	a, b = min(a, b), max(a, b)
	if n < m {
		a, b = max(a, b), min(a, b)
	}
	// a is below of first, b is right of first
	if a <= b {
		res -= b * (m - 1)
		res -= a * (n - 1) * m
	} else {
		res -= a * (n - 1)
		res -= b * n * (m - 1)
	}
	return res
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
