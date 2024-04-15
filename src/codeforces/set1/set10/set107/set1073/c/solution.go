package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	readNum(reader)
	seq := readString(reader)
	x, y := readTwoNums(reader)
	res := solve(seq, x, y)

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

func solve(seq string, x int, y int) int {
	n := len(seq)

	if abs(x)+abs(y) > n || n%2 != (abs(x)+abs(y))%2 {
		return -1
	}

	var dx, dy int

	suf := make([][]int, n+1)
	suf[n] = make([]int, 2)
	for i := n - 1; i >= 0; i-- {
		suf[i] = make([]int, 2)
		if seq[i] == 'U' {
			dy++
		} else if seq[i] == 'D' {
			dy--
		} else if seq[i] == 'R' {
			dx++
		} else {
			dx--
		}

		suf[i][0] = dx
		suf[i][1] = dy
	}

	if dx == x && dy == y {
		// already good
		return 0
	}

	best := n

	dx = 0
	dy = 0
	pref := make([][]int, n+1)
	pref[0] = make([]int, 2)

	check := func(l int, r int) bool {
		u := pref[l][0] + suf[r][0]
		v := pref[l][1] + suf[r][1]
		u = x - u
		v = y - v
		return abs(u)+abs(v) <= r-l
	}

	var l int

	for r := 0; r < n; r++ {
		for l+1 < r && check(l+1, r) {
			l++
		}

		if l < r && check(l, r) {
			best = min(best, r-l)
		}

		if seq[r] == 'U' {
			dy++
		} else if seq[r] == 'D' {
			dy--
		} else if seq[r] == 'R' {
			dx++
		} else {
			dx--
		}

		pref[r+1] = []int{dx, dy}
	}

	for l+1 < n && check(l+1, n) {
		l++
		best = min(best, n-l)
	}

	return best

}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
