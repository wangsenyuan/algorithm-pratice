package main

import "fmt"

func main() {
	gas := []int{4}
	cost := []int{5}
	fmt.Println(canCompleteCircuit(gas, cost))
}

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	for i := 0; i < n; {
		left := gas[i]
		j := i
		for k := 0; k < n; k++ {
			left -= cost[j]
			j = (j + 1) % n
			if left < 0 || j == i {
				break
			}
			left += gas[j]
		}

		if left >= 0 {
			return i
		}
		if j <= i {
			return -1
		}

		i = j
	}
	return -1
}
