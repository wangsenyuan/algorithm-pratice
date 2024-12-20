package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	res, _ := process(reader)
	if len(res) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	s := fmt.Sprintf("%v", res)
	fmt.Println(s[1 : len(s)-1])
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

func process(reader *bufio.Reader) (res []int, facts [][]int) {
	// 格式为 n m
	// 1 1 3
	n, m := readTwoNums(reader)
	facts = make([][]int, m)
	for i := range m {
		facts[i] = readNNums(reader, 3)
	}
	res = solve(n, facts)
	return
}

const inf = 1_000_000_000

func solve(n int, facts [][]int) []int {
	// sorted in non-decreasing order
	var x [][]int
	// unsorted in non-decreasing order
	var y [][]int

	for _, fact := range facts {
		l, r := fact[1]-1, fact[2]-1
		if fact[0] == 1 {
			x = append(x, []int{l, r})
		} else {
			y = append(y, []int{l, r})
		}
	}

	slices.SortFunc(x, func(a, b []int) int {
		return a[0] - b[0]
	})

	var arr [][]int

	prev := -1

	// arr代表的区间必须是sorted in non-decreasing order的
	for _, cur := range x {
		if cur[0] > prev {
			arr = append(arr, cur)
			prev = cur[1]
		} else {
			// cur[0] <= prev
			prev = max(prev, cur[1])
			arr[len(arr)-1][1] = prev
		}
	}

	res := make([]int, n)

	for i, j, val := 0, 0, inf; i < n; i++ {
		if j < len(arr) && i > arr[j][1] {
			j++
			val--
		} else if j >= len(arr) || i <= arr[j][0] {
			val--
		}
		res[i] = val
	}

	// dp[i]表示前面那个比i大的数的下标
	stack := make([]int, n)
	var top int
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		for top > 0 && res[stack[top-1]] <= res[i] {
			top--
		}
		if top == 0 {
			dp[i] = -1
		} else {
			dp[i] = stack[top-1]
		}
		stack[top] = i
		top++
	}
	// 现在检查y是否满足
	for _, cur := range y {
		l, r := cur[0], cur[1]
		if dp[r] < l {
			// 这个区间全部是相同的
			return nil
		}
	}

	return res
}
