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

	segs := make([][]int, n)

	for i := 0; i < n; i++ {
		segs[i] = readNNums(reader, 2)
	}

	res := solve(segs)

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

func solve(segs [][]int) int {
	n := len(segs)

	sort.Slice(segs, func(i, j int) bool {
		return segs[i][1] < segs[j][1] || segs[i][1] == segs[j][1] && segs[i][0] > segs[j][0]
	})

	ans := get(segs[1:])

	sort.Slice(segs, func(i, j int) bool {
		// 相等的情况下，去掉短的
		return segs[i][0] < segs[j][0] || segs[i][0] == segs[j][0] && segs[i][1] < segs[j][1]
	})

	ans = max(ans, get(segs[:n-1]))

	return ans
}

func get(segs [][]int) int {
	if len(segs) == 0 {
		return 0
	}
	l := segs[0][0]
	r := segs[0][1]
	for i := 1; i < len(segs); i++ {
		l = max(segs[i][0], l)
		r = min(segs[i][1], r)
	}
	return max(r-l, 0)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
