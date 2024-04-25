package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	a := readNNums(reader, 2)
	b := readNNums(reader, 2)
	c := readNNums(reader, 2)

	res := solve(a, b, c)

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("%d\n", len(res)))

	for i := 0; i < len(res); i++ {
		buf.WriteString(fmt.Sprintf("%d %d\n", res[i][0], res[i][1]))
	}

	fmt.Println(buf.String())
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadBytes('\n')
	for i := 0; i < len(s); i++ {
		if s[i] == '\n' || s[i] == '\r' {
			return string(s[:i])
		}
	}
	return string(s)
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

func solve(a []int, b []int, c []int) [][]int {
	pts := [][]int{a, b, c}
	slices.SortFunc(pts, func(x, y []int) int {
		if x[0] != y[0] {
			return x[0] - y[0]
		}
		return x[1] - y[1]
	})

	a, b, c = pts[0], pts[1], pts[2]
	var cells [][]int

	same := func(u, v []int) bool {
		return u[0] == v[0] && u[1] == v[1]
	}

	addUntil := func(cur []int, dest []int) {
		if cur[0] != dest[0] {
			dx := 1
			if cur[0] > dest[0] {
				dx = -1
			}
			// 水平+
			for i := cur[0]; ; i += dx {
				tmp := []int{i, cur[1]}

				if len(cells) > 0 && same(cells[len(cells)-1], tmp) {
					continue
				}
				cells = append(cells, tmp)
				if i == dest[0] {
					break
				}
			}
		} else if cur[1] != dest[1] {
			dy := 1
			if cur[1] > dest[1] {
				dy = -1
			}
			// 垂直
			for i := cur[1]; ; i += dy {
				tmp := []int{cur[0], i}

				if len(cells) > 0 && same(cells[len(cells)-1], tmp) {
					continue
				}
				cells = append(cells, tmp)
				if i == dest[1] {
					break
				}
			}
		} else if len(cells) == 0 || !same(cells[len(cells)-1], dest) {
			cells = append(cells, dest)
		}
	}

	y0 := min(a[1], c[1])
	y1 := max(a[1], c[1])

	if b[1] >= y0 && b[1] <= y1 {
		// b[0], b[1]
		addUntil(a, []int{b[0], a[1]})
		addUntil([]int{b[0], a[1]}, []int{b[0], c[1]})
		addUntil([]int{b[0], c[1]}, c)

		return cells
	}

	if b[1] < y0 {
		if a[1] == y0 {
			addUntil(a, []int{c[0], y0})
			addUntil([]int{c[0], y0}, c)
		} else {
			addUntil(c, []int{a[0], y0})
			addUntil([]int{a[0], y0}, a)
		}
		addUntil(b, []int{b[0], y0 - 1})

	} else {
		if a[1] == y1 {
			addUntil(a, []int{c[0], y1})
			addUntil([]int{c[0], y1}, c)
		} else {
			addUntil(c, []int{a[0], y1})
			addUntil([]int{a[0], y1}, a)
		}
		addUntil(b, []int{b[0], y1 + 1})
	}

	return cells
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
