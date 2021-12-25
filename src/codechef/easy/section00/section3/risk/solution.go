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
		B := make([]string, n)
		for i := 0; i < n; i++ {
			B[i], _ = reader.ReadString('\n')
		}
		res := solve(n, m, B)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

var d = []int{-1, 0, 1, 0, -1}

func solve(n, m int, B []string) int {
	uf := NewUF(n * m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if B[i][j] == '0' {
				continue
			}
			for k := 0; k < 4; k++ {
				x, y := i+d[k], j+d[k+1]
				if x >= 0 && x < n && y >= 0 && y < m && B[x][y] == '1' {
					uf.Union(i*m+j, x*m+y)
				}
			}
		}
	}
	sets := make([]int, 0, uf.size)

	for i := 0; i < n*m; i++ {
		if uf.arr[i] == i && B[i/m][i%n] == '1' {
			sets = append(sets, uf.count[i])
		}
	}

	sort.Ints(sets)

	var res int

	for i := len(sets) - 2; i >= 0; i -= 2 {
		res += sets[i]
	}

	return res
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
