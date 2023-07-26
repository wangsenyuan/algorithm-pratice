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
		A := readNNums(reader, n)
		res := solve(A)
		buf.WriteString(fmt.Sprintf("%d\n", len(res)))
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(A []int) []int {
	n := len(A)

	// 如果 sum(A) >= n / 2
	// 且 sum(A) % 2 == 0
	// 只需要保留所有的1即可
	// 如果 sum(A) % 2 == 1， 那么需要删除掉一个1，且增加两个0
	var sum int
	for _, a := range A {
		sum += a
	}
	if n == 2 {
		if sum == 0 || sum == 2 {
			return A
		}
		return []int{0}
	}

	h := n / 2
	if sum != h || h&1 == 0 {
		//eight all 0 or 1
		h += h & 1
		res := make([]int, h)
		if sum >= h {
			for i := 0; i < h; i++ {
				res[i] = 1
			}
		}
		return res
	}
	// half 1, half 0, and h is odd
	// 比如[1,0,1,0,1,0]
	res := make([]int, h+1)
	var x int
	if A[0] == A[1] || A[0] == A[2] {
		x = A[0]
	} else {
		// A[1] == A[2]
		x = A[1]
	}
	res[0] = x
	res[1] = x
	for i := 2; i <= h; i++ {
		res[i] = 1 - x
	}

	return res
}
