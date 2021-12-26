package p2125

import "sort"

func getDistances(arr []int) []int64 {
	n := len(arr)

	P := make([]Pair, n)

	for i := 0; i < n; i++ {
		P[i] = Pair{arr[i], i}
	}

	sort.Slice(P, func(i, j int) bool {
		return P[i].Less(P[j])
	})

	ans := make([]int64, n)

	handle := func(x, y int) {
		var sum int64 = int64(P[x].second)

		for i := x + 1; i <= y; i++ {
			j := P[i].second
			ans[j] = int64(i-x)*int64(j) - sum
			sum += int64(j)
		}

		sum = int64(P[y].second)

		for i := y - 1; i >= x; i-- {
			j := P[i].second
			ans[j] += sum - int64(y-i)*int64(j)
			sum += int64(j)
		}

	}

	for i, j := 1, 0; i <= n; i++ {
		if i == n || P[i].first > P[i-1].first {
			handle(j, i-1)
			j = i
		}
	}

	return ans
}

type Pair struct {
	first, second int
}

func (this Pair) Less(that Pair) bool {
	return this.first < that.first || this.first == that.first && this.second < that.second
}
