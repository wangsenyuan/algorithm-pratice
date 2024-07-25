package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)

	targets := make([]target, n)

	for i := 0; i < n; i++ {
		s := readString(reader)
		targets[i] = parse(s)
	}

	res := solve(targets)

	fmt.Printf("%.7f\n", res)
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

func readFloat64(bytes []byte, from int, val *float64) int {
	i := from
	var sign float64 = 1
	if bytes[i] == '-' {
		sign = -1
		i++
	}
	var real int64

	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		real = real*10 + int64(bytes[i]-'0')
		i++
	}

	if i == len(bytes) || bytes[i] != '.' {
		*val = float64(real)
		return i
	}

	// bytes[i] == '.'
	i++

	var fraq float64
	var base float64 = 0.1
	for i < len(bytes) && bytes[i] >= '0' && bytes[i] <= '9' {
		fraq += base * float64(bytes[i]-'0')
		base /= 10
		i++
	}

	*val = (float64(real) + fraq) * sign

	return i
}

type target struct {
	x int
	y int
	t int
	p float64
}

func parse(s string) target {
	var x, y, t int
	buf := []byte(s)
	pos := readInt(buf, 0, &x)
	pos = readInt(buf, pos+1, &y)
	pos = readInt(buf, pos+1, &t)
	var p float64
	readFloat64(buf, pos+1, &p)
	return target{x, y, t, p}
}

func canReach(a, b target) bool {
	dx := b.x - a.x
	dy := b.y - a.y
	dt := b.t - a.t
	return dx*dx+dy*dy <= dt*dt
}

func abs(num int) int {
	return max(num, -num)
}

func solve(targets []target) float64 {
	n := len(targets)

	sort.Slice(targets, func(i, j int) bool {
		return targets[i].t < targets[j].t
	})

	dp := make([]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = targets[i].p
		for j := 0; j < i; j++ {
			if canReach(targets[j], targets[i]) {
				dp[i] = math.Max(dp[i], dp[j]+targets[i].p)
			}
		}
	}

	return slices.Max(dp)
}
