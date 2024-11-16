package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

	fmt.Println(res)
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

func process(reader *bufio.Reader) int {
	n := readNum(reader)
	points := make([][]int, n)
	for i := range n {
		points[i] = readNNums(reader, 2)
	}
	return solve(points)
}

func solve(points [][]int) int {
	n := len(points)

	slices.SortFunc(points, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return a[1] - b[1]
	})

	lines := make(map[pair]map[pair]int)
	// cnt[0] for 垂直线
	// cnt[1] for 水平线
	vertical := make(map[int]int)
	horizontal := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			// dx >= 0
			x1 := points[i][0]
			y1 := points[i][1]
			if dx == 0 {
				// 垂直线
				vertical[x1]++
			} else if dy == 0 {
				// 水平线
				horizontal[y1]++
			} else {
				g := gcd(dx, abs(dy))
				dx /= g
				dy /= g
				// when x = 0, y = ?
				// y1 = dy / dx * x1 + b
				// y1 * dx = dy * x1 + dx + b
				// b = (y1 * dx - dy *x1) / dx

				b := calc(y1*dx-dy*x1, dx)
				k := pair{dx, dy}
				if lines[k] == nil {
					lines[k] = make(map[pair]int)
				}
				lines[k][b]++
			}
		}
	}

	tot := len(vertical) + len(horizontal)

	for _, vs := range lines {
		tot += len(vs)
	}

	res := len(vertical) * (tot - len(vertical))
	res += len(horizontal) * (tot - len(horizontal))

	for _, vs := range lines {
		tmp := len(vs)
		res += tmp * (tot - tmp)
	}

	return res / 2
}

func calc(a, b int) pair {
	g := gcd(abs(a), abs(b))
	a /= g
	b /= g
	return pair{a, b}
}

// y = u/v * x + b
type pair struct {
	first  int
	second int
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func abs(num int) int {
	return max(num, -num)
}

func sign(num int) int {
	if num < 0 {
		return -1
	}
	if num > 0 {
		return 1
	}
	return 0
}
