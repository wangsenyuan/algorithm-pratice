package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	tc := readNum(reader)

	var buf bytes.Buffer

	for tc > 0 {
		tc--
		n, m := readTwoNums(reader)
		teachers := readNNums(reader, n)
		groups := make([][]int, m)
		for i := 0; i < m; i++ {
			sz := readNum(reader)
			groups[i] = readNNums(reader, sz)
		}
		res := solve(teachers, groups)
		buf.WriteString(res)
		buf.WriteByte('\n')
	}
	fmt.Print(buf.String())
}

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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

func solve(teachers []int, groups [][]int) string {
	sort.Ints(teachers)
	reverse(teachers)

	avg := make([]StudentGroup, len(groups))

	for i, g := range groups {
		avg[i] = StudentGroup{Pair{sum(g), len(g)}, i}
	}

	sort.Sort(sort.Reverse(SGS(avg)))

	pos := make([]int, len(groups))

	for i, sg := range avg {
		pos[sg.pos] = i
	}

	m := len(avg)
	next := make([]int, m)
	prev := make([]int, m)
	curr := make([]int, m)
	for i := 0; i < m-1; i++ {
		if (Pair{int64(teachers[i+1]), 1}).Less(avg[i].group) {
			next[i] = 1
		}
	}

	for i := 0; i < m; i++ {
		if (Pair{int64(teachers[i]), 1}).Less(avg[i].group) {
			curr[i] = 1
		}
	}

	for i := 1; i < m; i++ {
		if (Pair{int64(teachers[i-1]), 1}).Less(avg[i].group) {
			prev[i] = 1
		}
	}

	for i := 1; i < m; i++ {
		next[i] += next[i-1]
		prev[i] += prev[i-1]
		curr[i] += curr[i-1]
	}

	tot := curr[m-1]
	var buf bytes.Buffer
	for i := 0; i < m; i++ {
		id := pos[i]
		cur := Pair{sum(groups[i]), len(groups[i]) - 1}
		for j := 0; j < len(groups[i]); j++ {
			cur = Pair{cur.first - int64(groups[i][j]), cur.second}
			l, r := -1, m
			for l+1 < r {
				mid := (l + r) / 2
				if !avg[mid].group.Less(cur) {
					l = mid
				} else {
					r = mid
				}
			}
			if r > id {
				r--
			}
			tr := 1
			if (Pair{int64(teachers[r]), 1}).Less(cur) {
				tr = 0
			}
			if min(r, id)-1 >= 0 && curr[min(r, id)-1] > 0 {
				tr = 0
			}
			if curr[max(r, id)] != tot {
				tr = 0
			}
			if r < id && (r > 0 && next[r-1] != next[id-1] || r == 0 && next[id-1] > 0) {
				tr = 0
			}
			if r > id && (prev[r] != prev[id]) {
				tr = 0
			}
			buf.WriteByte(byte('0' + tr))
			cur = Pair{cur.first + int64(groups[i][j]), cur.second}
		}
	}
	return buf.String()
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	return a + b - min(a, b)
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func sum(arr []int) int64 {
	var res int64
	for _, num := range arr {
		res += int64(num)
	}
	return res
}

type Pair struct {
	first  int64
	second int
}

func (this Pair) Less(that Pair) bool {
	return this.first*int64(that.second) < int64(this.second)*that.first
}

type StudentGroup struct {
	group Pair
	pos   int
}

type SGS []StudentGroup

func (ps SGS) Len() int {
	return len(ps)
}

func (ps SGS) Less(i, j int) bool {
	return ps[i].group.Less(ps[j].group)
}

func (ps SGS) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}
