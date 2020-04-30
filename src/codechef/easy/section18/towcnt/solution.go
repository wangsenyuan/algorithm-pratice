package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		H, n := readTwoNums(reader)
		A := make([][]int, n)
		for i := 0; i < n; i++ {
			A[i] = readNNums(reader, 3)
		}
		res := solve(n, H, A)

		for i := 0; i < n; i++ {
			if i < n-1 {
				fmt.Printf("%d ", res[i])
			} else {
				fmt.Printf("%d\n", res[i])
			}
		}
	}
}

func solve(n int, H int, obstacles [][]int) []int {
	items := make([]Obstacle, n)

	for i := 0; i < n; i++ {
		items[i] = Obstacle{obstacles[i][0], obstacles[i][1], obstacles[i][2], i}
	}

	sort.Sort(Obstacles(items))

	visible := func(ij, up, down Slop) bool {
		if !ij.Less(up) {
			return false
		}
		if !down.Less(ij) {
			return false
		}
		return true
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		up := Slop{0, 1}
		down := Slop{0, -1}

		for j := i + 1; j < n; j++ {
			ii := items[i]
			jj := items[j]
			dx := jj.x - ii.x
			dy := jj.a - ii.a
			ij := Slop{dx, dy}
			if j == i+1 || visible(ij, up, down) {
				res[items[i].i]++
				res[items[j].i]++
			}
			if items[j].tp == 0 && down.Less(ij) {
				down = ij
			} else if items[j].tp == 1 && ij.Less(up) {
				up = ij
			}
		}
	}

	return res
}

type Slop struct {
	dx int
	dy int
}

func (this Slop) Less(that Slop) bool {
	// dy / dx
	// this.dy / this.dx < that.dy / that.dx
	return this.dy*that.dx < this.dx*that.dy
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

type Obstacle struct {
	tp int
	x  int
	a  int
	i  int
}

type Obstacles []Obstacle

func (this Obstacles) Len() int {
	return len(this)
}

func (this Obstacles) Less(i, j int) bool {
	return this[i].x < this[j].x
}

func (this Obstacles) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
