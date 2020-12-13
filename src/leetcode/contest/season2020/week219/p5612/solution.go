package p5612

const INF = 1 << 30

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)

	diff := make([]int, n+1)

	for i := 1; i <= n; i++ {
		diff[i] = diff[i-1]
		if i > 1 && boxes[i-1][0] != boxes[i-2][0] {
			diff[i]++
		}
	}

	W := make([]int, n+1)

	que := make([]int, n+1)
	var front, end int
	que[end] = 0
	end++

	f := make([]int, n+1)
	g := make([]int, n+1)

	for i := 1; i <= n; i++ {
		W[i] = W[i-1] + boxes[i-1][1]

		for front < end && (i-que[front] > maxBoxes || W[i]-W[que[front]] > maxWeight) {
			front++
		}

		f[i] = g[que[front]] + diff[i] + 2

		if i < n {
			g[i] = f[i] - diff[i+1]
			for front < end && g[i] <= g[que[end-1]] {
				end--
			}
			que[end] = i
			end++
		}
	}
	return f[n]
}

func min(a, b Pair) Pair {
	if a.first <= b.first {
		return a
	}
	return b
}

type Pair struct {
	first, second int
}
