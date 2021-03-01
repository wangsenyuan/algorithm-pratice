package p1776

const INF = 1 << 30

func getCollisionTimes(cars [][]int) []float64 {
	n := len(cars)

	time := func(a, b []int) float64 {
		if a[1] <= b[1] {
			return INF
		}
		return float64(b[0]-a[0]) / float64(a[1]-b[1])
	}

	res := make([]float64, n)
	for i := 0; i < n; i++ {
		res[i] = -1
	}

	stack := make([]int, n)
	var p int

	for i := n - 1; i >= 0; i-- {
		cur := cars[i]
		if p > 0 {
			need := time(cur, cars[stack[p-1]])
			for p > 1 {
				tmp := time(cur, cars[stack[p-2]])
				if tmp > need {
					break
				}
				p--
				need = tmp
			}
			if need < INF {
				res[i] = need
			}
		}
		stack[p] = i
		p++
	}

	return res
}
