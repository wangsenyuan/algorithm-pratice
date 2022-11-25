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

	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		A := readNNums(reader, n)
		res := solve(A)
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

// 1024 * 1024
const H = 30

func solve(A []int) int64 {
	// 首先如果有A[i] & L != L => -1
	n := len(A)
	// L 是能改变的
	// L 肯定是
	sort.Ints(A)
	idx := make([]int, H)
	for i := 0; i < H; i++ {
		idx[i] = -1
	}

	uf := NewUF(n)

	var res int64

	for j := 0; j < H && uf.sz > 1; j++ {
		id := -1
		for i := n - 1; i >= 0; i-- {
			if (A[i]>>j)&1 == 1 {
				if id < 0 {
					id = i
				} else {
					pi := uf.Find(i)
					pid := uf.Find(id)
					if pi != pid {
						res += int64(1 << j)
						uf.Union(i, id)
					}
				}
			}
		}
	}

	if uf.sz != 1 {
		return -1
	}
	return res
}

type UF struct {
	arr []int
	cnt []int
	sz  int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &UF{arr, cnt, n}
}

func (uf *UF) Find(x int) int {
	if uf.arr[x] != x {
		uf.arr[x] = uf.Find(uf.arr[x])
	}
	return uf.arr[x]
}

func (uf *UF) Union(a, b int) {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa != pb {
		if uf.cnt[pa] < uf.cnt[pb] {
			pa, pb = pb, pa
		}
		uf.cnt[pa] += uf.cnt[pb]
		uf.arr[pb] = pa
		uf.sz--
	}
}
