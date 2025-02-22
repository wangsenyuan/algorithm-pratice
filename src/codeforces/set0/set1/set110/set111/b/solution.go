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
	qs := make([][]int, n)
	for i := 0; i < n; i++ {
		qs[i] = readNNums(reader, 2)
	}
	res := solve(qs)
	var buf bytes.Buffer
	for _, ans := range res {
		buf.WriteString(fmt.Sprintf("%d\n", ans))
	}
	fmt.Print(buf.String())
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

const MAX_X = 100010

func getFactors(num int) []int {
	var res []int
	for x := 1; x <= num/x; x++ {
		if num%x == 0 {
			res = append(res, x)
			y := num / x
			if x != y {
				res = append(res, y)
			}
		}
	}
	return res
}

func solve(qs [][]int) []int {
	//n := len(qs)

	used := make([]int, MAX_X)
	for i := 0; i < MAX_X; i++ {
		used[i] = -1
	}
	check := func(i int, q []int) int {
		x, y := q[0], q[1]
		fs := getFactors(x)
		var res int
		if y == 0 {
			res = len(fs)
		} else {
			for _, f := range fs {
				if used[f] < i-y {
					res++
				}
			}
		}
		for _, f := range fs {
			used[f] = i
		}
		return res
	}

	ans := make([]int, len(qs))

	for i, q := range qs {
		ans[i] = check(i, q)
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
