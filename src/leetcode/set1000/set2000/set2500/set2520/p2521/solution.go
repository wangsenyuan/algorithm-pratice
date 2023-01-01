package p2521

func distinctPrimeFactors(nums []int) int {

	factors := make(map[int]bool)

	for _, num := range nums {
		if num%2 == 0 {
			factors[2] = true

			for num%2 == 0 {
				num /= 2
			}
		}

		for x := 3; x <= num/x; x += 2 {
			if num%x == 0 {
				factors[x] = true
				for num%x == 0 {
					num /= x
				}
			}
		}
		if num > 1 {
			factors[num] = true
		}
	}

	return len(factors)
}
