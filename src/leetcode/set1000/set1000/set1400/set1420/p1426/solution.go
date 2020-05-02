package p1426

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max int

	for i := 0; i < len(candies); i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	res := make([]bool, len(candies))

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			res[i] = true
		}
	}

	return res
}
