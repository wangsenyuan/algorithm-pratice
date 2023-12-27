package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"os"
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

	n := readNum(reader)
	V := readNNums(reader, n)
	T := readNNums(reader, n)

	res := solve(V, T)

	var buf bytes.Buffer

	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%d ", res[i]))
	}
	buf.WriteByte('\n')

	fmt.Print(buf.String())
}

func solve(V []int, T []int) []int {
	n := len(V)

	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + T[i]
	}

	find := func(v int, i int) int {
		// pref[j] - pref[i] >= v
		l, r := i+1, n+1
		for l < r {
			mid := (l + r) / 2
			if pref[mid]-pref[i] >= v {
				r = mid
			} else {
				l = mid + 1
			}
		}
		return r
	}

	ans := make([]int, n)
	men := make(Snowmen, 0, n)

	for i := 0; i < n; i++ {
		j := find(V[i], i)
		quitVol := V[i] - (pref[j-1] - pref[i])
		heap.Push(&men, &Snowman{i, j - 1, quitVol, -1})

		for men.Len() > 0 && men[0].quitAt == i {
			cur := heap.Pop(&men).(*Snowman)
			ans[i] += cur.quitVol
		}
		ans[i] += men.Len() * T[i]
	}

	return ans
}

type Snowman struct {
	id      int
	quitAt  int
	quitVol int
	index   int
}

type Snowmen []*Snowman

func (set Snowmen) Len() int {
	return len(set)
}

func (set Snowmen) Less(i, j int) bool {
	return set[i].quitAt < set[j].quitAt
}

func (set Snowmen) Swap(i, j int) {
	set[i], set[j] = set[j], set[i]
	set[i].index = i
	set[j].index = j
}

func (set *Snowmen) Push(x interface{}) {
	it := x.(*Snowman)
	it.index = len(*set)
	*set = append(*set, it)
}

func (set *Snowmen) Pop() interface{} {
	old := *set
	n := len(old)
	res := old[n-1]
	res.index = -1
	*set = old[:n-1]
	return res
}
