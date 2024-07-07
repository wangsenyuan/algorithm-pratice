package p3206

import (
	"slices"
)

func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	mn := slices.Min(enemyEnergies)

	if currentEnergy < mn {
		return 0
	}

	sum := currentEnergy - mn

	for _, num := range enemyEnergies {
		sum += num
	}

	sum -= mn

	res := sum / mn
	res++
	return int64(res)
}
