package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	if res {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
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

func process(reader *bufio.Reader) bool {
	n := readNum(reader)
	points := make([][]int, n)
	for i := range n {
		points[i] = readNNums(reader, 2)
	}
	return solve(points)
}

func solve(points [][]int) bool {
	n := len(points)
	if n&1 == 1 {
		return false
	}

	verts := make([]pair, n)
	for i := 0; i < n; i++ {
		verts[i] = pair{points[i][0], points[i][1]}
	}

	center := make(map[pair]int)
	h := n / 2
	for i := 0; i < n; i++ {
		a := verts[i]
		b := verts[(i+h)%n]
		c := b.add(a)
		center[c]++
		if len(center) > 1 {
			return false
		}
	}

	return true
}

type pair struct {
	first  int
	second int
}

func (this pair) sub(that pair) pair {
	return pair{this.first - that.first, this.second - that.second}
}

func (this pair) add(that pair) pair {
	return pair{this.first + that.first, this.second + that.second}
}

func cw(a, b, c pair) bool {
	return (b.first-a.first)*(c.second-a.second)-(b.second-a.second)*(c.first-a.first) < 0
}
