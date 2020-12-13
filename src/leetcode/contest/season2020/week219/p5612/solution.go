package p5612

const INF = 1 << 30

func boxDelivering(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
	n := len(boxes)
	N := n + 1
	arr := make([]int, 2*N)

	for i := 0; i < len(arr); i++ {
		arr[i] = INF
	}

	update := func(p int, v int) {
		p += N
		arr[p] = v
		for p > 0 {
			arr[p>>1] = min(arr[p], arr[p^1])
			p >>= 1
		}
	}

	get := func(l int, r int) int {
		l += N
		r += N
		var res = INF
		for l < r {
			if l&1 == 1 {
				res = min(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = min(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	diff := make([]int, n+1)

	for i := 1; i <= n; i++ {
		diff[i] = diff[i-1]
		if i > 1 && boxes[i-1][0] != boxes[i-2][0] {
			diff[i]++
		}
	}
	var ans int
	var w int
	update(0, 0)
	for i, j := 1, 0; i <= n; i++ {
		w += boxes[i-1][1]

		for i-j > maxBoxes || w > maxWeight {
			w -= boxes[j][1]
			j++
		}

		tmp := get(j, i)

		ans = tmp + diff[i] + 2

		if i < n {
			update(i, ans-diff[i+1])
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func boxDelivering1(boxes [][]int, portsCount int, maxBoxes int, maxWeight int) int {
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
	var res int
	g := make([]int, n+1)

	for i := 1; i <= n; i++ {
		W[i] = W[i-1] + boxes[i-1][1]

		for front < end && (i-que[front] > maxBoxes || W[i]-W[que[front]] > maxWeight) {
			front++
		}

		res = g[que[front]] + diff[i] + 2

		if i < n {
			g[i] = res - diff[i+1]
			for front < end && g[i] <= g[que[end-1]] {
				end--
			}
			que[end] = i
			end++
		}
	}
	return res
}
