package p888_

func fairCandySwap(A []int, B []int) []int {
	sa := sumArray(A)
	sb := sumArray(B)
	diff := sb - sa
	half := diff / 2
	ii := make(map[int]int)
	for i := 0; i < len(A); i++ {
		ii[A[i]]++
	}

	for i := 0; i < len(B); i++ {
		b := B[i]
		// sa - a + b = sb - b + a => sb - sa = 2 * (b - a) => b - a == half
		a := b - half
		if ii[a] > 0 {
			return []int{a, b}
		}
	}

	return nil
}

func sumArray(num []int) int {
	var res int
	for i := 0; i < len(num); i++ {
		res += num[i]
	}
	return res
}
