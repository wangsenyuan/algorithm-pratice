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
	n, k := readTwoNums(reader)

	G := make([][]int, n)

	for i := 0; i < n; i++ {
		G[i] = readNNums(reader, n)
	}

	res := solve(n, k, G)

	for i := 0; i < len(res); i++ {
		for j := 0; j < len(res[i]); j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}

const INF = math.MaxInt32

func solve(n, k int, G [][]int) [][]int {
	X := make([][]int, n)
	deque := make([]int, n)
	for i := 0; i < n; i++ {
		X[i] = make([]int, n-k+1)
		findSlidWindownMin(G[i], n, k, X[i], deque)
	}

	res := make([][]int, n-k+1)
	for i := 0; i < n-k+1; i++ {
		res[i] = make([]int, n-k+1)
	}

	arr := make([]int, n)
	back := make([]int, n-k+1)
	for c := 0; c < n-k+1; c++ {
		for i := 0; i < n; i++ {
			arr[i] = X[i][c]
		}
		findSlidWindownMin(arr, n, k, back, deque)
		for i := 0; i < n-k+1; i++ {
			res[i][c] = back[i]
		}
	}

	return res
}

func findSlidWindownMin(arr []int, n int, k int, res []int, deque []int) {
	var p int

	for i := 0; i < k; i++ {
		for p > 0 && deque[p-1] > arr[i] {
			p--
		}
		deque[p] = arr[i]
		p++
	}
	res[0] = deque[0]
	var f int
	for i := k; i < n; i++ {
		for p > f && deque[p-1] > arr[i] {
			p--
		}
		deque[p] = arr[i]
		p++

		if arr[i-k] == deque[f] {
			f++
		}
		res[i-k+1] = deque[f]
	}
}
