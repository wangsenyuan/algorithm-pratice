package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d %d\n", x[0], x[1]))
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

func solve(a []int) [][]int {
	n := len(a)
	win := a[n-1]
	lose := 3 - win
	pos := make([][]int, 3)

	for i := 0; i < n; i++ {
		pos[a[i]] = append(pos[a[i]], i)
	}
	m := len(pos[win])

	find := func(who int, i int, t int) int {
		// 找到第i个位置
		if i+t <= len(pos[who]) {
			return pos[who][i+t-1]
		}
		return n
	}
	// if pass returns s
	check := func(t int) int {
		cnt := make([]int, 3)
		for i, j, k := 0, 0, 0; k < n; {
			// 要么是
			x := find(win, i, t)
			y := find(lose, j, t)
			if x == y {
				return 0
			}

			if x < y {
				i += t
				// win先到
				k = x + 1
				cnt[win]++
				j = sort.SearchInts(pos[lose], k)
			} else {
				j += t
				k = y + 1
				cnt[lose]++
				i = sort.SearchInts(pos[win], k)
			}
		}
		if cnt[win] > cnt[lose] {
			return cnt[win]
		}
		return 0
	}

	var res [][]int
	for i := 1; i <= m; i++ {
		s := check(i)
		if s > 0 {
			res = append(res, []int{s, i})
		}
	}

	slices.SortFunc(res, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})
	return res
}
