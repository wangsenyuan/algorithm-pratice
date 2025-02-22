package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	panels := make([][]int, n)
	for i := 0; i < n; i++ {
		panels[i] = readNNums(reader, m)
	}
	symbols := make([][]int, n)
	for i := 0; i < n; i++ {
		symbols[i] = readNNums(reader, m)
	}
	x, y := readTwoNums(reader)
	res := solve(n, m, panels, symbols, x, y)
	fmt.Println(res)
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func readString(reader *bufio.Reader) string {
	s, _ := reader.ReadString('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return s[:i]
		}
	}
	return s
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

func solve(n int, m int, panel [][]int, symbols [][]int, x int, y int) int64 {
	x--
	y--

	// find same panels
	first := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if _, ok := first[panel[i][j]]; !ok {
				first[panel[i][j]] = i*m + j
			}
		}
	}
	set := NewSet(n * m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			set.Union(i*m+j, first[panel[i][j]])
			first[panel[i][j]] = set.Find(i*m + j)
		}
	}

	panel_symbols := make(map[int][]Pair)

	//colors := make(map[int]int)

	for i := 0; i < n*m; i++ {
		u, v := i/m, i%m
		//colors[set.Find(i)] = panel[u][v]
		if symbols[u][v] >= 0 {
			if panel_symbols[panel[u][v]] == nil {
				panel_symbols[panel[u][v]] = make([]Pair, 0, 1)
			}
			panel_symbols[panel[u][v]] = append(panel_symbols[panel[u][v]], Pair{u, v})
		}
	}
	// var res int
	// 最多把所有的symbols移除掉
	que := make([]int, n*m)
	rem := make([]bool, n*m)
	rem[x*m+y] = true
	var front, end int
	que[end] = x*m + y
	end++
	var res int64
	for front < end {
		cur := que[front]
		front++
		u, v := cur/m, cur%m
		cur = set.Find(cur)
		if panel[cur/m][cur%m] > 0 && panel[cur/m][cur%m] != symbols[u][v] {
			// repaint
			res += int64(set.cnt[cur])

			ss := panel_symbols[panel[cur/m][cur%m]]
			sortSpiral(ss, u, v)

			for _, s := range ss {
				if rem[s.first*m+s.second] {
					continue
				}
				que[end] = s.first*m + s.second
				end++
				rem[s.first*m+s.second] = true
			}
			panel[cur/m][cur%m] = symbols[u][v]

			if symbols[u][v] == 0 {
				continue
			}
			// 需要找到symbols[u][v]这个背景色的某个panel, 建它和cur合并
			if j, ok := first[symbols[u][v]]; ok {
				set.Union(cur, j)
				first[symbols[u][v]] = set.Find(cur)
				// 当前部分的ss已经被放入了队列，不需要合并到新的背景色中去
			}
			//panel[cur/m][cur%m] = symbols[u][v]
		}
	}

	return res
}

func sortSpiral(arr []Pair, x int, y int) {
	sort.Slice(arr, func(i, j int) bool {
		a := arr[i]
		b := arr[j]

		da := max(abs(a.first-x), abs(a.second-y))
		db := max(abs(b.first-x), abs(b.second-y))
		if da != db {
			//
			return da < db
		}
		// da == db
		if b.first < x && b.second < y && x-b.first == y-b.second {
			// b in the top-left corner
			return true
		}
		if a.first < x && a.second < y && x-a.first == y-a.second {
			// a in the top-left corner
			return false
		}
		pa := getPosition(a, Pair{x, y})
		pb := getPosition(b, Pair{x, y})

		if pa != pb {
			return pa < pb
		}
		// a could be in top, right, bot, left
		// pa == pb
		if pa == 0 {
			return a.second < b.second
		}
		if pa == 1 {
			return a.first < b.first
		}
		if pa == 2 {
			return a.second > b.second
		}
		return a.first > b.first
	})
}

func getPosition(a Pair, b Pair) int {
	if a.first+a.second < b.first+b.second {
		// 在左上部分
		if a.first-a.second < b.first-b.second {
			return 0
		}
		return 3
	}
	// 在右下部分
	if a.first-a.second <= b.first-b.second {
		return 1
	}
	return 2
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

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Pair struct {
	first  int
	second int
}

type Set struct {
	arr []int
	cnt []int
}

func NewSet(n int) *Set {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &Set{arr, cnt}
}

func (s *Set) Find(x int) int {
	if s.arr[x] != x {
		s.arr[x] = s.Find(s.arr[x])
	}
	return s.arr[x]
}

func (s *Set) Union(a, b int) {
	a = s.Find(a)
	b = s.Find(b)

	if a != b {
		if s.cnt[a] < s.cnt[b] {
			a, b = b, a
		}
		s.cnt[a] += s.cnt[b]
		s.arr[b] = a
	}
}
