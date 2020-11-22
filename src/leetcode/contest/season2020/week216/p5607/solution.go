package p5607

func waysToMakeFair(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}
	odd := make([]int, n)
	even := make([]int, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			odd[i] = odd[i-1]
			even[i] = even[i-1]
		}
		if i%2 == 0 {
			even[i] += nums[i]
		} else {
			odd[i] += nums[i]
		}
	}

	var ans int

	for i := 0; i < n; i++ {
		// if we remove i
		var a, b int
		if i > 0 {
			a = odd[i-1]
			b = even[i-1]
		}
		c := odd[n-1] - odd[i]
		d := even[n-1] - even[i]

		if a+d == b+c {
			ans++
		}
	}
	return ans
}
