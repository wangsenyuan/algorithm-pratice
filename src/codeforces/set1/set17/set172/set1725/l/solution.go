package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	A := readNNums(reader, n)
	res := solve(A)
	fmt.Println(res)
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

func solve(a []int) int {
	// sum of a < 0 => -1
	n := len(a)
	b := make([]int, n+1)
	for i := 0; i < n; i++ {
		b[i+1] = b[i] + a[i]
	}
	if b[n] < b[0] {
		return -1
	}
	for i := 1; i < n; i++ {
		if b[i] < b[0] || b[i] > b[n] {
			return -1
		}
	}
	// b[0] <= b[1] < ... <= b[n]

	buf := make([]int, n+1)
	copy(buf, b)
	merge := func(l int, m int, r int) int {
		// [l...m) is sorted
		// [m...r) is sorted
		// if b[i] > b[j]
		var res int
		i, j := l, m
		for k := l; k < r; k++ {
			if j == r || i < m && b[i] <= b[j] {
				buf[k] = b[i]
				i++
				continue
			}
			res += m - i
			buf[k] = b[j]
			j++
		}
		copy(b[l:r], buf[l:r])
		return res
	}

	var dfs func(l int, r int) int
	dfs = func(l int, r int) int {
		if l+1 >= r {
			return 0
		}
		mid := (l + r) / 2
		res := dfs(l, mid) + dfs(mid, r)
		res += merge(l, mid, r)
		return res
	}

	return dfs(0, len(b))
}
