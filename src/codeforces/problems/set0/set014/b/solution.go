package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n, x0 := readTwoNums(reader)
	men := make([][]int, n)

	for i := 0; i < n; i++ {
		men[i] = readNNums(reader, 2)
	}
	res := solve(x0, men)

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

func solve(x0 int, men [][]int) int {

	for i := 0; i < len(men); i++ {
		if men[i][0] > men[i][1] {
			men[i][0], men[i][1] = men[i][1], men[i][0]
		}
	}

	x1, x2 := -1, -1

	for _, man := range men {
		if x1 < 0 || man[1] < x1 {
			x1 = man[1]
		}
		if x2 < 0 || man[0] > x2 {
			x2 = man[0]
		}
	}
	// x1 是最小的右边界
	// x2 是最大的左边界
	if x1 < x2 {
		return -1
	}

	// only in [x2...x1], Bob can take all photos

	for _, man := range men {
		if man[1] < x2 {
			return -1
		}
		if man[0] > x1 {
			return -1
		}
	}

	if x0 >= x2 && x0 <= x1 {
		return 0
	}

	d1 := abs(x0 - x1)
	d2 := abs(x0 - x2)
	return min(d1, d2)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
