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

	tc := 1
	for tc > 0 {
		tc--
		n := readNum(reader)
		C := readNNums(reader, n)
		buf.WriteString(fmt.Sprintf("%d\n", solve(n, C)))
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

func solve(n int, C []int) int {
	L, R := make([]int, n), make([]int, n)
	F := make([]int, n)

	for i := 0; i < n; i++ {
		L[i] = -1
	}

	for i := 0; i < n; i++ {
		x := C[i] - 1
		if L[x] < 0 {
			L[x] = i
		}
		R[x] = i
		F[x]++
	}

	CF := make([]int, n)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		dp[i] = dp[i+1]
		x := C[i] - 1
		CF[x]++
		if i == L[x] {
			// color x都留在原地，其他的都被移到到后端
			dp[i] = max(dp[i], dp[R[x]+1]+F[x])
		} else {
			dp[i] = max(dp[i], CF[x])
		}
	}

	return n - dp[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
