package p853

import (
	"sort"
)

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}
	cars := make(Cars, n)
	for i := 0; i < n; i++ {
		cars[i] = &Car{float64(speed[i]), float64(position[i])}
	}
	sort.Sort(cars)

	tt := float64(target)
	var ans int

	for i := n - 2; i >= 0; i-- {
		if cars[i].speed <= cars[i+1].speed {
			// can't catch up the next
			ans++
			continue
		}
		time := (cars[i+1].position - cars[i].position) / (cars[i].speed - cars[i+1].speed)
		nextPos := cars[i+1].position + cars[i+1].speed*time

		if nextPos > tt {
			// can't catch up next
			ans++
			continue
		}

		cars[i] = cars[i+1]
	}

	return ans + 1
}

type Car struct {
	speed    float64
	position float64
}

type Cars []*Car

func (cars Cars) Len() int {
	return len(cars)
}

func (cars Cars) Less(i, j int) bool {
	return cars[i].position < cars[j].position
}

func (cars Cars) Swap(i, j int) {
	cars[i], cars[j] = cars[j], cars[i]
}
