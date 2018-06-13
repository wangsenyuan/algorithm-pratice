package main

import (
	"bufio"
	"fmt"
	"os"
)

func readInt(bytes []byte, from int, val *int) int {
	i := from
	tmp := 0
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := -1
	scanner.Scan()
	for i := 0; i < n; i++ {
		x = readInt(scanner.Bytes(), x+1, &res[i])
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := readNum(scanner)
	nums := readNNums(scanner, n)
	fmt.Println(solve(n, nums))
}

// MOD
const MOD = 1000000007

func solve(n int, nums []int) int {
	stack := make([]int, n)
	calMaxSum := func() []int64 {
		res := make([]int64, n)
		p := 0
		for i := 0; i < n; i++ {
			for p > 0 && nums[stack[p-1]] <= nums[i] {
				p--
			}
			j := -1
			var tmp int64
			if p > 0 {
				j = stack[p-1]
				tmp = res[j]
			}
			res[i] = (tmp + int64(i-j)*int64(nums[i])) % MOD
			stack[p] = i
			p++
		}
		return res
	}

	calMinSum := func() []int64 {
		res := make([]int64, n)
		p := 0
		for i := n - 1; i >= 0; i-- {
			for p > 0 && nums[stack[p-1]] >= nums[i] {
				p--
			}
			j := n
			var tmp int64
			if p > 0 {
				j := stack[p-1]
				tmp = res[j]
			}
			res[i] = (tmp + int64(j-i)*int64(nums[i])) % MOD
			stack[p] = i
			p++
		}
		return res
	}

	calMinSumSuffix := func(minSum []int64) []int64 {
		res := make([]int64, n)
		for i := n - 1; i >= 0; i-- {
			if i+1 < n {
				res[i] = res[i+1]
			}
			res[i] = (res[i] + minSum[i]) % MOD
		}
		return res
	}

	maxSum := calMaxSum()
	minSum := calMinSum()
	minSumSuf := calMinSumSuffix(minSum)

	var ans int64
	for i := 0; i < n-1; i++ {
		ans = (ans + maxSum[i]*minSumSuf[i+1]) % MOD
	}
	return int(ans)
}
