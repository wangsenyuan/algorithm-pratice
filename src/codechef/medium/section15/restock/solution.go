package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	tc := readNum(reader)
	var buf bytes.Buffer
	for tc > 0 {
		tc--
		m, n := readTwoNums(reader)
		D, R, C := readThreeNums(reader)
		G := make([][]int, m)
		for i := 0; i < m; i++ {
			G[i] = readNNums(reader, n)
		}
		res := solve(R, C, D, G)
		buf.WriteString(fmt.Sprintf("%d\n", res))
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

const INF = 1 << 30

func solve(R, C, D int, G [][]int) int {

	m := len(G)
	n := len(G[0])

	T := make([][]int, 4*m)
	for i := 0; i < len(T); i++ {
		T[i] = make([]int, 4*n)
		for j := 0; j < len(T[i]); j++ {
			T[i][j] = INF
		}
	}

	var update_x func(vx int, lx int, rx int, x int, y int, v int)
	var update_y func(vx int, lx int, rx int, vy int, ly int, ry int, x int, y int, v int)

	update_x = func(vx, lx, rx, x, y, v int) {
		if lx != rx {
			mx := (lx + rx) / 2
			if x <= mx {
				update_x(vx*2, lx, mx, x, y, v)
			} else {
				update_x(vx*2+1, mx+1, rx, x, y, v)
			}
		}
		update_y(vx, lx, rx, 1, 0, n-1, x, y, v)
	}

	update_y = func(vx, lx, rx, vy, ly, ry, x, y, v int) {
		if ly == ry {
			if lx == rx {
				T[vx][vy] = min(T[vx][vy], v)
			} else {
				T[vx][vy] = min(T[vx*2][vy], T[vx*2+1][vy])
			}
			return
		}
		my := (ly + ry) / 2
		if y <= my {
			update_y(vx, lx, rx, vy*2, ly, my, x, y, v)
		} else {
			update_y(vx, lx, rx, vy*2+1, my+1, ry, x, y, v)
		}
		T[vx][vy] = min(T[vx][vy*2], T[vx][vy*2+1])
	}

	var get_x func(vx int, tlx int, trx int, lx int, rx int, ly int, ry int) int
	var get_y func(vx int, vy int, tly int, try int, ly int, ry int) int

	get_x = func(vx, tlx, trx, lx, rx, ly, ry int) int {
		if lx > rx {
			return INF
		}
		if lx == tlx && rx == trx {
			return get_y(vx, 1, 0, n-1, ly, ry)
		}
		tmx := (tlx + trx) / 2
		a := get_x(vx*2, tlx, tmx, lx, min(rx, tmx), ly, ry)
		b := get_x(vx*2+1, tmx+1, trx, max(tmx+1, lx), rx, ly, ry)
		return min(a, b)
	}

	get_y = func(vx, vy, tly, try, ly, ry int) int {
		if ly > ry {
			return INF
		}
		if ly == tly && ry == try {
			return T[vx][vy]
		}

		tmy := (tly + try) / 2

		a := get_y(vx, vy*2, tly, tmy, ly, min(tmy, ry))
		b := get_y(vx, vy*2+1, tmy+1, try, max(ly, tmy+1), ry)
		return min(a, b)
	}

	var jobs []Job

	cost := make([][]int, m)

	for i := 0; i < m; i++ {
		cost[i] = make([]int, n)
		for j := 0; j < n; j++ {
			job := Job{i*i + j*j, i, j}
			jobs = append(jobs, job)
		}
	}

	update_x(1, 0, m-1, 0, 0, 0)

	sort.Sort(Jobs(jobs))

	for i, j := 0, 0; i < len(jobs); i = j {
		for j = i; j < len(jobs) && jobs[j].dist == jobs[i].dist; j++ {
			r := jobs[j].i
			c := jobs[j].j
			x1 := max(0, r-D)
			x2 := min(m-1, r+D)
			y1 := max(0, c-D)
			y2 := min(n-1, c+D)
			v := get_x(1, 0, m-1, x1, x2, y1, y2)
			v += G[r][c]
			cost[r][c] = v
		}
		for j = i; j < len(jobs) && jobs[j].dist == jobs[i].dist; j++ {
			r := jobs[j].i
			c := jobs[j].j
			update_x(1, 0, m-1, r, c, cost[r][c])
		}
	}

	return cost[R][C]
}

type Job struct {
	dist int
	i    int
	j    int
}

type Jobs []Job

func (jobs Jobs) Len() int {
	return len(jobs)
}

func (jobs Jobs) Less(i, j int) bool {
	return jobs[i].dist < jobs[j].dist
}

func (jobs Jobs) Swap(i, j int) {
	jobs[i], jobs[j] = jobs[j], jobs[i]
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
