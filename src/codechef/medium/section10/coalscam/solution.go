package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		first := readNNums(reader, 3)
		N, M1, M2 := first[0], first[1], first[2]
		roads := make([][]int, M1+M2)
		for i := 0; i < len(roads); i++ {
			roads[i] = readNNums(reader, 3)
		}
		a, b := solve(N, M1, M2, roads)

		if a == -1 {
			fmt.Println("Impossible")
		} else {
			fmt.Printf("%d %d\n", a, b)
		}
	}
}

func solve(N int, M1 int, M2 int, road [][]int) (int64, int64) {
	roads1 := make([]Road, M1)
	roads2 := make([]Road, M2)

	for i := 0; i < len(road); i++ {
		u, v, c := road[i][0], road[i][1], road[i][2]
		cur := Road{u, v, c}
		if i < M1 {
			roads1[i] = cur
		} else {
			roads2[i-M1] = cur
		}
	}

	sort.Sort(Roads1(roads1))
	sort.Sort(Roads2(roads2))

	uf := NewUF(N)
	var profit int64
	for i := 0; i < M2; i++ {
		cur := roads2[i]
		if uf.Union(cur.first, cur.second) {
			profit += int64(cur.cost)
		}
	}

	total := profit

	for i := 0; i < M1; i++ {
		cur := roads1[i]
		if uf.Union(cur.first, cur.second) {
			total += int64(cur.cost)
		}
	}

	if uf.Size() > 1 {
		return -1, -1
	}
	return profit, total
}

type Road struct {
	first, second int
	cost          int
}

type Roads1 []Road

func (roads Roads1) Len() int {
	return len(roads)
}

func (roads Roads1) Swap(i, j int) {
	roads[i], roads[j] = roads[j], roads[i]
}

func (roads Roads1) Less(i, j int) bool {
	return roads[i].cost < roads[j].cost
}

type Roads2 []Road

func (roads Roads2) Len() int {
	return len(roads)
}

func (roads Roads2) Swap(i, j int) {
	roads[i], roads[j] = roads[j], roads[i]
}

func (roads Roads2) Less(i, j int) bool {
	return roads[i].cost > roads[j].cost
}

type UF struct {
	set   []int
	count []int
	size  int
}

func NewUF(n int) UF {
	set := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		set[i] = i
		count[i] = 1
	}
	return UF{set, count, n}
}

func (uf *UF) Find(x int) int {
	if uf.set[x] == x {
		return x
	}
	uf.set[x] = uf.Find(uf.set[x])
	return uf.set[x]
}

func (uf *UF) Union(a, b int) bool {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa == pb {
		return false
	}
	if uf.count[pa] >= uf.count[pb] {
		uf.set[pb] = pa
		uf.count[pa] += uf.count[pb]
	} else {
		uf.set[pa] = pb
		uf.count[pb] += uf.count[pa]
	}
	uf.size--
	return true
}

func (uf UF) Size() int {
	return uf.size
}
