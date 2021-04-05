package p1812

func squareIsWhite(coordinates string) bool {
	x := int(coordinates[0] - 'a')
	y := int(coordinates[1] - '1')

	return (x+y)&1 == 1
}
