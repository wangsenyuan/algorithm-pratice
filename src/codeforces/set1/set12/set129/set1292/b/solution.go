package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	first := readNNums(reader, 6)
	xs, ys, t := readThreeNums(reader)

	res := solve(first[0], first[1], first[2], first[3], first[4], first[5], xs, ys, t)

	fmt.Println(res)
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

func solve(x0, y0 int, ax int, ay int, bx int, by int, xs int, ys int, t int) int {
	type pair struct {
		first  int
		second int
	}

	arr := []pair{{x0, y0}}

	dist := func(a, b pair) int {
		dx := abs(a.first - b.first)
		dy := abs(a.second - b.second)
		if dx+dy < 0 {
			return -1
		}
		return dx + dy
	}

	a := big.NewInt(int64(ax))
	b := big.NewInt(int64(ay))

	inf := big.NewInt(math.MaxInt64)

	check := func(v int64, a *big.Int) bool {
		x := big.NewInt(v)
		x = x.Mul(x, a)
		return x.Cmp(inf) > 0
	}

	checkOverflow := func(v pair) bool {
		return check(int64(v.first), a) || check(int64(v.second), b)
	}

	origin := pair{xs, ys}

	prev := arr[0]

	for k := 0; k < 100 && !checkOverflow(prev); k++ {
		// 可能会益处的
		x, y := ax*prev.first+bx, ay*prev.second+by

		if ax*prev.first < 0 || x < 0 ||
			ay*prev.second < 0 || y < 0 {
			break
		}

		prev = pair{x, y}
		arr = append(arr, prev)
	}

	// 肯定是连续的一段
	var best int

	for l := 0; l < len(arr); l++ {
		var sum int
		for r := l; r < len(arr); r++ {
			// 能否访问这整个区间
			if r > l {
				tmp := dist(arr[r-1], arr[r])
				if tmp < 0 {
					break
				}
				sum += tmp
			}
			if sum > t {
				break
			}
			tmp := min(dist(origin, arr[l]), dist(origin, arr[r]))

			if tmp >= 0 && sum+tmp <= t {
				best = max(best, r-l+1)
			}
		}
	}
	return best
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
