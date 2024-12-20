package p3385

func beautifulSplits(nums []int) int {
	n := len(nums)

	z := make([]int, n)

	zf := func(arr []int) []int {
		n := len(arr)
		clear(z[:n])
		for i, l, r := 1, 0, 0; i < n; i++ {
			if i <= r {
				z[i] = min(r-i+1, z[i-l])
			}
			for i+z[i] < n && arr[z[i]] == arr[i+z[i]] {
				z[i]++
			}
			if i+z[i]-1 > r {
				l = i
				r = i + z[i] - 1
			}
		}
		return z[:n]
	}

	a := zf(nums)
	z = make([]int, n)

	var res int

	for i := 1; i+1 < n; i++ {
		b := zf(nums[i:])

		for j := i + 1; j < n; j++ {
			if a[i] >= i && i <= j-i || b[j-i] >= j-i && j-i <= n-j {
				res++
			}
		}
	}
	return res
}

func beautifulSplits1(nums []int) int {
	n := len(nums)
	// 1/2/3段，1是2的前缀，或者2是3的前缀
	// 且1/2/3都为空
	lcs := make([][]int, n+1)

	for i := n; i >= 0; i-- {
		lcs[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] == nums[j] {
				lcs[i][j] = lcs[i+1][j+1] + 1
			}
		}
	}
	var res int
	// 这里有重复计算的部分
	for i := 1; i < n; i++ {
		for j := i; j+1 < n; j++ {
			// a = [0...i]
			// b = [i...j]
			// c = [j+1....n]
			a := lcs[0][i]
			b := lcs[i][j+1]
			if a >= i && i <= j-i+1 || b >= j-i+1 && j-i+1 <= n-j-1 {
				res++
			}
		}
	}

	return res
}
