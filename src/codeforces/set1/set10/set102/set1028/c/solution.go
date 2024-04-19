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

	rects := make([][]int, n)
	for i := 0; i < n; i++ {
		rects[i] = readNNums(reader, 4)
	}

	res := solve(rects)

	// res always exists for the problem
	fmt.Printf("%d %d\n", res[0], res[1])
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

func solve(rects [][]int) []int {
	res := solve1(rects, 2)
	if len(res) != 0 {
		return res
	}
	return solve1(rects, 3)
}

func solve1(rects [][]int, id int) []int {
	sort.Slice(rects, func(i, j int) bool {
		return rects[i][id] < rects[j][id] || rects[i][id] == rects[j][id] && rects[i][id-2] > rects[j][id-2]
	})

	n := len(rects)

	res := find(rects[1:])

	if len(res) != 0 {
		return res
	}

	sort.Slice(rects, func(i, j int) bool {
		return rects[i][id-2] < rects[j][id-2] || rects[i][id-2] == rects[j][id-2] && rects[i][id] < rects[j][id]
	})

	return find(rects[:n-1])
}

func find(rects [][]int) []int {
	ox, x := check(rects, 0)
	oy, y := check(rects, 1)
	if !ox || !oy {
		return nil
	}
	return []int{x, y}
}

func check(rects [][]int, id int) (bool, int) {
	l := rects[0][id]
	r := rects[0][id+2]
	for _, rect := range rects {
		l = max(l, rect[id])
		r = min(r, rect[id+2])
	}
	if l > r {
		return false, -1
	}
	return true, l
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	return a + b - max(a, b)
}
