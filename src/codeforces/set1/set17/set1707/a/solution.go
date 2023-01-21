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
		n, k := readTwoNums(reader)
		A := readNNums(reader, n)
		res := solve(A, k)
		buf.WriteString(fmt.Sprintf("%s\n", res))
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

func solve(A []int, iq int) string {
	n := len(A)
	res := make([]byte, n)

	var cur int

	for i := n - 1; i >= 0; i-- {
		if A[i] > cur && cur == iq {
			res[i] = '0'
		} else {
			res[i] = '1'
			if A[i] > cur {
				cur++
			}
		}
	}

	return string(res)
}
func solve1(A []int, iq int) string {
	// when iq > A[i], always to test
	// when iq <= A[i], may test, but iq--, else not stay no change
	// only when iq > 0, can test
	n := len(A)

	check := func(x int) bool {
		// if x is the first bad contest
		// can we start from x, and take all the contests
		cur := iq
		for i := x; i < n; i++ {
			if cur == 0 {
				return false
			}
			if A[i] > cur {
				cur--
			}
		}
		return true
	}

	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	res := make([]byte, n)

	for i := 0; i < r; i++ {
		if iq >= A[i] {
			res[i] = '1'
		} else {
			res[i] = '0'
		}
	}

	for i := r; i < n; i++ {
		res[i] = '1'
	}

	return string(res)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
