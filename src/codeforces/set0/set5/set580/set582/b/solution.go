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

	n, k := readTwoNums(reader)
	a := readNNums(reader, n)

	res := solve(a, k)

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

const MAX_A = 300

func solve(a []int, k int) int {
	n := len(a)
	var stack []int
	var d int
	for i := 0; i < n && k > 0; i++ {
		m := len(stack)
		for _, x := range a {
			j := sort.Search(len(stack), func(j int) bool {
				return stack[j] > x
			})
			if j == len(stack) {
				stack = append(stack, x)
			} else {
				stack[j] = x
			}
		}
		d = len(stack) - m
		k--
	}

	return len(stack) + d*k
}

func solve1(a []int, k int) int {
	n := len(a)

	if k < 2*n {
		res := findLis(repeate(a, k))
		return slices.Max(res)
	}

	a2 := repeate(a, n)
	// k > n
	l1 := findLis(a2)
	l2 := findLes(a2)
	ma := slices.Max(a)
	// 找到最长的序列的连续序列
	freq := make([]int, ma+1)

	var best int

	for _, num := range a {
		freq[num]++
	}

	for i := 1; i <= ma; i++ {
		best = max(best, l1[i]+l2[i]+freq[i]*(k-2*n))
	}

	return best
}

func repeate(arr []int, k int) []int {
	n := len(arr)
	res := make([]int, n*k)
	for i := 0; i < k; i++ {
		copy(res[i*n:], arr)
	}
	return res
}

func findLes(a []int) []int {
	n := len(a)
	arr := make([]int, n)
	ma := slices.Max(a)
	copy(arr, a)
	res := make([]int, ma+1)
	var p int

	for i := n - 1; i >= 0; i-- {
		j := sort.Search(p, func(j int) bool {
			return arr[j] < a[i]
		})
		arr[j] = a[i]
		if j == p {
			p++
		}
		res[a[i]] = max(res[a[i]], j+1)
	}

	for i := ma - 1; i >= 0; i-- {
		res[i] = max(res[i], res[i+1])
	}
	return res
}

func findLis(a []int) []int {
	n := len(a)
	arr := make([]int, n)
	ma := slices.Max(a)
	copy(arr, a)
	// res[i] = 表示最后一个元素不超过i时的最大序列
	res := make([]int, ma+1)
	var p int

	for i := 0; i < n; i++ {
		j := sort.Search(p, func(j int) bool {
			return arr[j] > arr[i]
		})
		arr[j] = arr[i]
		if j == p {
			p++
		}
		// res[i]表示i前面的最长非递减序列的长度， 包括i
		res[arr[i]] = max(res[arr[i]], j+1)
	}

	for i := 1; i <= ma; i++ {
		res[i] = max(res[i], res[i-1])
	}

	return res
}
