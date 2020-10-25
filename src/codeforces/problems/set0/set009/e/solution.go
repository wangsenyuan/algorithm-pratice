package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	E := make([][]int, m)
	for i := 0; i < m; i++ {
		E[i] = readNNums(reader, 2)
	}
	cnt, res := solve(n, m, E)

	if cnt < 0 {
		fmt.Println("NO")
		return
	}
	var buf bytes.Buffer
	buf.WriteString("YES\n")
	buf.WriteString(fmt.Sprintf("%d\n", cnt))
	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
	}
	fmt.Print(buf.String())
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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
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

func solve(n, m int, E [][]int) (int, [][]int) {
	if m > n {
		return -1, nil
	}

	if n == 1 {
		if m == 0 {
			return 1, [][]int{{1, 1}}
		}
		return 0, nil
	}

	degree := make([]int, n)

	uf := NewUFSet(n)

	for i := 0; i < m; i++ {
		u, v := E[i][0], E[i][1]
		u--
		v--
		degree[u]++
		if degree[u] > 2 {
			return -1, nil
		}

		degree[v]++
		if degree[v] > 2 {
			return -1, nil
		}
		uf.Union(u, v)
	}
	root := uf.Find(0)
	if uf.cnt[root] == n {
		// only one comp

	}

	pair := make([]bool, n)
	for i := 0; i < n; i++ {
		if degree[i] == 2 {
			continue
		}
		if degree[i] == 0 {
			var j = i + 1
			for j < n {
				if degree[j] < 2 && !pair[j] && uf.Find(j) != uf.Find(i) {
					break
				}
				j++
			}

			if j == n {
				return -1, nil
			}
		}
	}
	return -1, nil
}

type Pair struct {
	first, second int
}

type Pairs []Pair

func (ps Pairs) Len() int {
	return len(ps)
}

func (ps Pairs) Less(i, j int) bool {
	return ps[i].first < ps[j].first || ps[i].first == ps[j].first && ps[i].second < ps[j].second
}

func (ps Pairs) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type UFSet struct {
	arr []int
	cnt []int
}

func NewUFSet(n int) UFSet {
	arr := make([]int, n)
	cnt := make([]int, n)

	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UFSet{arr, cnt}
}

func (uf *UFSet) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UFSet) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}

	if uf.cnt[pa] < uf.cnt[pb] {
		pa, pb = pb, pa
	}
	uf.cnt[pa] += uf.cnt[pb]
	uf.arr[pb] = pa
	return true
}
