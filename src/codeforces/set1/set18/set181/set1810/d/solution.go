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
		Q := make([][]int, n)
		for i := 0; i < n; i++ {
			s, _ := reader.ReadBytes('\n')
			var tp, pos int
			pos = readInt(s, 0, &tp) + 1
			if tp == 1 {
				Q[i] = make([]int, 4)
			} else {
				Q[i] = make([]int, 3)
			}
			Q[i][0] = tp
			for j := 1; j < len(Q[i]); j++ {
				pos = readInt(s, pos, &Q[i][j]) + 1
			}
		}
		res := solve(Q)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
		}
		buf.WriteByte('\n')
	}

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

func solve(Q [][]int) []int {
	// (a - b) * (n - 1) + a >= h
	// (a - b) * (n - 2) + a < h
	// if n == 1 , h <= a
	// if n == 2, a < h <= (a - b) + a
	var ans []int
	var i int
	for i < len(Q) && Q[i][0] == 2 {
		ans = append(ans, -1)
		i++
	}

	if i == len(Q) {
		return ans
	}
	var hi, lo int
	first := Q[i]
	if first[3] == 1 {
		hi = first[1]
		lo = 1
	} else {
		hi = (first[1]-first[2])*(first[3]-1) + first[1]
		lo = (first[1]-first[2])*(first[3]-2) + first[1] + 1
	}
	ans = append(ans, 1)

	get := func(h int, a, b int) int {
		if a >= h {
			return 1
		}
		// (a - b) * (n - 1) + a >= hi
		n := 1 + (h-a)/(a-b)
		for (a-b)*(n-1)+a < h {
			n++
		}
		for n >= 2 && (a-b)*(n-2)+a >= h {
			n--
		}
		return n
	}

	i++
	for i < len(Q) {
		cur := Q[i]
		if cur[0] == 2 {
			n1 := get(hi, cur[1], cur[2])
			n2 := get(lo, cur[1], cur[2])
			if n1 == n2 {
				ans = append(ans, n1)
			} else {
				ans = append(ans, -1)
			}
		} else {
			a, b, n := cur[1], cur[2], cur[3]
			x := (a-b)*(n-1) + a
			y := (a-b)*(n-2) + a + 1
			if n == 1 {
				y = 1
			}
			if max(lo, y) <= min(hi, x) {
				hi = min(hi, x)
				lo = max(lo, y)
				ans = append(ans, 1)
			} else {
				ans = append(ans, 0)
			}
		}
		i++
	}

	return ans
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
