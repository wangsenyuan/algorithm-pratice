package main

import (
	"bufio"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	C := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &C[i])
	}
	res := solve(C)
	fmt.Println(res)
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

func solve(C []int) int {
	// if pick i as k (0 < k < n - 1), what is the min value
	n := len(C)
	// R[i] = C[i...n] count C[i] > 0
	R := make([]int, n)
	if C[n-1] > 0 {
		R[n-1]++
	}

	for k := n - 2; k >= 0; k-- {
		R[k] = R[k+1]
		if C[k] > 0 {
			R[k]++
		}
	}
	// C[i] < 0 until k
	var cnt int
	if C[0] < 0 {
		cnt++
	}
	res := n
	for k := 1; k <= n-1; k++ {
		tmp := k - cnt + (n - k - R[k])
		res = min(res, tmp)
		if C[k] < 0 {
			cnt++
		}
	}
	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
