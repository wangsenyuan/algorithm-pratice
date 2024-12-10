package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	res := process(reader)

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

func process(reader *bufio.Reader) int {
	n, r := readTwoNums(reader)
	tasks := make([][]int, n)
	for i := range n {
		tasks[i] = readNNums(reader, 2)
	}
	return solve(r, tasks)
}

const inf = 1 << 60

func solve(r int, tasks [][]int) int {
	n := len(tasks)

	var arr []pair

	for i := 0; i < n; i++ {
		if tasks[i][1] >= 0 {
			arr = append(arr, pair{tasks[i][0], tasks[i][1]})
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return a.first - b.first
	})

	var pos int

	for _, cur := range arr {
		if cur.first > r {
			break
		}
		r += cur.second
		pos++
	}

	arr = arr[:0]
	for i := 0; i < n; i++ {
		if tasks[i][1] < 0 {
			arr = append(arr, pair{tasks[i][0], tasks[i][1]})
		}
	}

	slices.SortFunc(arr, func(a, b pair) int {
		return (b.first + b.second) - (a.first + a.second)
	})

	dp := make([]int, r+1)
	fp := make([]int, r+1)
	for i := 0; i <= r; i++ {
		dp[i] = -n
		fp[i] = -n
	}
	dp[r] = 0

	for _, cur := range arr {
		// x + b[i] >= 0
		// and x >= a[i]
		for x := cur.first; x <= r; x++ {
			if x+cur.second >= 0 {
				fp[x+cur.second] = max(fp[x+cur.second], dp[x]+1)
			}
		}
		for x := 0; x <= r; x++ {
			dp[x] = max(dp[x], fp[x])
			fp[x] = -n
		}
	}

	return pos + slices.Max(dp)
}

type pair struct {
	first  int
	second int
}
