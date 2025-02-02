package p3441

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	type pair struct {
		first  int
		second int
	}
	n := len(startTime)
	arr := make([]pair, n)
	for i := range n {
		arr[i] = pair{startTime[i], endTime[i]}
	}
	// 这里有k这个限制
	// 前面的都要往前移动，后面的都要往后移动
	// 假设前面移动x个，那么前面倒数第x的起点就是起点，移动的距离 = 当前的终点 - 第x的起点 - sum(x个)
	// 这k个肯定是连续的， 最后一个的终点 - 第一个起点 - sum...
	// 用个que就可以了
	var best int
	que := make([]int, n)
	var head, tail int
	var sum int
	for i := 0; i < n; i++ {
		que[head] = i
		head++
		sum += arr[i].second - arr[i].first
		if head-tail > k {
			j := que[tail]
			tail++
			sum -= arr[j].second - arr[j].first
		}
		prevPos := 0
		if que[tail] > 0 {
			prevPos = arr[que[tail]-1].second
		}
		// 少于等于k个， 全部可以移动
		nextPos := eventTime
		if i+1 < n {
			nextPos = arr[i+1].first
		}
		best = max(best, nextPos-prevPos-sum)
	}

	return best
}
