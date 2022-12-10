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
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
		if !res {
			buf.WriteString("-1\n")
		} else {
			for i := 0; i < n; i++ {
				buf.WriteString(fmt.Sprintf("%d ", A[i]))
			}
			buf.WriteByte('\n')
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

func solve(A []int) bool {
	if len(A) <= 2 {
		return true
	}
	// a < b > c
	// if a <= b, then we need to find a c, a > c < b or a < c > b
	// 相等的不能放在一起,
	// 只能是一个锯齿形的，
	// a < b > c < d > e ...
	sort.Ints(A)

	if zigZag(A) {
		return true
	}
	reverse(A)

	if zigZag(A) {
		return true
	}

	return false
}

func zigZag(A []int) bool {
	n := len(A)
	B := make([]int, n)
	j := (n + 1) / 2
	B[0] = A[0]
	k := 1
	for i := 1; i < (n+1)/2; i++ {
		B[k] = A[j]
		k++
		j++
		B[k] = A[i]
		k++
	}
	if k < n {
		B[k] = A[j]
	}
	if intersting(B) {
		copy(A, B)
		return true
	}
	return false
}

func intersting(arr []int) bool {
	for i := 1; i+1 < len(arr); i++ {
		if arr[i-1] <= arr[i] && arr[i] <= arr[i+1] {
			return false
		}
		if arr[i-1] >= arr[i] && arr[i] >= arr[i+1] {
			return false
		}
	}
	return true
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
