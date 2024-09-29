package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	n := readNum(reader)
	points := make([][]int, n)
	for i := 0; i < n; i++ {
		points[i] = readNNums(reader, 2)
	}
	res := solve(points)

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

type pair struct {
	first  int
	second int
}

func solve(points [][]int) int {

	slices.SortFunc(points, func(a, b []int) int {
		return (a[0] + a[1]) - (b[0] + b[1])
	})

	var last int
	ans := 1

	for i := 1; i < len(points); i++ {
		if points[i][0]-points[i][1] >= points[last][0]+points[last][1] {
			last = i
			ans++
		}
	}
	return ans
}
func solve1(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		return a[0] - b[0]
	})

	n := len(points)
	stack := make([]pair, n)
	var top int

	for _, cur := range points {
		// x[i] -  w[i]
		x := cur[0] - cur[1]
		j := search(top, func(j int) bool {
			return stack[j].first > x
		})
		j--
		var v int
		if j < 0 {
			v = 1
		} else {
			v = stack[j].second + 1
		}
		if top > 0 {
			v = max(v, stack[top-1].second)
		}
		for top > 0 && stack[top-1].first >= cur[0]+cur[1] {
			top--
		}
		stack[top] = pair{cur[0] + cur[1], v}
		top++
	}
	return stack[top-1].second
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
