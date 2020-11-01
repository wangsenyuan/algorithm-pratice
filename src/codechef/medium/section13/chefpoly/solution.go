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
	white := make([][]int, n)
	for i := 0; i < n; i++ {
		white[i] = readNNums(reader, 2)
	}

	black := make([][]int, m)
	for i := 0; i < m; i++ {
		black[i] = readNNums(reader, 2)
	}

	solver := NewSolver(n, white, m, black)

	q := readNum(reader)

	ord := make([]int, n)

	var buf bytes.Buffer

	for q > 0 {
		q--
		var k int
		s, _ := reader.ReadBytes('\n')
		pos := readInt(s, 0, &k)
		for i := 0; i < k; i++ {
			pos = readInt(s, pos+1, &ord[i])
		}
		res := solver.Query(ord[:k])
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

type Solver struct {
	cnt [][]int
}

const MAX_C = 10010

func NewSolver(n int, white [][]int, m int, black [][]int) Solver {
	ws := make([]Point, n)

	for i := 0; i < n; i++ {
		ws[i] = Point{white[i][0] + MAX_C, white[i][1] + MAX_C}
	}

	bs := make([]Point, m)

	for i := 0; i < m; i++ {
		bs[i] = Point{black[i][0] + MAX_C, black[i][1] + MAX_C}
	}

	cnt := make([][]int, n)
	for i := 0; i < n; i++ {
		cnt[i] = make([]int, n)
	}

	var zero Point

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if cw(ws[i], zero, ws[j]) > 0 {
				for k := 0; k < m; k++ {
					if inTriangle(ws[i], zero, ws[j], bs[k]) {
						cnt[i][j]++
					}
				}
				cnt[j][i] = -cnt[i][j]
			}
		}
	}

	return Solver{cnt}
}

func (solver Solver) Query(ord []int) int {
	for i := 0; i < len(ord); i++ {
		ord[i]--
	}

	var ans int

	for i := 0; i < len(ord); i++ {
		j := i + 1
		if j == len(ord) {
			j = 0
		}
		ans += solver.cnt[ord[i]][ord[j]]
	}

	return abs(ans)
}

type Point struct {
	x, y int
}

func cw(a, b, c Point) int {
	return (a.x-b.x)*(c.y-b.y) - (a.y-b.y)*(c.x-b.x)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func inTriangle(a, b, c, d Point) bool {
	bcd := abs(cw(b, c, d))
	acd := abs(cw(a, c, d))
	abd := abs(cw(a, b, d))
	abc := abs(cw(a, b, c))
	return abc == abd+acd+bcd
}
