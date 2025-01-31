package main

import (
	"bufio"
	"fmt"
	"os"
)

// https://atcoder.jp/contests/abc202/tasks/abc202_d

func main() {
	reader := bufio.NewReader(os.Stdin)
	first := readNInt64s(reader, 3)
	a, b := int(first[0]), int(first[1])
	k := first[2]
	res := solve(a, b, k)
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
	if len(s) == 0 || len(s) == 1 && s[0] == '\n' {
		return readNInt64s(reader, n)
	}
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

func solve(a int, b int, k int64) string {
	// 如果当前剩余 x个字符, y个字符
	// 那么有 nCr(x + y, x)中字符数
	// 如果这个数等于k, 就放置a，否则放置b
	C := make([][]int64, a+b+1)

	for i := 0; i < len(C); i++ {
		C[i] = make([]int64, a+b+1)
		C[i][0] = 1
		C[i][i] = 1
		for j := 1; j < i && j <= a; j++ {
			C[i][j] = C[i-1][j-1] + C[i-1][j]
		}
	}

	n := a + b
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		if a == 0 || a > 0 && k > C[b+a-1][a-1] {
			buf[i] = 'b'
			if a > 0 {
				k -= C[b+a-1][a-1]
			}
			b--
		} else {
			buf[i] = 'a'
			a--
		}
	}

	return string(buf)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
