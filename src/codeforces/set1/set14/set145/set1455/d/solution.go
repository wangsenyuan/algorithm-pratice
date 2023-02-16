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
		n, x := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(n, x, A)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(n int, x int, A []int) int {
	// 0,2,3,5,4, and 1
	// 0 no needs to swap in, either can't
	// 2 if there is no 5, 4, then no need to swap in either, else has to swap in
	// 3, too
	// but what if A[2] = 5? suppose the array is 0, 1, 5, 5, 4 and 2
	// also need to swap in,
	mn := A[n-1]
	F := make([]int, n)
	E := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		F[i] = F[i+1]
		if A[i] > mn {
			F[i]++
			E[i]++
		} else {
			mn = A[i]
		}
	}

	var res int

	for i := 0; i < n; i++ {
		if i < n-1 && F[i+1] > 0 {
			// has to replace x with A[i] to prepare A[i] for later usage
			if A[i] > x {
				res++
				A[i], x = x, A[i]
				continue
			}
		}

		if E[i] > 0 {
			if A[i] < x {
				return -1
			}
			res++
			A[i], x = x, A[i]
		}
	}

	for i := 1; i < n; i++ {
		if A[i] < A[i-1] {
			return -1
		}
	}

	return res
}

type Pair struct {
	first, second int
}
