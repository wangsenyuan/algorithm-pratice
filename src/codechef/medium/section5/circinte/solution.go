package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadBytes('\n')
	var m uint64
	var n int
	pos := readUint64(s, 0, &m)
	readInt(s, pos+1, &n)
	intervals := make([][]int64, n)
	for i := 0; i < n; i++ {
		s, _ = reader.ReadBytes('\n')
		var x, y uint64
		pos = readUint64(s, 0, &x)
		readUint64(s, pos+1, &y)
		intervals[i] = []int64{int64(x), int64(y)}
	}
	res := solve(int64(m), intervals)
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

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}
func solve(m int64, intervals [][]int64) int64 {
	first := []int64{intervals[0][0], intervals[0][1]}
	first[0] += m
	first[1] += m
	intervals = append(intervals, first)
	n := len(intervals)
	check := func(d int64) bool {
		r := intervals[n-1][1]
		for i := n - 2; i >= 0; i-- {
			r = min(r-d, intervals[i][1])
			if r < intervals[i][0] {
				return false
			}
		}
		x := r + m
		for i := 1; i < n; i++ {
			r = max(r+d, intervals[i][0])
			if r > intervals[i][1] {
				return false
			}
		}
		return r <= x
	}

	var l, r int64 = 0, m

	for l < r {
		mid := (l + r) / 2
		if !check(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r - 1
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int64) int64 {
	if a <= b {
		return a
	}
	return b
}
