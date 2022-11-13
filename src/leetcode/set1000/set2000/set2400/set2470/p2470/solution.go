package p2470

func subarrayLCM(nums []int, k int) int {
	var res int

	K := int64(k)

	n := len(nums)

	for i := 0; i < n; i++ {
		lcm := int64(nums[i])
		for j := i; j < n; j++ {
			g := gcd(lcm, int64(nums[j]))
			lcm = lcm / g * int64(nums[j])
			if lcm == K {
				res++
			}
			if lcm > K {
				break
			}
		}
	}

	return res
}

func gcd(a, b int64) int64 {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
