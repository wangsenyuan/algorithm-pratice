package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	lines := make([][]int, 4)

	for i := 0; i < 4; i++ {
		lines[i] = readNNums(reader, 4)
	}

	res := solve(lines)

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

func solve(lines [][]int) bool {
	var x [][]int
	var y [][]int

	for _, line := range lines {
		if line[0] == line[2] {
			if line[1] > line[3] {
				line[1], line[3] = line[3], line[1]
			}
			x = append(x, line)
		} else if line[1] == line[3] {
			if line[0] > line[2] {
				line[0], line[2] = line[2], line[0]
			}
			y = append(y, line)
		} else {
			return false
		}
	}

	if len(x) != len(y) {
		return false
	}

	if x[0][0] > x[1][0] {
		x[0], x[1] = x[1], x[0]
	}

	if y[0][1] > y[1][1] {
		y[0], y[1] = y[1], y[0]
	}

	if x[0][1] != x[1][1] || x[0][3] != x[1][3] || x[0][0] == x[1][0] {
		return false
	}

	if y[0][0] != y[1][0] || y[0][2] != y[1][2] || y[0][1] == y[1][1] {
		return false
	}

	if x[0][0] != y[0][0] || x[0][1] != y[0][1] {
		return false
	}

	if x[0][2] != y[1][0] || x[0][3] != y[1][1] {
		return false
	}

	if x[1][0] != y[0][2] || x[1][1] != y[0][3] {
		return false
	}

	if x[1][2] != y[1][2] || x[1][3] != y[1][3] {
		return false
	}

	return true
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
