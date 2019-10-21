package p1228

func missingNumber(arr []int) int {
	if len(arr) == 2 {
		return (arr[1]-arr[0])/2 + arr[0]
	}

	a := arr[1] - arr[0]
	b := arr[2] - arr[1]
	if a == b {
		i := 3
		for i < len(arr) && arr[i] == arr[i-1]+a {
			i++
		}
		// arr[i] != arr[i-1] + a
		return arr[i-1] + a
	}

	// a != b
	// x = arr[0] + b
	// y = arr[2] - a

	n := len(arr)
	if n > 3 {
		if arr[n-1]-arr[n-2] == b {
			return arr[0] + b
		}
		return arr[1] + a
	}
	// n == 3
	if a*2 == b {
		return arr[1] + a
	}
	return arr[0] + b
}
