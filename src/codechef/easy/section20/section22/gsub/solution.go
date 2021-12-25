package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

const N = 100010

var arr [N]int
var bak [N]int

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		A := readNNums(reader, n)
		solver := NewSolver(n, A)
		for q > 0 {
			q--
			x, y := readTwoNums(reader)
			res := solver.Update(x, y)
			buf.WriteString(fmt.Sprintf("%d\n", res))
		}
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
	res := bak[:n]
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

type Solver struct {
	arr []int
	cur int
	n   int
}

func NewSolver(n int, A []int) Solver {
	var ans = 2
	arr := make([]int, n+2)
	copy(arr[1:], A)
	arr[0] = -1
	arr[n+1] = -1

	for i := 1; i <= n; i++ {
		if arr[i] != arr[i-1] {
			ans++
		}
	}

	return Solver{arr, ans, n}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func (solver *Solver) Update(x int, y int) int {
	arr := solver.arr
	cur := solver.cur

	if arr[x] == y {
		// no change
		return cur - 2
	}

	if arr[x] == arr[x-1] && arr[x] == arr[x+1] {
		cur += 2
	} else if (arr[x] != arr[x-1]) && (arr[x] == arr[x+1]) {
		cur++
		if y == arr[x-1] {
			cur--
		}
	} else if arr[x] == arr[x-1] && arr[x] != arr[x+1] {
		cur++
		if y == arr[x+1] {
			cur--
		}
	} else {
		// arr[x] != arr[x-1] && arr[x] != arr[x+1]
		if y == arr[x-1] {
			cur--
		}
		if y == arr[x+1] {
			cur--
		}
	}

	arr[x] = y

	solver.cur = cur
	return cur - 2
}
