package main

import (
	"bufio"
	"fmt"
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
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + int(bytes[i]-'0')
		i++
	}
	*val = tmp * sign
	return i
}

func readNum(scanner *bufio.Scanner) (a int) {
	scanner.Scan()
	readInt(scanner.Bytes(), 0, &a)
	return
}

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	res := readNNums(scanner, 2)
	a, b = res[0], res[1]
	return
}

func readNNums(scanner *bufio.Scanner, n int) []int {
	res := make([]int, n)
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
	return res
}

func fillNNums(scanner *bufio.Scanner, n int, res []int) {
	x := 0
	scanner.Scan()
	for i := 0; i < n; i++ {
		for x < len(scanner.Bytes()) && scanner.Bytes()[x] == ' ' {
			x++
		}
		x = readInt(scanner.Bytes(), x, &res[i])
	}
}

func readUint64(bytes []byte, from int, val *uint64) int {
	i := from

	var tmp uint64
	for i < len(bytes) && bytes[i] != ' ' {
		tmp = tmp*10 + uint64(bytes[i]-'0')
		i++
	}
	*val = tmp

	return i
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	tc := readNum(scanner)

	for tc > 0 {
		tc--
		n, Q := readTwoNums(scanner)
		A := readNNums(scanner, n)
		K := make([]int, Q)
		for i := 0; i < Q; i++ {
			K[i] = readNum(scanner)
		}
		R := solve(n, Q, A, K)
		for _, r := range R {
			fmt.Println(r)
		}
	}
}

func solve(n, Q int, A []int, K []int) []int {
	dp := make([]int, 32)
	fp := make([]int, 32)

	var dfs func(i int, nums []int)

	dfs = func(i int, nums []int) {
		if i < 0 || len(nums) == 0 {
			return
		}
		ones := make([]int, 0, len(nums))
		zeros := make([]int, 0, len(nums))

		for j := 0; j < len(nums); j++ {
			num := nums[j]
			x := (num >> uint(i)) & 1
			if x == 1 {
				ones = append(ones, num)
				dp[i] += len(zeros)
			} else {
				zeros = append(zeros, num)
				fp[i] += len(ones)
			}
		}
		dfs(i-1, ones)
		dfs(i-1, zeros)
	}

	dfs(31, A)

	R := make([]int, Q)

	for q := 0; q < Q; q++ {
		k := K[q]
		var r int

		for i := 31; i >= 0; i-- {
			x := (k >> uint(i)) & 1
			if x == 0 {
				// will not change the result
				r += fp[i]
			} else {
				// will change
				r += dp[i]
			}
		}

		R[q] = r
	}

	return R
}
