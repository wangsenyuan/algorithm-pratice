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

func readTwoNums(scanner *bufio.Scanner) (a int, b int) {
	scanner.Scan()
	x := readInt(scanner.Bytes(), 0, &a)
	readInt(scanner.Bytes(), x+1, &b)
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
	n, k := readTwoNums(scanner)
	nums := readNNums(scanner, n)
	fmt.Println(solve(n, k, nums))
}

const MOD = 1000000007

func solve(n int, k int, nums []int) int {
	if k >= n {
		return pow(2, n-1)
	}

	if k == 0 {
		var hasZero bool
		for i := 0; i < n; i++ {
			if nums[i] == 0 {
				hasZero = true
				break
			}
		}
		if hasZero {
			return 0
		}
		return pow(2, n-1)
	}

	left := make([]int, n)
	cnt := make(map[int]int)

	for i, j := 0, 0; i < n; i++ {
		if nums[i] <= k {
			cnt[nums[i]]++
		}

		if len(cnt) == k+1 {
			for j < i && len(cnt) == k+1 {
				if nums[j] <= k {
					cnt[nums[j]]--
					if cnt[nums[j]] == 0 {
						delete(cnt, nums[j])
					}
				}
				j++
			}
		}
		left[i] = j
	}

	ans, sum := make([]int, n), make([]int, n)
	ans[0] = 1
	sum[0] = 1
	for i := 1; i < n; i++ {
		l := left[i]
		ans[i] = sum[i-1]
		if l > 1 {
			ans[i] = (ans[i] - sum[l-2] + 3*MOD) % MOD
		} else if l == 0 {
			ans[i] = (ans[i] + 1) % MOD
		}
		sum[i] = (sum[i-1] + ans[i]) % MOD
	}
	return ans[n-1]
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	c := pow(a, b/2)
	d := (c * c) % MOD
	if b%2 == 1 {
		return (a * d) % MOD
	}
	return d
}
