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
		x := readNNums(reader, n)
		direction := readString(reader)
		res := solve(m, x, direction)
		for i := 0; i < n; i++ {
			buf.WriteString(fmt.Sprintf("%d ", res[i]))
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

func solve(m int, x []int, direction string) []int {
	n := len(x)

	type robot struct {
		id  int
		pos int
		dir int
	}
	robots := make([]robot, n)
	for i := 0; i < n; i++ {
		var d int
		if direction[2*i] == 'R' {
			d = 1
		}
		robots[i] = robot{i, x[i], d}
	}

	sort.Slice(robots, func(i, j int) bool {
		return robots[i].pos < robots[j].pos
	})

	res := make([]int, n)

	stack := make([]robot, n)

	process := func(arr []robot) {
		var top int
		for _, r := range arr {
			if r.dir == 1 {
				// move right
				stack[top] = r
				top++
			} else {
				// move left
				if top == 0 {
					stack[top] = r
					top++
				} else {
					x := stack[top-1]
					top--
					dist := r.pos - x.pos

					if x.dir == 0 {
						dist += 2 * x.pos
					}
					res[r.id] = dist / 2
					res[x.id] = dist / 2
				}
			}
		}

		for top > 1 {
			// 它肯定是move right 的
			r := stack[top-1]
			l := stack[top-2]
			dist := m - r.pos
			if l.dir == 1 {
				dist += m - l.pos
			} else {
				dist += l.pos + m
			}

			res[r.id] = dist / 2
			res[l.id] = dist / 2

			top -= 2
		}
		if top > 0 {
			res[stack[top-1].id] = -1
		}
	}

	arr := make([][]robot, 2)
	for _, r := range robots {
		arr[r.pos%2] = append(arr[r.pos%2], r)
	}

	process(arr[0])
	process(arr[1])

	return res
}
