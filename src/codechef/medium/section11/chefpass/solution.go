package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		a, _ := reader.ReadBytes('\n')
		var n, m int
		pos := readInt(a, 0, &n)
		A := a[pos+1:]
		b, _ := reader.ReadBytes('\n')
		pos = readInt(b, 0, &m)
		B := b[pos+1:]
		res := solve(n, A, m, B)

		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
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

func solve(n int, A []byte, m int, B []byte) bool {
	degree := make([]int, 52)

	for i, j := 0, 0; i < n; i++ {
		a := A[j]
		b := A[j+1]
		degree[ind(a)]++
		degree[ind(b)]++
		j += 3
	}

	uf := NewUF(52)

	for i, j := 0, 0; i < m; i++ {
		a := B[j]
		b := B[j+1]
		uf.Union(ind(a), ind(b))
		j += 3
	}

	ovc := make([]int, 52)

	for i := 0; i < 52; i++ {
		if degree[i]&1 == 1 {
			ovc[uf.Find(i)]++
		}
	}

	for i := 0; i < 52; i++ {
		if ovc[i]&1 == 1 {
			return false
		}
	}
	return true
}

func ind(b byte) int {
	if b >= 'a' && b <= 'z' {
		return int(b - 'a')
	}
	return int(b-'A') + 26
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i] = 1
	}

	return UF{arr, cnt}
}

func (uf *UF) Find(u int) int {
	if uf.arr[u] == u {
		return u
	}
	uf.arr[u] = uf.Find(uf.arr[u])

	return uf.arr[u]
}

func (uf *UF) Union(u, v int) bool {
	pu := uf.Find(u)
	pv := uf.Find(v)
	if pu == pv {
		return false
	}
	if uf.cnt[pu] < uf.cnt[pv] {
		pu, pv = pv, pu
	}
	uf.cnt[pu] += uf.cnt[pv]
	uf.arr[pv] = pu
	return true
}
