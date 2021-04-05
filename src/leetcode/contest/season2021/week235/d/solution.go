package d

const MAX_X = 200010

func countDifferentSubsequenceGCDs(nums []int) int {
	n := len(nums)
	can := make([]bool, MAX_X)
	for i := 0; i < n; i++ {
		can[nums[i]] = true
	}
	var res int
	for x := 1; x < MAX_X; x++ {
		first := -1
		for y := x; y < MAX_X; y += x {
			if can[y] {
				if first < 0 {
					first = y
				} else {
					first = gcd(first, y)
				}
			}
		}
		if first == x {
			res++
		}
	}
	return res
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
