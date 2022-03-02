package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func solve(A []int, K int64) int64 {
	n := len(A)
	if n == 1 {
		return max(0, int64(A[0])-K)
	}
	sort.Ints(A)
	var sum int64

	for _, num := range A {
		sum += int64(num)
	}

	if sum <= K {
		return 0
	}

	var best int64 = sum - K

	sum -= int64(A[0])

	for i := n - 1; i > 0; i-- {
		// decrease first to some m
		// and set A[i...]  to A[0]
		// (n - i + 1) * x + sum <= K
		sum -= int64(A[i])
		x := (K - sum) / int64(n-i+1)
		if x*int64(n-i+1) > (K - sum) {
			x--
		}
		a := max(0, int64(A[0])-x)
		best = min(best, a+int64(n-i))
	}

	return best
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		var n int
		var K uint64
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &n) + 1
		readUint64(s, pos, &K)
		A := readNNums(reader, n)
		res := solve(A, int64(K))
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
