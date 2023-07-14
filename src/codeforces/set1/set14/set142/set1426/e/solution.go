package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	A := readNNums(reader, 3)
	B := readNNums(reader, 3)
	res := solve(A, B)
	fmt.Println(fmt.Sprintf("%d %d", res[0], res[1]))
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

func solve(A []int, B []int) []int {
	// min number to win
	// a[0] for rock
	// a[1] for scissors
	// a[2] for paper
	//n := A[0] + A[1] + A[2]
	x := winMax(A, B)

	id := make([]int, 6)
	ord := make([][]int, 6)
	for i := 0; i < 6; i++ {
		id[i] = i
	}
	ord[0] = []int{0, 0}
	ord[1] = []int{0, 2}
	ord[2] = []int{1, 1}
	ord[3] = []int{1, 0}
	ord[4] = []int{2, 2}
	ord[5] = []int{2, 1}

	process := func(a []int, b []int) int {
		for i := 0; i < 6; i++ {
			z := min(a[ord[id[i]][0]], b[ord[id[i]][1]])
			a[ord[id[i]][0]] -= z
			b[ord[id[i]][1]] -= z
		}
		ans := min(a[0], b[1]) + min(a[1], b[2]) + min(a[2], b[0])
		return ans
	}
	y := x
	a := make([]int, 6)
	b := make([]int, 6)
	for {
		copy(a, A)
		copy(b, B)
		y = min(y, process(a, b))
		if !nextPermutation(id) {
			break
		}
	}

	return []int{y, x}
}

func winMax(A []int, B []int) int {
	var res int
	x := min(A[0], B[1])

	res += x
	x = min(A[1], B[2])

	res += x
	x = min(A[2], B[0])

	res += x

	return res
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func nextPermutation(arr []int) bool {
	n := len(arr)
	i := n - 1
	for i > 0 && arr[i-1] > arr[i] {
		i--
	}
	//arr[i-1] < arr[i]
	if i == 0 {
		//unable to permutate
		return false
	}
	i--
	j := i + 1
	for k := j; k < n; k++ {
		if arr[k] > arr[i] {
			j = k
		}
	}
	arr[i], arr[j] = arr[j], arr[i]

	for a, b := i+1, n-1; a < b; a, b = a+1, b-1 {
		arr[a], arr[b] = arr[b], arr[a]
	}
	return true
}
