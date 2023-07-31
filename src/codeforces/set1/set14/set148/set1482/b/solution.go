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
		n := readNum(reader)
		nums := readNNums(reader, n)
		res := solve(nums)
		for i := 0; i < len(res); i++ {
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

func solve(A []int) []int {
	n := len(A)
	// A[0] = s % m
	// s 没有任何关系
	// A[i] = (A[i-1] + c) % m
	// A[i] < m
	// 如果 A[i-1] + c < m => A[i] = A[i-1] + c
	// 也就是说 可能存在连续的连续的两个数 A[i] - A[i-1] = c
	// 如果 A[i] - A[i] = 固定的数
	if n <= 2 {
		return []int{0}
	}
	diff := make(map[int]int)

	for i := 1; i < n; i++ {
		diff[A[i]-A[i-1]]++
	}
	// A[i] = A[i-1] + c， 或者 A[i] = A[i-1] + c - m
	// diff = c or c - m
	if len(diff) > 2 {
		return []int{-1}
	}
	if len(diff) == 1 {
		return []int{0}
	}
	// diff = 2
	var arr []int
	for k := range diff {
		arr = append(arr, k)
	}
	a := arr[0]
	b := arr[1]
	// a = c, b = c - m
	m := a - b
	if m < 0 {
		m = b - a
	}
	for i := 0; i < n; i++ {
		if A[i] >= m {
			return []int{-1}
		}
	}
	if a > 0 {
		return []int{m, a}
	}
	return []int{m, b}
}
