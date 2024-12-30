package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var buf bytes.Buffer
	tc := readNum(reader)
	for tc > 0 {
		tc--
		res, _ := process(reader)
		if len(res) == 0 {
			buf.WriteString("0\n")
		} else {
			buf.WriteString(fmt.Sprintf("1 %d %d\n", res[0], res[1]))
		}
	}
	fmt.Print(buf.String())
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

func process(reader *bufio.Reader) (res []int, robots [][]int) {
	n := readNum(reader)
	robots = make([][]int, n)
	for i := range n {
		robots[i] = readNNums(reader, 6)
	}
	res = solve(robots)
	return
}

const inf = 1 << 60

type robot struct {
	x int
	y int
	f []int
}

func solve(robots [][]int) []int {
	n := len(robots)
	rs := make([]robot, n)
	for i, cur := range robots {
		rs[i] = robot{cur[0], cur[1], cur[2:]}
	}

	slices.SortFunc(rs, func(a, b robot) int {
		return a.y - b.y
	})
	// wy 表示是否所有的在y下面的都可以向上运
	wy := rs[0].y
	for i := 0; i < n; i++ {
		wy = rs[i].y
		if rs[i].f[1] == 0 {
			break
		}
	}
	vy := rs[n-1].y
	for i := n - 1; i >= 0; i-- {
		vy = rs[i].y
		if rs[i].f[3] == 0 {
			break
		}
	}

	if vy > wy {
		return nil
	}

	slices.SortFunc(rs, func(a, b robot) int {
		return a.x - b.x
	})
	wx := rs[0].x
	for i := 0; i < n; i++ {
		wx = rs[i].x
		if rs[i].f[2] == 0 {
			break
		}
	}

	vx := rs[n-1].x
	for i := n - 1; i >= 0; i-- {
		vx = rs[i].x
		if rs[i].f[0] == 0 {
			break
		}
	}

	if vx > wx {
		return nil
	}
	return []int{wx, wy}
}
