package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	n := readNum(reader)
	a := readNNums(reader, n)
	res := solve(a)
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

func compact(a []int) []int {
	n := len(a)
	arr := make([]int, n)
	copy(arr, a)
	sort.Ints(arr)
	x := slices.Max(a)
	id := make([]int, x+1)
	var m int
	for i := 1; i <= n; i++ {
		if i == n || arr[i] > arr[i-1] {
			id[arr[i-1]] = m
			m++
		}
	}
	for i := range n {
		a[i] = id[a[i]]
	}
	return a
}

func solve(a []int) int {
	a = compact(a)
	x := slices.Max(a) + 1

	n := len(a)

	check := func(x int, y int) int {
		// 把所有的x移动到y的前面
		var res int
		var cnt int
		for i := 0; i < n; i++ {
			if a[i] == x {
				res += cnt
			} else if a[i] == y {
				cnt++
			}
		}
		return res
	}

	cnt := make([][]int, x)
	for i := 0; i < x; i++ {
		cnt[i] = make([]int, x)
		for j := 0; j < x; j++ {
			if i != j {
				cnt[i][j] = check(i, j)
			}
		}
	}
	T := 1 << x
	dp := make([]int, T)
	for i := 0; i < T; i++ {
		dp[i] = 1 << 60
	}
	for i := 0; i < x; i++ {
		dp[1<<i] = 0
	}

	for state := 0; state < T; state++ {
		var arr []int
		for i := 0; i < x; i++ {
			if (state>>i)&1 == 1 {
				arr = append(arr, i)
			}
		}
		for i := 0; i < x; i++ {
			if (state>>i)&1 == 0 {
				var sum int
				for _, j := range arr {
					sum += cnt[j][i]
				}
				dp[state|(1<<i)] = min(dp[state|(1<<i)], dp[state]+sum)
			}
		}
	}

	return dp[T-1]
}
