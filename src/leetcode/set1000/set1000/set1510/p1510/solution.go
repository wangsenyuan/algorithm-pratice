package p1510

func numWaterBottles(numBottles int, numExchange int) int {
	var res = numBottles

	for numBottles >= numExchange {
		x, y := numBottles/numExchange, numBottles%numExchange
		res += x
		numBottles = x + y
	}

	return res
}
