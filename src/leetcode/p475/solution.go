package main

import (
	"fmt"
	"sort"
)

func main() {
	//houses := []int{1, 2, 3, 4}
	houses := []int{1, 5}
	//heaters := []int{1, 4}
	heaters := []int{2}
	fmt.Println(findRadius(houses, heaters))
}

func findRadius(houses []int, heaters []int) int {

	sort.Ints(houses)
	sort.Ints(heaters)

	ans := 0
	i, j := 0, 0

	for i < len(houses) {
		for j < len(heaters)-1 && heaters[j] < houses[i] {
			j++
		}

		if heaters[j] < houses[i] {
			ans = max(ans, houses[i]-heaters[j])
			i++
			continue
		}

		dist := heaters[j] - houses[i]
		if j > 0 {
			dist = min(dist, houses[i]-heaters[j - 1])
		}

		ans = max(ans, dist)
		i++
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

/*

func findRadius(houses []int, heaters []int) int {
  sort.Ints(houses)
  sort.Ints(heaters)
  checkRadius := func(r int) bool {
    i, j := 0, 0
    for i < len(heaters) && j < len(houses) {
      x := heaters[i]
      if houses[j] < x && x - houses[j] > r {
        return false
      }
      for j < len(houses) && houses[j] <= r + x {
        j++
      }
      i++
    }
    return j == len(houses)
  }

  l, r :=0, 10000000000
  ans := 0
  for l <= r {
      mid := l + (r - l) / 2
      res := checkRadius(mid)

      if res {
        ans = mid
      }
      if l == mid {
        break
      }
      //fmt.Println(mid)
      //fmt.Printf("%d %d %d\n", l, mid, r)
      if res {
        r = mid
      } else {
        l = mid
      }
  }

  return ans
}
*/
