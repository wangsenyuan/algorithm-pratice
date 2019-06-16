package main

import "fmt"

func isValid(a, b, c, d []int) bool {
	x1, y1 := a[0], a[1]
	x2, y2 := b[0], b[1]

	if x1 == x2 && y1 == y2 {
		return false
	}

	x3, y3 := x1+y1-y2, y1+x2-x1
	x4, y4 := x2+y1-y2, y2+x2-x1

	return c[0] == x4 && c[1] == y4 && d[0] == x3 && d[1] == y3
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	ps := [][]int{p1, p2, p3, p4}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}

			for k := 0; k < 4; k++ {
				if i == k || j == k {
					continue
				}
				l := 6 - i - j - k
				if isValid(ps[i], ps[j], ps[k], ps[l]) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	p1 := []int{0, 0}
	p2 := []int{1, 1}
	p3 := []int{1, 0}
	p4 := []int{0, 1}
	fmt.Println(validSquare(p1, p2, p3, p4))
}
