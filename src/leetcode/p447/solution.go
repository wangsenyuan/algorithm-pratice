package main

import "fmt"

func main() {
	points := [][]int{[]int{0, 0}, []int{1, 0}, []int{2, 0}}
	fmt.Println(numberOfBoomerangs(points))
}

func numberOfBoomerangs(points [][]int) int {
	if len(points) < 3 {
		return 0
	}

	ans := 0
	for i, a := range points {
		dist := make(map[int64]int)
		for j, b := range points {
			if i == j {
				continue
			}

			d := distance(a, b)
			if dist[d] > 0 {
				ans += 2 * dist[d]
			}
			dist[d]++
		}
	}

	return ans
}

func distance(a, b []int) int64 {
	x := int64(b[0] - a[0])
	y := int64(b[1] - a[1])

	return x*x + y*y
}
