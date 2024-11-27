package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
}

func process(reader *bufio.Reader) string {
	n := readNum(reader)
	words := make([]string, 3)
	for i := range 3 {
		words[i] = readString(reader)
	}

	return solve(n, words)
}

var users = []string{"Kuro", "Shiro", "Katie"}

func solve(n int, words []string) string {

	check := func(word string) int {
		m := len(word)

		if n > 0 && n%m == 0 {
			// 不管怎么处理，都可以得到m的结果
			return m
		}

		freq := make(map[byte]int)
		for _, x := range []byte(word) {
			freq[x]++
		}

		var best int

		for _, u := range freq {
			if n == 1 {
				if u == m {
					best = m - 1
				} else {
					best = max(best, u+1)
				}
				continue
			}
			v := m - u

			if v < n {
				// 如果 n - v 等于1，那么可以选择一个，进行2次操作，从而将其变成0
				best = m
			} else {
				// v >= n
				best = max(best, u+n)
			}
		}

		return min(best, m)
	}

	cnt := make([]int, 3)
	for i := range 3 {
		cnt[i] = check(words[i])
	}
	x := slices.Max(cnt)
	var winers int
	for i := 0; i < 3; i++ {
		if cnt[i] == x {
			winers++
		}
	}
	if winers > 1 {
		return "Draw"
	}

	for i := 0; i < 3; i++ {
		if cnt[i] == x {
			return users[i]
		}
	}

	return "Draw"
}
