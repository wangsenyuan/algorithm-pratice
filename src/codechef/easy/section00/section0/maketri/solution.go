package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Int64Slice []int64

func (s Int64Slice) Len() int {
	return len(s)
}

func (s Int64Slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Int64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n, l, r := firstLine(scanner)

	nums := strands(scanner, n)

	res := solve(n, l, r, nums)

	fmt.Println(res)
}

func solve(n int, left int64, right int64, nums []int64) int64 {
	sort.Sort(Int64Slice(nums))

	stack := make([][]int64, n)
	p := 0

	for i := 0; i < n-1; i++ {
		a, b := nums[i], nums[i+1]
		x := b - a + 1
		y := b + a - 1

		for p > 0 && stack[p-1][1] >= x {
			x = stack[p-1][0]
			if stack[p-1][1] > y {
				y = stack[p-1][1]
			}
			p--
		}

		stack[p] = []int64{x, y}
		p++
	}

	var ans int64
	for i := 0; i < p; i++ {
		x, y := stack[i][0], stack[i][1]
		if x < left {
			x = left
		}
		if y > right {
			y = right
		}
		if y >= x {
			ans += y - x + 1
		}
	}
	return ans
}

func firstLine(scanner *bufio.Scanner) (n int, l int64, r int64) {
	scanner.Scan()
	bytes := scanner.Bytes()

	i := 0
	for bytes[i] != ' ' {
		n = n*10 + int(bytes[i]-'0')
		i++
	}
	i++
	for bytes[i] != ' ' {
		l = l*10 + int64(bytes[i]-'0')
		i++
	}
	i++
	for i < len(bytes) {
		r = r*10 + int64(bytes[i]-'0')
		i++
	}
	return
}

func strands(scanner *bufio.Scanner, n int) []int64 {
	nums := make([]int64, n)

	scanner.Scan()
	bytes := scanner.Bytes()

	i, j := 0, 0
	for i < len(bytes) {
		var tmp int64
		for i < len(bytes) && bytes[i] != ' ' {
			tmp = tmp*10 + int64(bytes[i]-'0')
			i++
		}
		nums[j] = tmp
		j++

		i++
	}

	return nums
}
