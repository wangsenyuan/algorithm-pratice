package p3274

func checkTwoChessboards(coordinate1 string, coordinate2 string) bool {
	getColor := func(coord string) int {
		x := int(coord[0] - 'a')
		y := int(coord[1] - '1')
		return (x + y) % 2
	}

	return getColor(coordinate1) == getColor(coordinate2)
}
