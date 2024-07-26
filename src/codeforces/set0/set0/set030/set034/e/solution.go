package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, x := readTwoNums(reader)
	ss := make([][]int, n)
	for i := 0; i < n; i++ {
		ss[i] = readNNums(reader, 3)
	}
	res := solve(ss, x)
	var buf bytes.Buffer
	for i := 0; i < n; i++ {
		buf.WriteString(fmt.Sprintf("%.6f\n", res[i]))
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

type ball struct {
	x  float64
	v  float64
	m  float64
	id int
}

func (a ball) update(x float64, v float64) ball {
	return ball{x, v, a.m, a.id}
}

const inf = 1 << 50

const eps = 1e-10

func solve(ss [][]int, t int) []float64 {
	n := len(ss)
	balls := make([]ball, n)
	for i := 0; i < n; i++ {
		balls[i] = ball{float64(ss[i][0]), float64(ss[i][1]), float64(ss[i][2]), i}
	}

	update := func(dt float64) {
		for i := 0; i < n; i++ {
			if i+1 < n {
				tmp := (balls[i].x - balls[i+1].x) / (balls[i+1].v - balls[i].v)
				if tmp == dt || math.Abs(tmp-dt) < eps {
					x := balls[i].x + balls[i].v*dt
					v1 := ((balls[i].m-balls[i+1].m)*balls[i].v + 2*balls[i+1].m*balls[i+1].v) / (balls[i].m + balls[i+1].m)
					v2 := ((balls[i+1].m-balls[i].m)*balls[i+1].v + 2*balls[i].m*balls[i].v) / (balls[i].m + balls[i+1].m)
					balls[i] = balls[i].update(x, v1)
					balls[i+1] = balls[i+1].update(x, v2)
					i++
					continue
				}
			}
			x := balls[i].x + balls[i].v*dt
			balls[i] = balls[i].update(x, balls[i].v)
		}
	}

	sort.Slice(balls, func(i, j int) bool {
		return balls[i].x < balls[j].x
	})

	var tt = float64(t)
	var cur float64

	for cur < tt {
		// 找出最近的相撞的时间
		var dt float64 = tt - cur

		for i := 1; i < n; i++ {
			a := balls[i-1]
			b := balls[i]
			tmp := (a.x - b.x) / (b.v - a.v)
			if tmp > 0 {
				dt = math.Min(dt, tmp)
			}
		}
		update(dt)
		cur += dt
	}

	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[balls[i].id] = balls[i].x
	}
	return res
}
