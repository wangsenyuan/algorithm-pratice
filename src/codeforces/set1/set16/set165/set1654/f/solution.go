package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	s := readString(reader)
	res := solve(n, s)
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

func solve(n int, s string) string {
	a := make([]int, 1<<n)
	h := 1 << n
	for i := 0; i < h; i++ {
		a[i] = i
	}

	sort.Slice(a, func(i, j int) bool {
		return s[a[i]] < s[a[j]]
	})

	v := make([]int, h)
	for i := 1; i < h; i++ {
		v[a[i]] = v[a[i-1]]
		if s[a[i]] != s[a[i-1]] {
			v[a[i]]++
		}
	}

	tmp := make([]int, h)

	for k := 1; k < h; k <<= 1 {
		cmp := func(i, j int) bool {
			if v[a[i]] == v[a[j]] {
				return v[a[i]^k] < v[a[j]^k]
			}
			return v[a[i]] < v[a[j]]
		}
		sort.Slice(a, cmp)
		tmp[a[0]] = 0
		for i := 1; i < h; i++ {
			tmp[a[i]] = tmp[a[i-1]]
			if cmp(i-1, i) {
				tmp[a[i]]++
			}
		}
		copy(v, tmp)
	}
	buf := make([]byte, h)
	for i := 0; i < h; i++ {
		buf[i] = s[i^a[0]]
	}
	return string(buf)
}
