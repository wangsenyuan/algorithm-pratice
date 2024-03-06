package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, x := readTwoNums(reader)
		res := solve(n, x)
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

const H = 62

func solve(n int, x int) int {
	if n == x {
		return n
	}
	for n > x {
		lo := n & -n
		n ^= lo
		if n == x && n&(lo<<1) == 0 {
			return n | (lo << 1)
		}
	}
	return -1
}

func solve1(n int, x int) int {
	// if n[d] = 0, but x[d] = 1 => -1
	// if n[d] = 0, x[d] must be 0, and no constraints on m
	//    n[d] = 1, but x[d] = 1 => m[d] = 1,
	//    n[d] = 1, but x[d] = 0 => m[d] = 0
	var lo, hi int = 0, math.MaxInt64
	for i := 0; i < H; i++ {
		a := (n >> i) & 1
		b := (x >> i) & 1
		if a == 0 {
			if b == 1 {
				return -1
			}
			continue
		}
		// a = 1
		mask := (1 << i) - 1
		if b == 1 {
			hi = min(hi, mask-(n&mask))
		} else {
			lo = max(lo, mask-(n&mask)+1)
		}
		if hi < lo {
			return -1
		}
	}
	return n + lo
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
