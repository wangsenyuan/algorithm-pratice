package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, h := readTwoNums(reader)
	segs := make([][]int, n)
	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 2)
	}

	res := solve(h, segs)

	fmt.Println(res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(h int, segs [][]int) int {
	if len(segs) == 0 {
		return h
	}

	if len(segs) == 1 {
		return h + segs[0][1] - segs[0][0]
	}
	//
	//slices.SortFunc(segs, func(a, b []int) int {
	//	return a[0] - b[0]
	//})
	n := len(segs)

	pref := make([]int, n+1)

	for i := 1; i < n; i++ {
		pref[i] = pref[i-1] + segs[i][0] - segs[i-1][1]
	}

	best := h

	for l, r := 0, 0; l < n; l++ {
		// 如果从l的左端开始
		x := segs[l][0]
		for r < n && pref[r]-pref[l] < h {
			r++
		}
		// pref[r] - pref[l] > h
		dx := h - (pref[r-1] - pref[l])
		x1 := segs[r-1][1] + dx
		best = max(best, x1-x)
	}

	return best
}
