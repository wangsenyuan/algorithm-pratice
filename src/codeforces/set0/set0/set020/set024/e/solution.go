package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res < 0 {
		fmt.Println("-1")
		return
	}
	fmt.Printf("%.10f\n", res)
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

func process(reader *bufio.Reader) float64 {
	n := readNum(reader)
	particles := make([][]int, n)
	for i := 0; i < n; i++ {
		particles[i] = readNNums(reader, 2)
	}
	return solve(particles)
}

const inf = 1 << 60

const eps = 1e-9

func solve(particles [][]int) float64 {
	n := len(particles)
	// x[0] < x[1] ... < x[n-1]
	// v[0] > v[1] ... > v[n-1]
	check := func(t float64) bool {
		// 找到最大的 s2 + v2 * t (v2 > 0) 的部分
		var best float64 = -inf
		for i := 0; i < n; i++ {
			x := particles[i][0]
			if particles[i][1] < 0 {
				// 它是向左的
				v := -particles[i][1]
				// s1 - v1 * t <= best
				tmp := float64(x) - float64(v)*t
				if tmp < best || math.Abs(tmp-best) < eps {
					return true
				}
			} else {
				v := particles[i][1]
				tmp := float64(x) + float64(v)*t
				best = math.Max(best, tmp)
			}
		}
		return false
	}

	var l, r float64

	prev := -1
	for i := range n {
		if particles[i][1] > 0 {
			prev = i
		} else if prev >= 0 {
			dist := float64(particles[i][0] - particles[prev][0])
			v := float64(particles[prev][1] - particles[i][1])
			r = dist / v
			break
		}
	}
	if r == 0 {
		// 没有可以相遇的
		return -1
	}

	for range 60 {
		mid := (l + r) / 2
		if check(mid) {
			r = mid
		} else {
			l = mid
		}
		if (r - l) < eps {
			break
		}
	}

	return (l + r) / 2
}
