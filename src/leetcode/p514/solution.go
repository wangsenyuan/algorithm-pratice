package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findRotateSteps("godding", "godding"))
	fmt.Println(findRotateSteps("godding", "gd"))

}

func findRotateSteps(ring string, key string) int {
	n := len(ring)
	cache := make(map[string]int)

	var process func(ring int, key string) int

	process = func(ringPos int, key string) int {
		if len(key) == 0 {
			return 0
		}

		cacheKey := fmt.Sprintf("%s@%d", key, ringPos)

		if res, cached := cache[cacheKey]; cached {
			return res
		}

		cur := key[0]

		steps := math.MaxInt32

		for i := 0; i < n; i++ {
			j := (ringPos + i) % n
			if cur != ring[j] {
				continue
			}

			x := i
			if 2*i > n {
				x = n - i
			}
			subRes := process(j, key[1:])

			if subRes+x+1 < steps {
				steps = subRes + x + 1
			}
		}

		cache[cacheKey] = steps

		return steps
	}

	return process(0, key)
}
