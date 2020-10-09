package smallestposiblesum

func Solution(arr []int) int {
	// return smallest possible sum of all numbers in array
	// if X[i] > X[j] then X[i] = X[i] - X[j]
	// when no more operations can perform? when all nums equal
	// gcd of the arr
	if len(arr) == 0 {
		return 0
	}
	var g = arr[0]

	for i := 1; i < len(arr); i++ {
		g = gcd(g, arr[i])
	}

	return g * len(arr)
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
