package p2079

func wateringPlants(plants []int, capacity int) int {
	// steps
	n := len(plants)
	// n <= 1000
	var res int
	for i := 0; i < n; {
		j := i
		var sum int
		for j < n && sum+plants[j] <= capacity {
			sum += plants[j]
			j++
		}
		// j == n || sum + plants[j] > capacity
		res -= 2
		res += 2 * (i + 1)
		res += (j - i)
		i = j
	}

	return res
}
