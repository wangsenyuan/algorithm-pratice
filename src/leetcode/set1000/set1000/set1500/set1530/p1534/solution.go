package p1534

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var res int

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[j]-arr[k]) > b || abs(arr[i]-arr[k]) > c {
					continue
				}
				res++
			}
		}
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
