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
		A := readNNums(reader, n)
		X := readNNums(reader, m)
		res := solve(A, X)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
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

func solve(A []int, X []int) []int64 {
	n := len(A)
	sum := make([]int64, n)

	var inc []int

	for i := 0; i < n; i++ {
		sum[i] = int64(A[i])
		if i > 0 {
			sum[i] += sum[i-1]
		}
		if len(inc) == 0 || sum[i] > sum[inc[len(inc)-1]] {
			inc = append(inc, i)
		}
	}

	tot := sum[n-1]

	m := len(inc)

	var check func(int64) int64

	check = func(x int64) int64 {
		if tot <= 0 && sum[inc[m-1]] < x {
			return -1
		}
		var time int64
		if sum[inc[m-1]] < x {
			time = (x - sum[inc[m-1]] + tot - 1) / tot
		}
		x -= time * tot
		i := search(m, func(i int) bool {
			return sum[inc[i]] >= x
		})

		return int64(n)*time + int64(inc[i])
	}

	ans := make([]int64, len(X))

	for i, cur := range X {
		ans[i] = check(int64(cur))
	}

	return ans
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
