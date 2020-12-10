package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	ans := solve(n, A)
	var buf bytes.Buffer

	for i := 0; i < len(ans); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", ans[i].left, ans[i].right))
		for j := 0; j < len(ans[i].nums); j++ {
			buf.WriteString(fmt.Sprintf("%d ", ans[i].nums[j]))
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

func solve(n int, A []int) [3]Operation {
	var res [3]Operation
	if n == 1 {
		res[0] = Operation{1, 1, []int64{0}}
		res[1] = Operation{1, 1, []int64{0}}
		res[2] = Operation{1, 1, []int64{int64(-A[0])}}
		return res
	}
	res[0] = Operation{1, 1, []int64{int64(-A[0])}}
	arr := make([]int64, n)
	for i := 1; i < n; i++ {
		arr[i] = -int64(n) * int64(A[i])
	}
	res[1] = NewOperation(1, n, arr)
	arr = make([]int64, n-1)
	for i := 1; i < n; i++ {
		arr[i-1] = int64(n-1) * int64(A[i])
	}
	res[2] = NewOperation(2, n, arr)

	return res
}

type Operation struct {
	left, right int
	nums        []int64
}

func NewOperation(l, r int, arr []int64) Operation {
	tmp := make([]int64, len(arr))
	copy(tmp, arr)
	return Operation{l, r, tmp}
}
