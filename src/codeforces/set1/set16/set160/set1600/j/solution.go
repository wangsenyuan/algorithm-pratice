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

	n, m := readTwoNums(reader)
	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		mat[i] = readNNums(reader, m)
	}
	res := solve(mat)
	var buf bytes.Buffer

	for _, x := range res {
		buf.WriteString(fmt.Sprintf("%d ", x))
	}
	fmt.Println(buf.String())
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

func readNInt64s(reader *bufio.Reader, n int) []int64 {
	res := make([]int64, n)
	s, _ := reader.ReadBytes('\n')

	var pos int

	for i := 0; i < n; i++ {
		pos = readInt64(s, pos, &res[i]) + 1
	}

	return res
}

func readInt64(bytes []byte, from int, val *int64) int {
	i := from
	var sign int64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var tmp int64
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		tmp = tmp*10 + int64(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
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

func solve(mat [][]int) []int {
	n := len(mat)
	m := len(mat[0])
	uf := NewUF(n * m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := mat[i][j]
			if i > 0 && x&8 == 0 && mat[i-1][j]&2 == 0 {
				uf.Union(i*m+j, (i-1)*m+j)
			}
			if j > 0 && x&1 == 0 && mat[i][j-1]&4 == 0 {
				uf.Union(i*m+j, i*m+j-1)
			}
		}
	}
	var rooms []int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			x := uf.Find(i*m + j)
			if x == i*m+j {
				rooms = append(rooms, uf.count[x])
			}
		}
	}
	sort.Ints(rooms)
	reverse(rooms)

	return rooms
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

type UF struct {
	arr   []int
	count []int
	size  int
}

func NewUF(n int) UF {
	arr := make([]int, n)
	count := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		count[i]++
	}

	return UF{arr, count, n}
}

func (uf *UF) Find(a int) int {
	if uf.arr[a] != a {
		uf.arr[a] = uf.Find(uf.arr[a])
	}
	return uf.arr[a]
}

func (uf *UF) Union(a, b int) {
	pa := uf.Find(a)
	pb := uf.Find(b)
	if pa != pb {
		if uf.count[pa] < uf.count[pb] {
			pa, pb = pb, pa
		}
		uf.count[pa] += uf.count[pb]
		uf.arr[pb] = pa
		uf.size--
	}
}
