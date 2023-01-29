package p2550

func countQuadruplets(nums []int) int64 {
	// n <= 4000
	n := len(nums)
	// i < j < k < l
	// A[i] < A[k] < A[j] < A[l]

	set := func(arr []int, p int, v int) {
		p += n
		arr[p] += v
		for p > 1 {
			arr[p>>1] = arr[p] + arr[p^1]
			p >>= 1
		}
	}

	get := func(arr []int, l int, r int) int {
		var res int
		l += n
		r += n
		for l < r {
			if l&1 == 1 {
				res += arr[l]
				l++
			}
			if r&1 == 1 {
				r--
				res += arr[r]
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	arr := make([]int, 2*n)
	arr2 := make([]int, 2*n)
	var res int64
	for j := 0; j+1 < n; j++ {
		for i := 0; i < 2*n; i++ {
			arr2[i] = 0
		}
		for k := n - 1; k > j; k-- {
			if nums[j] > nums[k] {
				a := get(arr, 0, nums[k])
				b := get(arr2, nums[j], n)
				res += int64(a) * int64(b)
			}
			set(arr2, nums[k]-1, 1)
		}
		set(arr, nums[j]-1, 1)
	}

	return res
}
