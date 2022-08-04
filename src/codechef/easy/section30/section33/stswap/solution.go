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

	n, k := readTwoNums(reader)
	A := readNNums(reader, n)
	res := solve(A, k)

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

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

func solve(A []int, k int) []int {
	n := len(A)
	k++
	if k >= n {
		return A
	}
	// A[0] swap with A[i] where i >= k
	// and A[0] should be placed at position j, where j < k
	// then if there is some element x, (x - j) >= k, then A[0] can go to position j, by x
	// so if x is position n - 1
	// for any A[i], the position i can be placed range is [i+k, n - 1], i + k <= n - 1
	// and [0, n - 1 - k] when n - 1 - i >= k
	// 0, 1, 2, 3, 4, 5
	// k = 3
	// or [0... i - k] when i >= k, and [k, n-1]
	// if k <= n - 1 - k, result is the sorted one
	if k <= n-1-k {
		sort.Ints(A)
		return A
	}
	// k > n - 1 - k
	B := make([]int, 0, n)
	for i := 0; i <= n-1-k; i++ {
		B = append(B, A[i])
	}
	for i := k; i < n; i++ {
		B = append(B, A[i])
	}
	sort.Ints(B)

	for i := 0; i <= n-1-k; i++ {
		A[i] = B[i]
	}

	for j, u := n-k, k; j < len(B); j, u = j+1, u+1 {
		A[u] = B[j]
	}
	return A
}
