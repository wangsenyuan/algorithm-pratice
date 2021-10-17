package p2037

import "sort"

func minMovesToSeat(seats []int, students []int) int {
	n := len(seats)
	sort.Ints(seats)
	sort.Ints(students)

	var res int

	for i := 0; i < n; i++ {
		res += abs(students[i] - seats[i])
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
