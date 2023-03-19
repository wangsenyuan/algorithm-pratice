package p2591

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	money -= children

	x := money / 7
	y := money % 7

	if x >= children {
		if y == 0 && x == children {
			return children
		}
		return children - 1
	}

	if y == 3 && x == children-1 {
		x--
	}

	return x
}
