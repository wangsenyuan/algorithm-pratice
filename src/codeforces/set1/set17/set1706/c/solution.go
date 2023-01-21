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
		n := readNum(reader)
		H := readNNums(reader, n)
		res := solve(H)
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(H []int) int64 {
	n := len(H)

	get := func(i int) int64 {
		cur := max3(H[i-1], H[i], H[i+1])

		if cur == H[i-1] {
			cur++
		}
		if cur == H[i+1] {
			cur++
		}
		return int64(cur - H[i])
	}

	// 0 1, 0, 1
	var res int64
	if n%2 == 1 {
		for i := 1; i+1 < n; i += 2 {
			res += get(i)
		}
	} else {
		// 0, 1, 2, 3, 4, 5
		// could be 1, 3 or 2, 4
		dp := make([]int64, n+1)
		dp[n] = 0
		dp[n-1] = 0

		for i := n - 2; i > 0; i-- {
			dp[i] = dp[i+2] + get(i)
		}

		res = min(dp[1], dp[2])

		fp := make([]int64, 2)
		fp[0] = 1
		fp[1] = get(1)

		for i := 2; i < n-2; i++ {
			// if keep i not change
			if i%2 == 0 {
				tmp := fp[1] + min(dp[i+1], dp[i+2])
				res = min(res, tmp)
			} else {
				tmp := min(fp[0], fp[1]) + dp[i+1]
				res = min(res, tmp)
			}
			fp[i%2] += get(i)
		}
	}

	return res
}

func max3(a, b, c int) int {
	if a >= b && a >= c {
		return a
	}
	if b >= a && b >= c {
		return b
	}
	return c
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
