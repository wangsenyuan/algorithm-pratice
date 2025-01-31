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
	pts := make([][]int, n)
	for i := 0; i < n; i++ {
		pts[i] = readNNums(reader, 3)
	}
	res := solve(pts)
	fmt.Println(res)
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

func solve(points [][]int) int64 {
	rows := make(map[int]int64)
	cols := make(map[int]int64)
	grid := make(map[int]map[int]int)
	for _, pt := range points {
		// x, y, v
		rows[pt[0]] += int64(pt[2])
		cols[pt[1]] += int64(pt[2])
		if grid[pt[0]] == nil {
			grid[pt[0]] = make(map[int]int)
		}
		grid[pt[0]][pt[1]] = pt[2]
	}

	xs := convertToItems(rows)
	ys := convertToItems(cols)

	sort.Slice(ys, func(i, j int) bool {
		return ys[i].val > ys[j].val
	})

	get := func(x int, y int) (int, bool) {
		if _, ok := grid[x]; !ok {
			return 0, false
		}
		if v, ok := grid[x][y]; !ok {
			return 0, false
		} else {
			return v, true
		}
	}

	var ans int64
	for _, x := range xs {
		for _, y := range ys {
			v, ok := get(x.pos, y.pos)
			if !ok {
				ans = max(ans, x.val+y.val)
				break
			}
			ans = max(ans, x.val+y.val-int64(v))
		}
	}

	return ans
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Item struct {
	val int64
	pos int
}

func convertToItems(vals map[int]int64) []Item {
	its := make([]Item, len(vals))
	var j int
	for i, v := range vals {
		its[j] = Item{v, i}
		j++
	}
	return its
}
