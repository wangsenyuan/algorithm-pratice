package p2960

func countTestedDevices(batteryPercentages []int) int {
	var res int
	var dec int
	for i := 0; i < len(batteryPercentages); i++ {
		if batteryPercentages[i]-dec > 0 {
			res++
			dec++
		}
	}

	return res
}
