package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, k := readTwoNums(reader)
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = readString(reader)
	}

	res := solve(k, words)

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

func solve(k int, words []string) int {
	sort.Strings(words)

	n := len(words)
	f := make([]int, n)
	for i := 1; i < n; i++ {
		for f[i] < len(words[i-1]) && f[i] < len(words[i]) && words[i-1][f[i]] == words[i][f[i]] {
			f[i]++
		}
	}

	var get func(l int, r int) []int

	get = func(l int, r int) []int {
		if r-l < 2 {
			return []int{0, 0}
		}
		mid := l + 1
		for i := l + 1; i < r; i++ {
			if f[i] < f[mid] {
				mid = i
			}
		}
		left := get(l, mid)
		right := get(mid, r)
		ans := make([]int, r-l+1)

		for i := 0; i < len(left); i++ {
			for j := 0; j < len(right); j++ {
				ans[i+j] = max(ans[i+j], left[i]+right[j]+i*j*f[mid])
			}
		}

		return ans
	}

	ans := get(0, n)

	return ans[k]
}
