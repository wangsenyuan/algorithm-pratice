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
		n, m, k := readThreeNums(reader)
		x := readNNums(reader, n)
		y := readNNums(reader, m)
		ps := make([][]int, k)
		for i := 0; i < k; i++ {
			ps[i] = readNNums(reader, 2)
		}

		res := solve(x, y, ps)
		buf.WriteString(fmt.Sprintf("%d\n", res))
	}

	fmt.Print(buf.String())
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

func solve(x []int, y []int, persons [][]int) int {
	var px []pair
	var py []pair
	for i := 0; i < len(persons); i++ {
		a := sort.SearchInts(x, persons[i][0])
		b := sort.SearchInts(y, persons[i][1])
		if a < len(x) && x[a] == persons[i][0] && b < len(y) && y[b] == persons[i][1] {
			continue
		}
		if a < len(x) && x[a] == persons[i][0] {
			// 它在垂直方向的街道上
			px = append(px, pair{persons[i][0], persons[i][1]})
		} else {
			py = append(py, pair{persons[i][1], persons[i][0]})
		}
	}

	res := solve_x(px, y)

	res += solve_x(py, x)

	return res
}

type pair struct {
	first  int
	second int
}

func solve_x(people []pair, ys []int) int {
	// 将所有的y进行排序，进而进行压缩
	id := make(map[int]int)
	for _, p := range people {
		id[p.second]++
	}
	for _, num := range ys {
		id[num]++
	}
	arr := make([]int, 0, len(id))
	for k := range id {
		arr = append(arr, k)
	}
	sort.Ints(arr)

	for i, num := range arr {
		id[num] = i
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].first < people[j].first
	})

	tr := make(fenwick, len(arr)+10)

	var res int

	for i := 0; i < len(people); {
		j := i
		for i < len(people) && people[i].first == people[j].first {
			y := people[i].second
			k := sort.SearchInts(ys, y)
			// ys[k] > y
			// ys[k-1] < k
			// k1 肯定是存在的
			y1 := id[ys[k]]
			y2 := id[ys[k-1]]

			res += tr.query(y2, y1)

			i++
		}
		for j < i {
			y := id[people[j].second]
			tr.update(y, 1)
			j++
		}
	}

	return res
}

type fenwick []int

func (tr fenwick) update(p int, v int) {
	p++
	for p < len(tr) {
		tr[p] += v
		p += p & -p
	}
}

func (tr fenwick) pre(p int) int {
	var res int
	p++
	for p > 0 {
		res += tr[p]
		p -= p & -p
	}
	return res
}

func (tr fenwick) query(l int, r int) int {
	return tr.pre(r) - tr.pre(l-1)
}
