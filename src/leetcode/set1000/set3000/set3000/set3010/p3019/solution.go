package p3019

const D = 30

func minOrAfterOperations(nums []int, k int) int {
	//n := len(nums)

	var mask int

	var res int

	for d := D - 1; d >= 0; d-- {
		mask |= 1 << d
		var cnt int
		and := -1

		for _, num := range nums {
			and &= num & mask
			if and != 0 {
				cnt++
			} else {
				and = -1
			}
		}
		if cnt > k {
			res |= 1 << d
			mask ^= 1 << d
		}
	}

	return res
}
