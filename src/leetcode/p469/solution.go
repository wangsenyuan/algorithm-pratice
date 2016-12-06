package main

import "fmt"

func main() {
	points := [][]int{[]int{0, 0}, []int{0, 10}, []int{10, 10}, []int{10, 0}, []int{5, 5}}
	fmt.Println(isConvex(points))
}

func isConvex(points [][]int) bool {
	if len(points) < 3 {
		return false
	}
	n := len(points)
	last := 0
	tmp := 0
	for i := 2; i <= len(points)+2; i++ {
		p0, p1, p2 := points[(i-2)%n], points[(i-1)%n], points[i%n]
		tmp = (p1[0]-p0[0])*(p2[1]-p0[1]) - (p2[0]-p0[0])*(p1[1]-p0[1])
		if tmp != 0 {
			if last*tmp < 0 {
				return false
			}
			last = tmp
		}
	}
	return true
}
