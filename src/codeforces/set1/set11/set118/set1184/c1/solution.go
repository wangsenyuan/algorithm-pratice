package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res := process(reader)
	fmt.Println(res[0], res[1])
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

func process(reader *bufio.Reader) []int {
	n := readNum(reader)
	points := make([][]int, 4*n+1)
	for i := range 4*n + 1 {
		points[i] = readNNums(reader, 2)
	}
	return solve(points)
}

const inf = 1e12

func solve(points [][]int) []int {
	// 这个地方不对， 有可能最远的那个正好在外面
	rect := []int{inf, inf, -inf, -inf}

	for _, cur := range points {
		x, y := cur[0], cur[1]
		rect[0] = min(rect[0], x)
		rect[1] = min(rect[1], y)
		rect[2] = max(rect[2], x)
		rect[3] = max(rect[3], y)
	}

	borders := make([][]int, 4)
	var sigle [][]int

	for j, cur := range points {
		x, y := cur[0], cur[1]
		var add bool
		for i := 0; i < 3; i += 2 {
			if x == rect[i] {
				borders[i] = append(borders[i], j)
				add = true
			}
			if y == rect[i+1] {
				borders[i+1] = append(borders[i+1], j)
				add = true
			}
		}
		if !add {
			sigle = append(sigle, cur)
		}
	}

	for i := 0; i < 4; i++ {
		if len(borders[i]) == 1 {
			return points[borders[i][0]]
		}
	}

	if len(sigle) == 1 {
		return sigle[0]
	}

	return nil
}
