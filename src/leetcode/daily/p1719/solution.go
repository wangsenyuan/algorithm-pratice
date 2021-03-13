package p1719

import "math"

func missingTwo(nums []int) []int {
	var sum int
	var sum2 int
	for _, num := range nums {
		sum += num
		sum2 += num * num
	}
	n := len(nums) + 2
	var tot, tot2 int

	for i := 1; i <= n; i++ {
		tot += i
		tot2 += i * i
	}
	tot -= sum
	tot2 -= sum2
	// a + b = tot - sum =
	// a ^^ 2 + b ^^ 2 = tot2 - sum2
	// ( a + b) ^^ 2 = tot ^^ 2 = a2 + 2ab + b2 = 2ab + tot2
	diff := tot*tot - tot2
	diff /= 2
	// a + b = tot
	// a * b = diff
	// (tot - a) * a = diff
	// a*a - tot * a + diff = 0
	// tot + sqrt(tot * tot + 4 * diff)
	x := int(math.Sqrt(float64(tot)*float64(tot) - 4*float64(diff)))
	a := (tot - x) / 2
	b := tot - a
	return []int{a, b}
}

func missingTwo1(nums []int) []int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	n := len(nums) + 2
	tot := (1 + n) * n / 2
	// 1 + 2 + .. + n = (1 + n) * n / 2
	// a + b = tot - sum
	ab := tot - sum
	var n1 int
	var n2 int
	for i := 0; i < len(nums); i++ {
		x := abs(nums[i])
		if x == n-1 {
			n1 = -1
		} else if x == n {
			n2 = -1
		} else {
			nums[x-1] = -abs(nums[x-1])
		}
	}
	if n2 == 0 {
		return []int{ab - n, n}
	}
	if n1 == 0 {
		return []int{ab - n + 1, n - 1}
	}
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x > 0 {
			return []int{ab - (i + 1), i + 1}
		}
	}
	return nil
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
