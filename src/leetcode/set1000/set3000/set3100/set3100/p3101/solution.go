package p3101

func maxBottlesDrunk(numBottles int, numExchange int) int {
	var res int
	var empty int
	for numBottles+empty >= numExchange {
		res += numBottles
		empty += numBottles
		if empty < numExchange {
			break
		}
		numBottles = 1
		empty -= numExchange
		numExchange++
	}

	return res + numBottles
}
