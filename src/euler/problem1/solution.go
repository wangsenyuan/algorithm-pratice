package problem1

func MultipleOf3Or5(n int) int {
	var ans int

	for i := 3; i < n; i += 3 {
		ans += i
	}

	for i := 5; i < n; i += 5 {
		ans += i
	}

	for i := 15; i < n; i += 15 {
		ans -= i
	}

	return ans
}
