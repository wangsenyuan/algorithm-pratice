package p1276

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	// x, y for jumbo & small burgers
	// x + y = cheese
	// 4 * x + 2 * y = tom
	// 4 * (ch - y) + 2 * y = tom
	// 4 * ch - - 2* y = tom
	//
	if 4*cheeseSlices < tomatoSlices {
		return nil
	}

	d := 4*cheeseSlices - tomatoSlices
	if d%2 != 0 {
		return nil
	}

	y := d / 2
	x := cheeseSlices - y
	if x < 0 {
		return nil
	}

	return []int{x, y}
}
