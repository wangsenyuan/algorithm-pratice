package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	H := readNNums(reader, n)
	fmt.Println(solve(n, H))
}

func solve(n int, H []int) int64 {
	hills := make([]int, n+1)

	var pivot int
	for i := 1; i < n; i++ {
		if H[i] > H[pivot] {
			pivot = i
		}
	}

	copy(hills, H[pivot:])
	copy(hills[n-pivot:], H[0:])
	hills[n] = H[pivot]
	//n++

	C := make([]int, n+1)
	R := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		R[i] = i + 1
		for R[i] < n && hills[i] > hills[R[i]] {
			R[i] = R[R[i]]
		}
		if R[i] < n && hills[i] == hills[R[i]] {
			C[i] = C[R[i]] + 1
			R[i] = R[R[i]]
		}
	}

	L := make([]int, n+1)
	for i := 1; i <= n; i++ {
		L[i] = i - 1
		for L[i] > 0 && hills[i] >= hills[L[i]] {
			L[i] = L[L[i]]
		}
	}

	var res int64

	for i := 0; i < n; i++ {
		res += int64(C[i])
		if hills[i] < hills[0] {
			if L[i] == 0 && R[i] == n {
				res++
			} else {
				res += 2
			}
		}
	}

	return res
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
