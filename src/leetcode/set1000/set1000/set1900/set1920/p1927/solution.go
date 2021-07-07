package p1927

import (
	"math"
	"sort"
)

func minDayskVariants(points [][]int, k int) int {
	S_2 := math.Sqrt(2)

	change := func(p []int) []float64 {
		x := float64(p[0])/S_2 - float64(p[1])/S_2
		y := float64(p[0])/S_2 + float64(p[1])/S_2
		return []float64{x, y}
	}

	n := len(points)
	K1 := (1 << k) - 1
	K2 := K1 << (n - k)

	cur := K1
	var best = math.MaxInt32
	X := make([]float64, k)
	Y := make([]float64, k)
	for cur <= K2 {
		var j int

		for i := 0; i < n; i++ {
			if cur&(1<<i) > 0 {
				p := change(points[i])
				X[j] = p[0]
				Y[j] = p[1]
				j++
			}
		}

		sort.Float64s(X)
		sort.Float64s(Y)

		x_c := (X[0] + X[k-1]) / 2
		y_c := (Y[0] + Y[k-1]) / 2

		x := int(x_c/S_2 + y_c/S_2)
		y := int(-x_c/S_2 + y_c/S_2)

		U := []int{x, x + 1}
		V := []int{y, y + 1}

		for a := 0; a < 2; a++ {
			for b := 0; b < 2; b++ {
				var tmp int
				for i := 0; i < n; i++ {
					if cur&(1<<i) > 0 {
						p := points[i]
						dis := abs(p[0]-U[a]) + abs(p[1]-V[b])
						if dis > tmp {
							tmp = dis
						}
					}
				}
				if tmp < best {
					best = tmp
				}
			}
		}

		cur = nextNum(cur)
		if cur <= K1 {
			break
		}
	}

	return best
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func nextNum(x int) int {
	// right most set bit
	rightOne := x & -x

	// reset the pattern and set next higher bit
	// left part of x will be here
	nextHigherOneBit := x + rightOne

	// nextHigherOneBit is now part [D] of the above explanation.

	// isolate the pattern
	rightOnesPattern := x ^ nextHigherOneBit

	// right adjust pattern
	rightOnesPattern = (rightOnesPattern) / rightOne

	// correction factor
	rightOnesPattern >>= 2

	// rightOnesPattern is now part [A] of the above explanation.

	// integrate new pattern (Add [D] and [A])
	next := nextHigherOneBit | rightOnesPattern

	return next
}
