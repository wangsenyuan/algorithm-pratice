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
		n, m, s := readThreeNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, m)
		}
		res := solve(n, m, s, A)

		buf.WriteString(fmt.Sprintf("%d\n", res))
	}
	fmt.Print(buf.String())
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
		if s[i] == '\n' {
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

func solve(n int, m, s int, A [][]int) int64 {
	items := make([]Item, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			items[i*m+j] = Item{i, j, A[i][j]}
		}
	}

	sort.Sort(Items(items))

	N := n + m - 2
	f1 := make([]int64, N)
	f2 := make([]int64, N)
	f3 := make([]int64, N)
	f4 := make([]int64, N)

	var ans int64

	for _, it := range items {
		v := it.v
		x := it.x
		y := it.y
		f12 := x + y
		f34 := x - y + m - 1
		var tmp int64
		tmp = max(tmp, get(f1, f12-s))
		tmp = max(tmp, get(f2, n-1+m-1-f12-s))
		tmp = max(tmp, get(f3, f34-s))
		tmp = max(tmp, get(f4, n-1+m-1-f34-s))
		tmp += int64(v)
		ans = max(ans, tmp)
		update(f1, f12, tmp)
		update(f2, n-1+m-1-f12, tmp)
		update(f3, f34, tmp)
		update(f4, n-1+m-1-f34, tmp)
	}

	return ans
}

func update(arr []int64, p int, v int64) {
	for p < len(arr) {
		arr[p] = max(arr[p], v)
		p |= (p + 1)
	}
}

func get(arr []int64, p int) int64 {
	var res int64
	for p > 0 {
		res = max(res, arr[p-1])
		p &= (p - 1)
	}
	return res
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

type Item struct {
	x int
	y int
	v int
}

type Items []Item

func (this Items) Len() int {
	return len(this)
}

func (this Items) Less(i, j int) bool {
	return this[i].v > this[j].v
}

func (this Items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
