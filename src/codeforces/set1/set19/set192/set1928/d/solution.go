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
		n, b, x := readThreeNums(reader)
		c := readNNums(reader, n)
		res := solve(b, x, c)

		buf.WriteString(fmt.Sprintf("%d\n", res))
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

func solve(b int, x int, c []int) int {

	calc := func(k int) int {
		if k == 1 {
			return 0
		}
		// 当k组时，它的收益时多少
		var res int
		for _, num := range c {
			if num <= k {
				// <= 1e16
				res += num * (num - 1) / 2 * b
			} else {
				avg, r := num/k, num%k
				res += (k - r) * avg * (num - avg) / 2 * b
				res += r * (avg + 1) * (num - (avg + 1)) / 2 * b
			}
		}

		res -= (k - 1) * x

		return res
	}

	var r int
	for _, num := range c {
		r = max(r, num)
	}
	l := 1
	for r-l > 2 {
		mid1 := l + (r-l)/3
		mid2 := r - (r-l)/3
		a := calc(mid1)
		b := calc(mid2)
		if a <= b {
			l = mid1
		} else {
			r = mid2
		}
	}

	if l == r {
		return calc(l)
	}
	if l+1 == r {
		return max(calc(l), calc(r))
	}
	// r - l <= 2
	return max(calc(l), calc(l+1), calc(r))
}
