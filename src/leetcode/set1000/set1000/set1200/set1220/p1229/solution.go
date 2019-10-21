package p1229

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Sort(Slots(slots1))
	sort.Sort(Slots(slots2))

	var i, j int

	for i < len(slots1) && j < len(slots2) {
		if intersect(slots1[i], slots2[j]) {
			a := max(slots1[i][0], slots2[j][0])
			b := min(slots1[i][1], slots2[j][1])
			if b-a >= duration {
				return []int{a, a + duration}
			}
		}

		if i < len(slots1)-1 && slots1[i+1][0] < slots2[j][1] {
			i++
		} else {
			j++
		}
	}

	return []int{}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func intersect(a, b []int) bool {
	return a[1] > b[0] && b[1] > a[0]
}

type Slots [][]int

func (this Slots) Len() int {
	return len(this)
}

func (this Slots) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this Slots) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}
