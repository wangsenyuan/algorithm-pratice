package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n, q := readTwoNums(reader)
		solver := NewSolver(n)
		for n > 1 {
			reader.ReadString('\n')
			n--
		}
		for q > 0 {
			q--
			s, _ := reader.ReadBytes('\n')
			if s[0] == '1' {
				u, v, w := readUpdate(s[2:])
				if solver.Update(u, v, w) {
					buf.WriteString("AC\n")
				} else {
					buf.WriteString("WA\n")
				}
			} else {
				u, v := readQuery(s[2:])
				res := solver.Query(u, v)
				buf.WriteString(fmt.Sprintf("%d\n", res))
			}
		}
	}

	fmt.Print(buf.String())
}

func readQuery(s []byte) (u, v int) {
	pos := readInt(s, 0, &u)
	readInt(s, pos+1, &v)
	return
}

func readUpdate(s []byte) (u, v, w int) {
	pos := readInt(s, 0, &u)
	pos = readInt(s, pos+1, &v)
	readInt(s, pos+1, &w)
	return
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

type Solver struct {
	uf   *UF
	arrs []map[int]int
}

func NewSolver(n int) *Solver {
	uf := NewUF(n)
	arrs := make([]map[int]int, n)
	for i := 0; i < n; i++ {
		arrs[i] = make(map[int]int)
		arrs[i][i] = 0
	}
	return &Solver{uf, arrs}
}

func (solver *Solver) Update(u, v, w int) bool {
	u--
	v--
	pu := solver.uf.Find(u)
	pv := solver.uf.Find(v)
	if pu == pv {
		// same component
		arr := solver.arrs[pu]
		return arr[u]^arr[v] == w
	}
	// pu != pv
	if solver.uf.cnt[pu] < solver.uf.cnt[pv] {
		// swap, make sure cnt[pu] >= cnt[pv]
		pu, pv = pv, pu
		u, v = v, u
	}
	ww := solver.arrs[pv][v] ^ w ^ solver.arrs[pu][u]
	for x, y := range solver.arrs[pv] {
		solver.arrs[pu][x] = y ^ ww
	}
	solver.uf.cnt[pu] += solver.uf.cnt[pv]
	solver.uf.arr[pv] = pu
	return true
}

func (solver *Solver) Query(u, v int) int {
	u--
	v--
	pu := solver.uf.Find(u)
	pv := solver.uf.Find(v)

	if pu != pv {
		return -1
	}
	return solver.arrs[pu][u] ^ solver.arrs[pu][v]
}

type UF struct {
	arr []int
	cnt []int
}

func NewUF(n int) *UF {
	arr := make([]int, n)
	cnt := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		cnt[i]++
	}
	return &UF{arr, cnt}
}

func (uf *UF) Find(u int) int {
	if uf.arr[u] == u {
		return u
	}
	uf.arr[u] = uf.Find(uf.arr[u])
	return uf.arr[u]
}
