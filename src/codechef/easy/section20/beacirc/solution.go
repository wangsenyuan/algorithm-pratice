package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)

	for tc > 0 {
		tc--
		n := readNum(reader)
		cc := make([][]int, n)
		for i := 0; i < n; i++ {
			cc[i] = readNNums(reader, 3)
		}

		fmt.Println(solve(n, cc))
	}
}

func solve(n int, circles [][]int) int64 {
	cnt := make(map[Circle]int64)

	for i := 0; i < n; i++ {
		cc := Circle{Pair{circles[i][0], circles[i][1]}, circles[i][2]}
		cnt[cc]++
	}
	var res int64

	xx := make([]Circle, 0, len(cnt))

	for c := range cnt {
		xx = append(xx, c)
	}

	for i := 0; i < len(xx); i++ {
		a := xx[i]
		for j := i + 1; j < len(xx); j++ {
			b := xx[j]
			if (a.center.first+b.center.first)%2 == 1 || (a.center.second+b.center.second)%2 == 1 {
				continue
			}

			dx := float64(b.center.first - a.center.first)
			dy := float64(b.center.second - a.center.second)
			ar := float64(a.radius)
			br := float64(b.radius)
			dd := dx*dx + dy*dy

			if dd != ar*ar+br*br {
				continue
			}

			d := math.Sqrt(dd)

			if d*d != dd {
				continue
			}

			di := int(d)

			if math.Abs(float64(di)-d) > 1e-6 {
				continue
			}

			if di >= a.radius+b.radius {
				continue
			}

			if di <= abs(a.radius-b.radius) {
				continue
			}

			if di%2 == 1 {
				continue
			}

			x := (a.center.first + b.center.first) / 2
			y := (a.center.second + b.center.second) / 2

			c := Circle{Pair{x, y}, di / 2}

			res += cnt[a] * cnt[b] * cnt[c]
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

type Pair struct {
	first, second int
}

type Circle struct {
	center Pair
	radius int
}
