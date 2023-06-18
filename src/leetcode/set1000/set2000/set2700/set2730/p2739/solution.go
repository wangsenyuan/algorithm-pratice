package p2739

func distanceTraveled(mainTank int, additionalTank int) int {
	var res int

	for mainTank >= 5 {
		res += 5 * 10
		mainTank -= 5
		if additionalTank > 0 {
			mainTank++
			additionalTank--
		}
	}
	res += mainTank * 10

	return res
}
