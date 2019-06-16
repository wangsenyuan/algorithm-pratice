package p885

import "sort"

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)

	i, j := 0, len(people)-1
	var ans int
	for i <= j {
		if people[i]+people[j] <= limit {
			i++
		}
		j--
		ans++
	}

	return ans
}
