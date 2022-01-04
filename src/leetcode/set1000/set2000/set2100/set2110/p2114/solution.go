package p2114

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	var sum = int64(mass)

	for i := 0; i < len(asteroids); i++ {
		cur := int64(asteroids[i])
		if sum < cur {
			return false
		}
		sum += cur
	}
	return true
}
