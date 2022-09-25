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
		Q := make([][]int, m)
		for i := 0; i < m; i++ {
			Q[i] = readNNums(reader, 2)
		}
		res := solve(A, Q)
		for i := 0; i < m; i++ {
			buf.WriteString(fmt.Sprintf("%d\n", res[i]))
		}
	}
	fmt.Print(buf.String())
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(A []int, Q [][]int) []int {
	// if A[i] > A[i+1], have to apply one operation
	// 因为后面的数字增加的比前面的多， 所以始终选择r = n是个更好的策略，且不会影响前面的结果
	n := len(A)
	m := n - 1
	arr := make([]int, 2*m)

	for i := m; i < 2*m; i++ {
		arr[i] = A[i-m] - A[i-m+1]
	}

	for i := m - 1; i > 0; i-- {
		arr[i] = max(arr[i*2], arr[i*2+1])
	}

	set := func(p int, v int) {
		p += m
		arr[p] = v
		for p > 1 {
			arr[p>>1] = max(arr[p], arr[p^1])
			p >>= 1
		}
	}

	ans := make([]int, len(Q))

	for i, qr := range Q {
		j, x := qr[0], qr[1]
		j--
		if j > 0 {
			set(j-1, A[j-1]-x)
		}
		if j+1 < n {
			set(j, x-A[j+1])
		}
		ans[i] = max(0, arr[1])
		A[j] = x
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
