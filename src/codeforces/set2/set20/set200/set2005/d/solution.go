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
		a := readNNums(reader, n)
		b := readNNums(reader, n)
		res := solve(a, b)
		buf.WriteString(fmt.Sprintf("%d %d\n", res[0], res[1]))
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

func solve(a []int, b []int) []int {
	n := len(a)

	type state struct {
		g1  int
		g2  int
		pos int
	}

	suf := make([]state, n+1)
	for i := n - 1; i >= 0; i-- {
		suf[i] = state{gcd(suf[i+1].g1, a[i]), gcd(suf[i+1].g2, b[i]), -1}
	}

	ans := make([]int, 2)

	var states []state

	var fa, fb int

	for i := range n {
		for j := range states {
			states[j].g1 = gcd(states[j].g1, b[i])
			states[j].g2 = gcd(states[j].g2, a[i])
		}
		states = append(states, state{gcd(fa, b[i]), gcd(fb, a[i]), i})

		j := 1
		for k := 1; k < len(states); k++ {
			if states[k-1].g1 != states[k].g1 || states[k-1].g2 != states[k].g2 {
				states[j] = states[k]
				j++
			}
		}
		states = states[:j]

		pre := i + 1

		for j > 0 {
			p := states[j-1]
			j--
			tmp := gcd(suf[i+1].g1, p.g1) + gcd(suf[i+1].g2, p.g2)
			if tmp > ans[0] {
				ans[0] = tmp
				ans[1] = pre - p.pos
			} else if tmp == ans[0] {
				ans[1] += pre - p.pos
			}
			pre = p.pos
		}
		fa = gcd(fa, a[i])
		fb = gcd(fb, b[i])
	}

	return ans
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
