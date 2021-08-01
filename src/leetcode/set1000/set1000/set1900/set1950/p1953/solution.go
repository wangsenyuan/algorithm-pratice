package p1953

func minimumPerimeter(neededApples int64) int64 {
	// suppose there n cols
	// then first is 1, 2, 3, ... n + 1 => (1 + n + 1) * (n + 1) / 2
	// second is 2, 3, 4. ..... n + 2 => (2 + n + 2) * (n+1) / 2
	// ...
	// last is  n, n + 1, ...... n + n => (n + n + n) * (n + 1) / 2
	// sum => ((n + 1) * n / 2 + (n + 1 + 2 *n ) * n/2) * (n + 1) / 2
	// n * (n + 1) * ((n + 1) + 3 * n + 1) / 4
	// n * (n + 1) * (4 * n + 2)
	// sum *= 4
	// sum = 2 * ((1 + n) * n / 2 + (3 * n + 1) * n / 2)
	// = (1 + n) * n + (3 * n + 1) * n = (4 * n + 2) * n
	if neededApples == 1 {
		return 8
	}
	var left, right int64 = 0, 504000
	for left < right {
		mid := (left + right) / 2
		if count(mid) >= neededApples {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return 8 * right
}

func count(n int64) int64 {
	return n * (n + 1) * (4*n + 2)
}
