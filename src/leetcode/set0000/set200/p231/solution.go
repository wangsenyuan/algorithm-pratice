package main

func isPowerOfTwo(n int) bool {
	ones := 0
	for n > 0 {
		if n&1 == 1 {
			ones++
		}
		if ones > 1 {
			return false
		}
		n >>= 1
	}

	return ones == 1
}
