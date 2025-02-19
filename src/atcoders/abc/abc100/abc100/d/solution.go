package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	n, m := readTwoNums(reader)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = readNNums(reader, 3)
	}
	return solve(m, a)
}

const inf = 1 << 60

type pair struct {
	first  int
	second int
}

func solve(m int, a [][]int) int {
	n := len(a)

	arr := make([]pair, n)
	get := func(state int) int {
		for i := 0; i < n; i++ {
			arr[i].first = 0
			arr[i].second = i
			for j := 0; j < 3; j++ {
				if (state>>j)&1 == 0 {
					arr[i].first -= a[i][j]
				} else {
					arr[i].first += a[i][j]
				}
			}
		}

		sort.Slice(arr, func(i, j int) bool {
			return arr[i].first > arr[j].first
		})

		sum := []int{0, 0, 0}
		for i := 0; i < m; i++ {
			j := arr[i].second
			sum[0] += a[j][0]
			sum[1] += a[j][1]
			sum[2] += a[j][2]
		}
		return abs(sum[0]) + abs(sum[1]) + abs(sum[2])
	}

	var ans int

	for state := 0; state < 8; state++ {
		ans = max(ans, get(state))
	}
	return ans
}

func abs(num int) int {
	return max(num, -num)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
