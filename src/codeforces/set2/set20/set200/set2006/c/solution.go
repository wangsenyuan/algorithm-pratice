package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/bits"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n := readNum(reader)
		a := readNNums(reader, n)
		res := solve(a)
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

func solve(a []int) int {
	n := len(a)
	if n == 1 {
		return 1
	}
	n--
	st := make([][]int, n+1)
	h := bits.Len(uint(n))
	for i := range n {
		st[i] = make([]int, h+1)
		st[i][0] = abs(a[i] - a[i+1])
	}
	st[n] = make([]int, h+1)
	for j := 0; j <= h; j++ {
		for i := 0; i+1<<j < n; i++ {
			st[i][j+1] = gcd(st[i][j], st[i+(1<<j)][j])
		}
	}

	next := make([]int, n+1)
	next[n] = n
	for i := n - 1; i >= 0; i-- {
		if a[i] == a[i+1] {
			next[i] = next[i+1]
		} else {
			next[i] = i
		}
	}

	query := func(l int, r int) int {
		s := bits.Len(uint(r-l+1)) - 1
		return gcd(st[l][s], st[r-(1<<s)+1][s])
	}

	ans := 1
	var l, r int

	for i := 0; i < n; i++ {
		l = max(l, next[i])
		r = max(r, l)
		for r < n && countDigits(query(l, r)) > 1 {
			r++
		}
		ans += n - r
		ans += next[i] - i + 1
	}

	return ans
}

func countDigits(num int) int {
	return bits.OnesCount(uint(num))
}

func abs(a int) int {
	return max(a, -a)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
