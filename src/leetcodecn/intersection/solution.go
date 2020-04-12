package intersection

func intersection(start1 []int, end1 []int, start2 []int, end2 []int) []float64 {
	det := func(a, b []int) int {
		return a[0]*b[1] - a[1]*b[0]
	}

	xdiff := []int{start1[0] - end1[0], start2[0] - end2[0]}
	ydiff := []int{start1[1] - end1[1], start2[1] - end2[1]}

	div := det(xdiff, ydiff)

	if div == 0 {
		return nil
	}

	d := []int{det(start1, end1), det(start2, end2)}

	x := det(d, xdiff)
	y := det(d, ydiff)

	return []float64{float64(x) / float64(div), float64(y) / float64(div)}
}
