package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		A := readString(reader)[:n]
		B := readString(reader)[:n]
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, B, Q)
		for _, r := range res {
			if r {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}

	fmt.Print(buf.String())
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func solve(A, B string, Q [][]int) []bool {
	// 对于[l..r]
	// A[l..r] => B[l...r]
	// diff[i] = (B[i] - A[i]) % 26
	//  考虑将 l...r 变成一致的一种操作方式是
	// 先对l进行处理, 这时候要进行 diff[l]次
	// 同时 l+1 也需要进行 diff[l]次操作，但是 l + 2 不进行操作
	//  l + 1, 现在它变成了需要进行 change[l+1] = diff[l+1] - diff[l]
	//  l + 2 需要进行 change[l + 2] = diff[l+2] - (diff[l+1] - diff[l] )
	//  l + i ... change[l+i] =   diff[l + i] - change[l + i - 1]
	//  如果对于 r, 如果 change[r-1] == diff[r]
	// 那么 [l...r] 是可以变成一致的
	// 怎么快速的计算 change[l+i]
	// change[l+i] = diff[l+i] - change[l + i - 1] = diff[l+i] - (diff[l+i-1] - change[l + i - 2])
	//         = sum of diff at even distance - sum of diff at odd distance
	n := len(A)
	sum := make([]int, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i]
		x := byteDiff(B[i], A[i])
		if i%2 == 0 {
			sum[i+1] += x
		} else {
			sum[i+1] += 26 - x
		}
		sum[i+1] %= 26
	}

	check := func(l int, r int) bool {
		if l == r {
			// can't make any change
			return A[l-1] == B[l-1]
		}
		return (sum[r]-sum[l-1])%26 == 0
	}

	res := make([]bool, len(Q))

	for i, cur := range Q {
		res[i] = check(cur[0], cur[1])
	}

	return res
}

func byteDiff(x, y byte) int {
	return int(x) - int(y)
}
