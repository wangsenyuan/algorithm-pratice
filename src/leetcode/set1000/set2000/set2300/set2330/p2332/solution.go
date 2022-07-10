package p2332

import "sort"

func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	m := len(buses)
	n := len(passengers)

	mem := make(map[int]bool)

	for _, num := range passengers {
		mem[num] = true
	}

	// a 先到，但是没有赶上的话，也需要上车
	var ans int
	var k int
	var j int
	for i := 0; i < n && j < m; i++ {
		if passengers[i] <= buses[j] {
			if i == 0 || passengers[i-1]+1 != passengers[i] {
				ans = max(ans, passengers[i]-1)
			}
			k++

			if k+1 <= capacity {
				if passengers[i] < buses[j] && (i+1 == n || passengers[i+1] > buses[j]) {
					ans = buses[j]
				}
				// can arrive at buses[j]
			}
			if k == capacity {
				k = 0
				j++
			}
		} else {
			// pass[i] > bus[j]
			if k < capacity && !mem[buses[j]] {
				ans = buses[j]
			}
			i--
			k = 0
			j++
		}
	}

	for j < m {
		if k < capacity && !mem[buses[j]] {
			ans = buses[j]
		}
		k = 0
		j++
	}

	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
