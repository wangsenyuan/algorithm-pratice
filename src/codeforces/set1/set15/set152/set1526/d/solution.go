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
		s := readString(reader)
		res := solve(s)
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

const ATON = "ATON"

func solve(s string) string {
	cnt := make([]int, 4)

	get := func(x byte) int {
		if x == 'A' {
			return 0
		} else if x == 'T' {
			return 1
		} else if x == 'O' {
			return 2
		} else {
			return 3
		}
	}

	for _, x := range []byte(s) {
		cnt[get(x)]++
	}
	n := len(s)

	flips := make([][4]int, 4)

	for i := range 4 {
		var cnt int
		for j := 0; j < n; j++ {
			// 如果把i排在s(j)的后面，需要移动cnt次
			flips[i][get(s[j])] += cnt
			if get(s[j]) == i {
				cnt++
			}
		}
	}

	count := func(p []int) int {
		var res int
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				res += flips[p[j]][p[i]]
			}
		}
		return res
	}

	perm := []int{0, 1, 2, 3}

	var ans []int
	var best int
	var dfs func(i int)

	dfs = func(i int) {
		if i == 4 {
			tmp := count(perm)
			if ans == nil || tmp > best {
				if ans == nil {
					ans = make([]int, 4)
				}
				copy(ans, perm)
				best = tmp
			}
			return
		}
		dfs(i + 1)
		for j := i + 1; j < 4; j++ {
			perm[i], perm[j] = perm[j], perm[i]
			dfs(i + 1)
			perm[i], perm[j] = perm[j], perm[i]
		}
	}

	dfs(0)

	buf := make([]byte, n)

	for i, j := 0, 0; i < 4; i++ {
		x := ans[i]
		for cnt[x] > 0 {
			buf[j] = ATON[x]
			j++
			cnt[x]--
		}
	}

	return string(buf)
}

func countMinMoves(s string, t string) int {
	if s == t {
		return 0
	}
	for i := 0; i < len(t); i++ {
		if t[i] == s[0] {
			return i + countMinMoves(s[1:], t[:i]+t[i+1:])
		}
	}
	return 0
}
