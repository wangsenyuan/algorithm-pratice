package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	persons := make([][]int, 3)
	for i := 0; i < 3; i++ {
		persons[i] = readNNums(reader, 3)
	}
	res := solve(persons)
	fmt.Println(res)
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

func solve(persons [][]int) int {
	// 每个人都可以move, lift, throw 也就是最多只有9个操作
	// a (move and ) lift b, c (move and ) lift a & b, and (move and ) throw a & b, (a throw b) and b move

	p := make([]int, 3)
	m := make([]int, 3)
	t := make([]int, 3)

	for i := 0; i < 3; i++ {
		p[i] = persons[i][0]
		m[i] = persons[i][1]
		t[i] = persons[i][2]
	}

	check := func(i int, p []int) bool {
		return i != p[0] && i != p[1] && i != p[2]
	}

	var f [48][48][48][8][8]bool

	var ans int

	var dfs func(p []int, x int, y int)
	dfs = func(p []int, x int, y int) {
		if f[p[0]][p[1]][p[2]][x][y] {
			return
		}
		f[p[0]][p[1]][p[2]][x][y] = true
		far, near := 0, 50
		for i := 0; i < 3; i++ {
			if p[i] < 45 && p[i] > far {
				far = p[i]
			}
			if p[i] < near {
				near = p[i]
			}
		}

		ans = max(ans, far)

		for j := 0; j < 3; j++ {
			q := []int{p[0], p[1], p[2]}
			if p[j] >= 45 {
				T := p[j] - 45
				for i := p[T] + t[T]; p[T] < 45 && i >= p[T]-t[T] && i > near-2 && i > 0; i-- {
					if check(i, p) {
						q[j] = i
						dfs(q, x, y)
					}
				}
			} else if check(j+45, p) {
				z := 1 << j
				if y&z > 0 {
					for i := 0; i < 3; i++ {
						if abs(p[i]-p[j]) == 1 {
							q[i] = 45 + j
							dfs(q, x, y^z)
							q[i] = p[i]
						}
					}
				}

				if x&z > 0 {
					for i := p[j] + m[j]; i >= p[j]-m[j] && i > near-2 && i > 0; i-- {
						if check(i, p) {
							q[j] = i
							dfs(q, x^z, y)
						}
					}
				}
			}
		}
	}

	dfs(p, 7, 7)

	return ans
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
