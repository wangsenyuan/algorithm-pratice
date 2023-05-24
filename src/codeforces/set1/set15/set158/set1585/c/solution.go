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
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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

func solve(X []int, k int) int64 {
	var pos []int
	var neg []int
	for _, x := range X {
		if x >= 0 {
			pos = append(pos, x)
		} else {
			neg = append(neg, -x)
		}
	}

	a := calc(pos, k, false) + calc(neg, k, true)
	b := calc(pos, k, true) + calc(neg, k, false)
	return min(a, b)
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}

func calc(x []int, k int, keepLast bool) int64 {
	// a positive
	n := len(x)
	if n == 0 {
		return 0
	}
	sort.Ints(x)

	// better to from backend
	res := int64(x[n-1])
	if keepLast {
		res += int64(x[n-1])
	}
	m := max(0, n-k)
	for i := m; i > 0; i -= k {
		res += 2 * int64(x[i-1])
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
