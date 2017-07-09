package main

import (
	"math"
	"fmt"
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	if satisfied(needs) {
		return 0
	}

	best := math.MaxInt32
	tmpNeeds := make([]int, len(needs))
	for _, offer := range special {
		i := 0
		for i < len(needs) {
			if offer[i] > needs[i] {
				break
			}
			tmpNeeds[i] = needs[i] - offer[i]
			i++
		}

		if i == len(needs) {
			//can use this offer
			tmp := offer[i] + shoppingOffers(price, special, tmpNeeds)
			if tmp < best {
				best = tmp
			}
		}
	}

	//try use the single price
	ans := 0

	for i := 0; i < len(needs); i++ {
		ans += needs[i] * price[i]
	}

	if ans < best {
		return ans
	}
	return best
}

func satisfied(needs []int) bool {
	for i := 0; i < len(needs); i++ {
		if needs[i] > 0 {
			return false
		}
	}
	return true
}

func main() {
	//items := []int{2, 5}
	//offers := [][]int{{3, 0, 5}, {1, 2, 10}}
	//needs := []int{3, 2}

	items := []int{2, 3, 4}
	offers := [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}
	needs := []int{1, 2, 1}
	fmt.Println(shoppingOffers(items, offers, needs))
}
