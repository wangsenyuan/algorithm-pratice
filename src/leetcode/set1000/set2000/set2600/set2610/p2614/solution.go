package p2614

import "sort"

func diagonalPrime(nums [][]int) int {
	n := len(nums)

	diag := make([]int, 0, 2*n-1)

	for i := 0; i < n; i++ {
		diag = append(diag, nums[i][i])
		diag = append(diag, nums[i][n-1-i])
	}

	sort.Ints(diag)

	for i := len(diag) - 1; i >= 0; i-- {
		if i+1 < len(diag) && diag[i] == diag[i+1] {
			continue
		}
		num := diag[i]
		if checkPrime(num) {
			return num
		}
	}

	return 0
}

func checkPrime(num int) bool {
	if num == 1 {
		return false
	}

	if num == 2 || num == 3 || num == 5 || num == 7 {
		return true
	}

	if num%2 == 0 {
		return false
	}

	for x := 3; x <= num/3; x += 2 {
		if num%x == 0 {
			return false
		}
	}
	return true
}
