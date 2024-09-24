package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	pos := readNNums(reader, 4)
	vx, t := readTwoNums(reader)
	v := readNNums(reader, 2)
	w := readNNums(reader, 2)
	res := solve(pos, vx, t, v, w)

	fmt.Printf("%.8f\n", res)
}

func readString(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
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

const iter = 100

func solve(pos []int, vx int, t int, v []int, w []int) float64 {
	x1, y1 := float64(pos[0]), float64(pos[1])
	x2, y2 := float64(pos[2]), float64(pos[3])
	x2 -= x1
	y2 -= y1

	f := func(time float64) bool {
		var dx, dy float64
		if time <= float64(t) {
			dx = time * float64(v[0])
			dy = time * float64(v[1])
		} else {
			dx = float64(t)*float64(v[0]) + (time-float64(t))*float64(w[0])
			dy = float64(t)*float64(v[1]) + (time-float64(t))*float64(w[1])
		}
		dx -= x2
		dy -= y2

		return math.Sqrt(dx*dx+dy*dy) <= time*float64(vx)
	}

	var l, r float64 = 0, 1 << 30

	for i := 0; i < iter; i++ {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid
		}
	}

	return (l + r) / 2
}
