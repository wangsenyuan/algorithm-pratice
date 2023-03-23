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
	Q := make([][]int, n)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadBytes('\n')
		var t, x int
		pos := readInt(s, 0, &t)
		pos = readInt(s, pos+1, &x)

		if t == 1 {
			Q[i] = []int{1, x}
		} else {
			var y int
			pos = readInt(s, pos+1, &y)
			Q[i] = []int{2, x, y}
		}
	}

	res := solve(Q)

	var buf bytes.Buffer

	for _, num := range res {
		buf.WriteString(fmt.Sprintf("%d ", num))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
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

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
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

func solve(Q [][]int) []int {
	// 2 x, y
	n := len(Q)

	p := make(map[int]int)

	get := func(x int) int {
		if v, ok := p[x]; ok {
			return v
		}
		p[x] = x
		return x
	}

	var res []int
	for i := n - 1; i >= 0; i-- {
		if Q[i][0] == 1 {
			x := Q[i][1]
			res = append(res, get(x))
		} else {
			x := Q[i][1]
			y := Q[i][2]
			p[x] = get(y)
		}
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return res
}
