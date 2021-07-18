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
	n := readNum(reader)
	A := readNNums(reader, n)
	B := readNNums(reader, n)
	res := solve(n, A, B)

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

const MOD = 1000000007

func solve(n int, A, B []int) []int {
	for i := 0; i < n; i++ {
		A[i] += i + 1
		B[i] += i + 1
	}
	L1, R1 := getMaxPosition(n, A)
	L2, R2 := getMaxPosition(n, B)

	f := func(k int, A []int, L, R []int) int64 {
		var sum int64
		for i := 0; i < n; i++ {
			cnt := max(0, min(i, R[i]-k)+1-max(L[i]+1, i-k+1))
			sum += int64(A[i]) * int64(cnt) % MOD
		}
		return sum % MOD
	}
	res := make([]int, n)
	for k := 1; k <= n; k++ {
		a := f(k, A, L1, R1)
		b := f(k, B, L2, R2)
		res[k-1] = int(a * b % MOD)
	}
	return res
}

func getMaxPosition(n int, A []int) ([]int, []int) {
	// A[L[i]] >= A[i]
	L := make([]int, n)
	// A[R[i]] > A[i]
	R := make([]int, n)

	stack := make([]int, n)
	var p int

	for i := 0; i < n; i++ {
		L[i] = -1
		for p > 0 && A[stack[p-1]] < A[i] {
			p--
		}
		if p > 0 {
			// A[stack[p-1]] >= A[i]
			L[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}

	p = 0
	for i := n - 1; i >= 0; i-- {
		R[i] = n
		for p > 0 && A[stack[p-1]] <= A[i] {
			p--
		}
		if p > 0 {
			// A[stack[p-1] > A[i]
			R[i] = stack[p-1]
		}
		stack[p] = i
		p++
	}
	return L, R
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
