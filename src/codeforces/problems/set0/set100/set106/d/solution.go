package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, m := readTwoNums(reader)
	mat := make([]string, n)
	for i := 0; i < n; i++ {
		mat[i] = readString(reader)[:m]
	}
	k := readNum(reader)
	insts := make([]string, k)
	for i := 0; i < k; i++ {
		insts[i] = readString(reader)
	}
	res := solve(mat, insts)
	if len(res) == 0 {
		fmt.Println("no solution")
	} else {
		fmt.Println(res)
	}
}

func readFloat64(bs []byte, pos int, r *float64) int {
	var first float64
	var second float64
	exp := 1.0
	var dot bool
	for pos < len(bs) && (bs[pos] == '.' || isDigit(bs[pos])) {
		if bs[pos] == '.' {
			dot = true
		} else if !dot {
			first = first*10 + float64(bs[pos]-'0')
		} else {
			second = second*10 + float64(bs[pos]-'0')
			exp *= 10
		}
		pos++
	}
	*r = first + second/exp
	//fmt.Fprintf(os.Stderr, "%.6f ", *r)
	return pos
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
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

func solve(mat []string, instructions []string) string {
	var res []byte

	m := len(mat)
	n := len(mat[0])

	ins := make([]Instruction, len(instructions))

	for i, s := range instructions {
		ins[i] = ParseInstruction(s)
	}

	sum := make([][]int, m)
	for i := 0; i < m; i++ {
		sum[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == '#' {
				sum[i][j] = 1
			}
			if i > 0 {
				sum[i][j] += sum[i-1][j]
			}
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
			if i > 0 && j > 0 {
				sum[i][j] -= sum[i-1][j-1]
			}
		}
	}

	canReach := func(x int, y int, u int, v int) bool {
		if u >= m || u < 0 || v >= n || v < 0 {
			return false
		}
		x, u = min(x, u), max(x, u)
		y, v = min(y, v), max(y, v)

		tmp := sum[u][v]
		if x > 0 {
			tmp -= sum[x-1][v]
		}
		if y > 0 {
			tmp -= sum[u][y-1]
		}
		if x > 0 && y > 0 {
			tmp += sum[x-1][y-1]
		}
		return tmp == 0
	}

	check := func(x int, y int) bool {
		// start from x, y
		for _, in := range ins {
			u, v := x, y
			if in.dir == "N" {
				u -= in.steps
			} else if in.dir == "E" {
				v += in.steps
			} else if in.dir == "S" {
				u += in.steps
			} else {
				v -= in.steps
			}
			if !canReach(x, y, u, v) {
				return false
			}
			x, y = u, v
		}
		return true
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] >= 'A' && mat[i][j] <= 'Z' {
				if check(i, j) {
					res = append(res, mat[i][j])
				}
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})

	return string(res)
}

type Instruction struct {
	dir   string
	steps int
}

func ParseInstruction(s string) Instruction {
	dir := s[:1]
	steps, _ := strconv.Atoi(s[2:])
	return Instruction{dir, steps}
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
