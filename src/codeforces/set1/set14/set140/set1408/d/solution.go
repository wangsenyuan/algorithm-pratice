package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, m := readTwoNums(reader)
	robs := make([][]int, n)
	for i := 0; i < n; i++ {
		robs[i] = readNNums(reader, 2)
	}
	lights := make([][]int, m)
	for i := 0; i < m; i++ {
		lights[i] = readNNums(reader, 2)
	}

	res := solve(robs, lights)

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
		for x < len(bs) && (bs[x] < '0' || bs[x] > '9') {
			x++
		}
		x = readInt(bs, x, &res[i])
	}
	return res
}

func solve(robbers [][]int, searchLights [][]int) int {
	R := make([]int, X)
	for i := 0; i < len(robbers); i++ {

		for j := 0; j < len(searchLights); j++ {
			x := searchLights[j][0] - robbers[i][0]
			y := searchLights[j][1] - robbers[i][1] + 1
			// for z <= x, r[z] >
			if x >= 0 {
				R[x] = max(R[x], y)
			}
		}
	}
	ans := inf
	for i := X - 2; i >= 0; i-- {
		R[i] = max(R[i], R[i+1])
		ans = min(ans, i+R[i])
	}

	return ans
}

const X = 1_000_010

const inf = 1 << 30

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
