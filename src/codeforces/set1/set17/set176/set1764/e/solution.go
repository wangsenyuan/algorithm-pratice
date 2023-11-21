package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, k := readTwoNums(reader)
		a := make([]int, n)
		b := make([]int, n)
		for i := 0; i < n; i++ {
			a[i], b[i] = readTwoNums(reader)
		}
		res := solve(k, a, b)
		if res {
			buf.WriteString("YES\n")
		} else {
			buf.WriteString("NO\n")
		}
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

func solve(k int, a []int, b []int) bool {
	n := len(a)

	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	sort.Slice(id, func(i, j int) bool {
		return a[id[i]] < a[id[j]]
	})

	mx := make([]int, n+1)
	for i := 0; i < n; i++ {
		mx[i+1] = max(a[id[i]]+b[id[i]], mx[i])
	}

	ans := make([]bool, n)

	for i := 0; i+1 < n; i++ {
		if a[id[i]]+b[id[i]] >= k {
			ans[id[i]] = true
		}
	}

	var calc func(i int) int

	calc = func(i int) int {
		if i == 1 {
			return a[id[0]]
		}

		far := max(calc(i-1), mx[i-2])

		x := min(far, a[id[i-1]]) + b[id[i-1]]

		return max(x, a[id[i-1]])
	}

	if calc(n) >= k {
		ans[id[n-1]] = true
	}

	return ans[0]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
