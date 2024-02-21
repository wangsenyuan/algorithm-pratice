package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	a := readNNums(reader, n)

	m := readNum(reader)
	ops := make([][]int, m)
	for i := 0; i < m; i++ {
		s, _ := reader.ReadBytes('\n')
		if s[0] == '1' {
			var p, x int
			pos := readInt(s, 2, &p)
			readInt(s, pos+1, &x)
			ops[i] = []int{1, p, x}
		} else {
			var x int
			readInt(s, 2, &x)
			ops[i] = []int{2, x}
		}
	}

	res := solve(n, a, ops)

	s := fmt.Sprintf("%v", res)

	s = s[1 : len(s)-1]
	fmt.Println(s)
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
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

func solve(n int, a []int, ops [][]int) []int {
	marked := make([]bool, n)

	val := make([]int, n)
	copy(val, a)

	hi := -1

	for i := len(ops) - 1; i >= 0; i-- {
		cur := ops[i]
		if cur[0] == 1 {
			// ops 1
			p := cur[1] - 1
			x := cur[2]
			if marked[p] {
				continue
			}
			val[p] = max(x, hi)
			marked[p] = true
		} else {
			x := cur[1]
			hi = max(hi, x)
		}
	}

	for i := 0; i < n; i++ {
		if !marked[i] {
			val[i] = max(val[i], hi)
		}
	}

	return val
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
