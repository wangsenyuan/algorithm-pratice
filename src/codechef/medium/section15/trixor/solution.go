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
		for _, cur := range res {
			for j := 0; j < len(cur); j++ {
				buf.WriteString(fmt.Sprintf("%d ", cur[j]))
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

const H = 30

func solve(A []int) [][]int {
	n := len(A)
	// n >= 6

	var res [][]int

	stack := make([]int, 4)

	for d := 0; d < H; d++ {
		var p int
		for i := 0; i < n; i++ {
			if (A[i]>>d)&1 == 1 {
				stack[p] = i
				p++
			}
			if p == 3 {
				x, y, z := stack[p-1], stack[p-2], stack[p-3]
				p = 0
				u, v, w := A[x]^A[y], A[y]^A[z], A[z]^A[x]
				res = append(res, []int{A[x], A[y], A[z]})
				A[x], A[y], A[z] = u, v, w
			}
		}
		if p == 0 {
			continue
		}
		for p < 3 {
			var x int
			for contains(stack[:p], x) {
				x++
			}
			y := x + 1
			for contains(stack[:p], y) {
				y++
			}
			z := stack[p-1]
			u, v, w := A[x]^A[y], A[y]^A[z], A[z]^A[x]
			res = append(res, []int{A[x], A[y], A[z]})
			A[x], A[y], A[z] = u, v, w
			stack[p] = y
			p++
		}
		x, y, z := stack[p-1], stack[p-2], stack[p-3]
		u, v, w := A[x]^A[y], A[y]^A[z], A[z]^A[x]
		res = append(res, []int{A[x], A[y], A[z]})
		A[x], A[y], A[z] = u, v, w
	}

	return res
}

func contains(arr []int, target int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return true
		}
	}
	return false
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
