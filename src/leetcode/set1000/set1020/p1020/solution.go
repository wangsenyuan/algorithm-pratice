package p1020

func canThreePartsEqualSum(A []int) bool {
	var sum int

	for _, a := range A {
		sum += a
	}

	if sum%3 != 0 {
		return false
	}
	sum /= 3

	var x int
	for i := 0; i < len(A); i++ {
		x += A[i]
		if x == sum {
			var y int
			for j := i + 1; j < len(A)-1; j++ {
				y += A[j]
				if y == sum {
					return true
				}
			}
		}
	}
	return false
}
