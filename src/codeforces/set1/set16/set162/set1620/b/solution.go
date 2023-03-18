package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)

	for tc > 0 {
		tc--
		w, h := readTwoNums(reader)
		coords := make([][]int, 4)
		for i := 0; i < 4; i++ {
			coords[i] = readLenNums(reader)
		}
		res := solve(w, h, coords)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
}

func readLenNums(r *bufio.Reader) []int {
	s, _ := r.ReadBytes('\n')
	var n int
	pos := readInt(s, 0, &n) + 1
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		pos = readInt(s, pos, &arr[i]) + 1
	}
	return arr
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(w int, h int, coords [][]int) int64 {
	// coords[0] for y = 0
	// coords[1] for y = h
	// coords[2] for x = 0
	// coords[3] for x = w
	a := area(coords[0][len(coords[0])-1]-coords[0][0], h)
	b := area(coords[1][len(coords[1])-1]-coords[1][0], h)
	c := area(coords[2][len(coords[2])-1]-coords[2][0], w)
	d := area(coords[3][len(coords[3])-1]-coords[3][0], w)

	return max(a, b, c, d)
}

func area(w int, h int) int64 {
	return int64(w) * int64(h)
}

func max(arr ...int64) int64 {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > res {
			res = arr[i]
		}
	}
	return res
}
