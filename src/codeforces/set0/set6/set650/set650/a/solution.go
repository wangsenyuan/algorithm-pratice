package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	pts := make([][]int, n)

	for i := 0; i < n; i++ {
		pts[i] = readNNums(reader, 2)
	}

	res := solve(pts)

	fmt.Println(res)
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

func solve(pts [][]int) int {
	x := make(map[int]int)
	y := make(map[int]int)
	p := make(map[Pair]int)

	for _, pt := range pts {
		x[pt[0]]++
		y[pt[1]]++
		p[Pair{pt[0], pt[1]}]++
	}

	var res int
	for _, v := range x {
		res += v * (v - 1) / 2
	}
	for _, v := range y {
		res += v * (v - 1) / 2
	}

	for _, v := range p {
		res -= v * (v - 1) / 2
	}

	return res
}

type Pair struct {
	first  int
	second int
}
