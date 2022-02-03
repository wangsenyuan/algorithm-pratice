package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	div := func(x, y, d int) bool {
		x++
		y++
		fmt.Printf("1 %d %d %d", x, y, d)
		s, _ := reader.ReadString('\n')
		return s[0] == 'Y'
	}

	less := func(x, y int) bool {
		x++
		y++
		fmt.Printf("2 %d %d", x, y)
		s, _ := reader.ReadString('\n')
		return s[0] == 'Y'
	}

	tc := readNum(reader)
	for tc > 0 {
		tc--
		n := readNum(reader)
		res := solve(n, div, less)
		var buf bytes.Buffer
		buf.WriteString("3")
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf(" %d", res[i]))
		}
		buf.WriteByte('\n')
		fmt.Print(buf.String())
	}
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

func solve(n int, div func(int, int, int) bool, less func(int, int) bool) []int {

	var loop func(indices []int, divisor int) Result

	loop = func(indices []int, divisor int) Result {
		if len(indices) == 1 {
			return NewResult([]int{1}, indices[0])
		}
		if len(indices) == 2 {
			if less(indices[0], indices[1]) {
				return NewResult([]int{1, 2}, indices[0])
			}
			return NewResult([]int{2, 1}, indices[1])
		}
		n := len(indices)
		leftZip := []int{indices[0]}
		var rightZip []int
		leftIndices := []int{0}
		var rightIndices []int
		for i := 1; i < n; i++ {
			if div(indices[0], indices[i], divisor) {
				leftZip = append(leftZip, indices[i])
				leftIndices = append(leftIndices, i)
			} else {
				rightZip = append(rightZip, indices[i])
				rightIndices = append(rightIndices, i)
			}
		}
		leftRes := loop(leftZip, 2*divisor)
		rightRes := loop(rightZip, 2*divisor)
		pos := leftRes.second
		if less(rightRes.second, leftRes.second) {
			pos = rightRes.second
		}
		res := make([]int, n)
		for i := 0; i < len(leftIndices); i++ {
			val := 2*leftRes.first[i] - 1
			if pos != leftRes.second {
				val++
			}
			res[leftIndices[i]] = val
		}
		for i := 0; i < len(rightIndices); i++ {
			val := 2*rightRes.first[i] - 1
			if pos != rightRes.second {
				val++
			}
			res[rightIndices[i]] = val
		}

		return NewResult(res, pos)
	}
	ind := make([]int, n)
	for i := 0; i < n; i++ {
		ind[i] = i
	}
	res := loop(ind, 2)
	return res.first
}

type Result struct {
	first  []int
	second int
}

func NewResult(ind []int, pos int) Result {
	return Result{first: ind, second: pos}
}
