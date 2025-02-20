package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for range tc {
		res := process(reader)
		for _, x := range res {
			if x {
				buf.WriteString("YES\n")
			} else {
				buf.WriteString("NO\n")
			}
		}
	}
	buf.WriteTo(os.Stdout)
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}
func process(reader *bufio.Reader) []bool {
	n, m := readTwoNums(reader)
	p := readNNums(reader, n)
	s := readString(reader)
	qs := make([]int, m)
	for i := range m {
		qs[i] = readNum(reader)
	}
	return solve(p, s, qs)
}

func solve(p []int, s string, qs []int) []bool {
	n := len(s)

	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = p[i] - 1
		if i > 0 {
			dp[i] = max(dp[i], dp[i-1])
		}
	}

	var bad int

	set := func(i int) {
		if dp[i] > i {
			bad++
		}
	}

	unset := func(i int) {
		if dp[i] > i {
			bad--
		}
	}

	for i := 0; i+1 < n; i++ {
		if s[i] == 'L' && s[i+1] == 'R' {
			set(i)
		}
	}

	buf := []byte(s)

	play := func(i int) bool {
		if i+1 < n && buf[i] == 'L' && buf[i+1] == 'R' {
			unset(i)
		}
		if i > 0 && buf[i-1] == 'L' && buf[i] == 'R' {
			unset(i - 1)
		}
		if buf[i] == 'L' {
			buf[i] = 'R'
		} else {
			buf[i] = 'L'
		}
		if i+1 < n && buf[i] == 'L' && buf[i+1] == 'R' {
			set(i)
		}
		if i > 0 && buf[i-1] == 'L' && buf[i] == 'R' {
			set(i - 1)
		}

		return bad == 0
	}

	ans := make([]bool, len(qs))
	for i, x := range qs {
		ans[i] = play(x - 1)
	}

	return ans
}
