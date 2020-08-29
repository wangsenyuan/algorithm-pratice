package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
		n := readNum(reader)
		A := readNNums(reader, n)
		fmt.Println(solve(n, A))
	}
}

func solve(n int, A []int) int64 {
	if n == 1 {
		return int64(A[0])
	}
	// i is for postive & j is for negative

	var res int64
	var i int
	if A[0] < 0 {
		var x int = math.MinInt32
		var j int
		for j < n && A[j] < 0 {
			x = max(x, A[j])
			j++
		}
		res = int64(x)
		i = j
	}
	var x = 0
	for i <= n {
		// i is at postive position
		if i < n && A[i] > 0 {
			x = max(x, A[i])
			i++
			continue
		}
		res += int64(x)
		x = 0
		if i == n {
			break
		}
		// A[i] < 0
		var y int = math.MinInt32
		for i < n && A[i] < 0 {
			y = max(y, A[i])
			i++
		}
		res += int64(y)
	}

	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
