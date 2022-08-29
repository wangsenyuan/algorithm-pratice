package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var buf bytes.Buffer

	for {
		n := readNum(reader)
		if n == 0 {
			break
		}
		points := make([][]int, n)
		for i := 0; i < n; i++ {
			points[i] = readNNums(reader, 2)
		}

		res := solve(points)

		buf.WriteString(fmt.Sprintf("%.6f\n", res))
	}

	fmt.Print(buf.String())
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

const INF = 1000000100
const eps = 1e-8

func solve(points [][]int) float64 {
	min_x := INF
	max_x := -INF
	min_y := INF
	max_y := -INF

	for _, point := range points {
		x, y := point[0], point[1]
		min_x = min(min_x, x)
		max_x = max(max_x, x)
		min_y = min(min_y, y)
		max_y = max(max_y, y)
	}

	calc := func(x, y int) float64 {
		var sum float64

		for _, point := range points {
			dx := float64(point[0] - x)
			dy := float64(point[1] - y)
			sum += math.Sqrt(dx*dx + dy*dy)
		}

		return sum
	}

	calc_x := func(x int) float64 {
		lo, hi := min_y, max_y
		best := math.MaxFloat64
		for lo <= hi {
			if hi-lo < 3 {
				for y := lo; y <= hi; y++ {
					best = math.Min(best, calc(x, y))
				}
				break
			}

			m1 := (hi - lo) / 3
			m2 := hi - m1
			m1 += lo

			d1 := calc(x, m1)
			d2 := calc(x, m2)

			if d1 <= d2+eps {
				if d1 < best-eps {
					best = d1
				}
				hi = m2 - 1
			} else {
				if d2 < best-eps {
					best = d2
				}
				lo = m1 + 1
			}
		}

		return best
	}

	lo, hi := min_x, max_x

	var best float64 = math.MaxFloat64

	for lo <= hi {
		if hi-lo < 3 {
			for x := lo; x <= hi; x++ {
				best = math.Min(best, calc_x(x))
			}
			break
		}
		m1 := (hi - lo) / 3
		m2 := hi - m1
		m1 += lo
		d1 := calc_x(m1)
		d2 := calc_x(m2)
		if d1 <= d2+eps {
			if d1 < best-eps {
				best = d1
			}
			hi = m2 - 1
		} else {
			if d2 < best-eps {
				best = d2
			}
			lo = m1 + 1
		}
	}

	return best
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
