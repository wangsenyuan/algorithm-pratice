package p2567

func minMaxDifference(num int) int {
	if num == 0 {
		return 9
	}
	a, b := num, num
	// always need to remap the high num
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num /= 10
	}
	n := len(arr)

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			// map x to y
			var tmp int
			for i := n - 1; i >= 0; i-- {
				if arr[i] == x {
					tmp = tmp*10 + y
				} else {
					tmp = tmp*10 + arr[i]
				}
			}
			if tmp > a {
				a = tmp
			}
			if tmp < b {
				b = tmp
			}
		}
	}
	return a - b
}
