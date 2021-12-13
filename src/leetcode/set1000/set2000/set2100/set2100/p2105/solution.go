package p2105

func minimumRefill(plants []int, capacityA int, capacityB int) int {
	var res int
	n := len(plants)
	a, b := capacityA, capacityB

	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if i == j {
			if a >= b {
				// alice
				if a < plants[i] {
					res++
				}
			} else {
				if b < plants[i] {
					res++
				}
			}
			continue
		}
		if a < plants[i] {
			res++
			a = capacityA
		}
		a -= plants[i]
		if b < plants[j] {
			res++
			b = capacityB
		}
		b -= plants[j]
	}

	return res
}
