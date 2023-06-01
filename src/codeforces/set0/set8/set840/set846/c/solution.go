package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//tc := readNum(reader)
	tc := 1
	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d %d %d\n", res[0], res[1], res[2]))
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

func solve(A []int) []int {
	n := len(A)
	// n <= 5000
	// d0, d1, d2  (0 <= d0 <= d1 <= d2 <= n)
	// sum(0..d0) - sum(d0..d1) + sum(d1..d2) - sum(d2..n) 最大
	// 对于d1, 找到d0，使得 sum(0..d0) - sum(d0...d1) 最大
	// sum(d0...d1) = sum(0...d1) - sum(0...d0)
	// 2 * sum(0...d0) - sum(0...d1)

	// tmp := sum[d0] - (sum[d1] - sum[d0]) + (sum[d2] - sum[d1]) - (sum[n] - sum[d2])
	// 2 * sum[d0] - 2 * sum[d1] + 2 * sum[d2] - sum[n]
	// 所以对于d1来说，找到它前面的最大值，找到它后面的最大值即可
	sum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		sum[i+1] = sum[i] + int64(A[i])
	}

	d0 := make([]int, n+1)
	d0[0] = 0
	for i := 1; i <= n; i++ {
		d0[i] = d0[i-1]
		if sum[i] > sum[d0[i-1]] {
			d0[i] = i
		}
	}
	d2 := make([]int, n+1)
	d2[n] = n
	for i := n - 1; i >= 0; i-- {
		d2[i] = d2[i+1]
		if sum[i] > sum[d2[i+1]] {
			d2[i] = i
		}
	}
	best := 2*sum[0] - 2*sum[0] + 2*sum[0] - sum[n]
	ans := []int{0, 0, 0}
	for i := 0; i <= n; i++ {
		tmp := 2*sum[d0[i]] - 2*sum[i] + 2*sum[d2[i]] - sum[n]
		if tmp > best {
			best = tmp
			ans = []int{d0[i], i, d2[i]}
		}
	}
	return ans
}
