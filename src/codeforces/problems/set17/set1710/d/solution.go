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
		n := readNum(reader)

		good := make([]string, n)
		for i := 0; i < n; i++ {
			good[i] = readString(reader)
		}

		res := solve(n, good)

		for _, edge := range res {
			buf.WriteString(fmt.Sprintf("%d %d\n", edge[0], edge[1]))
		}
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

func normalize(s string) string {

	for i := len(s); i > 0; i-- {
		if s[i-1] >= 'a' && s[i-1] <= 'z' {
			return s[:i]
		}
	}
	return ""
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
func solve(n int, good []string) [][]int {

	lv := make([]int, n+1)
	rv := make([]int, n+1)

	for i := 1; i <= n; i++ {
		lv[i] = i
		rv[i] = i
	}

	tmp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		tmp[i] = make([]int, n)
	}
	pos := make([]int, 2)
	back := func(id int) int {
		return tmp[id][pos[id]-1]
	}

	push := func(id int, v int) {
		tmp[id][pos[id]] = v
		pos[id]++
	}

	reset := func() {
		for i := 0; i < 2; i++ {
			pos[i] = 0
		}
	}

	var edges [][]int

	for L := 2; L <= n; L++ {
		for l := 1; l <= n+1-L; l++ {
			r := l + L - 1
			if good[l-1][r-l] == '1' && lv[l] != lv[r] && rv[l] != rv[r] {
				edges = append(edges, []int{l, r})

				reset()
				push(0, l)
				push(1, r)

				id := 1

				for i := rv[l] + 1; i < lv[r]; i++ {
					if lv[i] == i {
						// the left bound
						edges = append(edges, []int{back(id), i})
						push(id, i)
						id ^= 1
					}
				}
				lm := lv[l]
				rm := rv[r]
				for i := lm; i <= rm; i++ {
					lv[i] = lm
					rv[i] = rm
				}
			}
		}
	}
	return edges
}
