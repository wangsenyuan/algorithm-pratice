package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	roads := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		roads[i] = readNNums(reader, 2)
	}
	res := solve(n, roads)
	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))
	for _, cur := range res {
		for i := 0; i < len(cur); i++ {
			buf.WriteString(fmt.Sprintf("%d ", cur[i]))
		}
		buf.WriteByte('\n')
	}

	fmt.Print(buf.String())
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

func solve(n int, roads [][]int) [][]int {
	set := NewUFSet(n)

	var extra [][]int

	for _, r := range roads {
		u, v := r[0]-1, r[1]-1
		if !set.Union(u, v) {
			extra = append(extra, r)
		}
	}

	if len(extra) == 0 {
		return nil
	}

	var comp []int

	for i := 0; i < n; i++ {
		j := set.Find(i)
		if i == j {
			comp = append(comp, i)
		}
	}

	res := make([][]int, len(extra))
	for i, cur := range extra {
		u, v := comp[i], comp[i+1]
		res[i] = []int{cur[0], cur[1], u + 1, v + 1}
	}

	return res
}

type UFSet struct {
	arr  []int
	size []int
	cnt  int
}

func NewUFSet(n int) *UFSet {
	arr := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		size[i] = 1
	}

	return &UFSet{arr, size, n}
}

func (set *UFSet) Find(x int) int {
	if set.arr[x] != x {
		set.arr[x] = set.Find(set.arr[x])
	}
	return set.arr[x]
}

func (set *UFSet) Union(a, b int) bool {
	pa := set.Find(a)
	pb := set.Find(b)
	if pa == pb {
		return false
	}
	set.cnt--
	if set.size[pa] < set.size[pb] {
		pa, pb = pb, pa
	}
	set.size[pa] += set.size[pb]
	set.arr[pb] = pa
	return true
}
