package p1438

func longestSubarray(nums []int, limit int) int {
	n := len(nums)

	maxQueue := NewDQueue(n)
	minQueue := NewDQueue(n)

	var best int

	for i, j := 0, 0; i < n; i++ {
		for !maxQueue.IsEmpty() && nums[maxQueue.PeekLast()] < nums[i] {
			maxQueue.PopLast()
		}

		maxQueue.PushLast(i)

		for !minQueue.IsEmpty() && nums[minQueue.PeekLast()] > nums[i] {
			minQueue.PopLast()
		}

		minQueue.PushLast(i)

		for nums[maxQueue.PeekFirst()]-nums[minQueue.PeekFirst()] > limit {
			if j == maxQueue.PeekFirst() {
				maxQueue.PopFirst()
			}
			if j == minQueue.PeekFirst() {
				minQueue.PopFirst()
			}
			j++
		}

		if best < i-j+1 {
			best = i - j + 1
		}
	}

	return best
}

type DQueue struct {
	arr        []int
	front, end int
}

func NewDQueue(n int) DQueue {
	arr := make([]int, n)
	return DQueue{arr, 0, 0}
}

func (que *DQueue) PushLast(v int) {
	que.arr[que.end] = v
	que.end++
}

func (que *DQueue) PopLast() int {
	res := que.arr[que.end-1]
	que.end--
	return res
}

func (que DQueue) PeekLast() int {
	return que.arr[que.end-1]
}

func (que DQueue) IsEmpty() bool {
	return que.front == que.end
}

func (que *DQueue) PeekFirst() int {
	return que.arr[que.front]
}

func (que *DQueue) PopFirst() int {
	res := que.arr[que.front]
	que.front++
	return res
}

func longestSubarray1(nums []int, limit int) int {
	n := len(nums)

	arr1 := make([]int, 2*n)
	arr2 := make([]int, 2*n)
	for i := n; i < len(arr1); i++ {
		arr1[i] = nums[i-n]
		arr2[i] = nums[i-n]
	}

	for i := n - 1; i > 0; i-- {
		arr1[i] = min(arr1[2*i], arr1[2*i+1])
		arr2[i] = max(arr2[2*i], arr2[2*i+1])
	}

	get := func(arr []int, l, r int, res int, fn func(int, int) int) int {
		l += n
		r += n
		// var res = math.MaxInt32
		for l < r {
			if l&1 == 1 {
				res = fn(res, arr[l])
				l++
			}
			if r&1 == 1 {
				r--
				res = fn(res, arr[r])
			}
			l >>= 1
			r >>= 1
		}
		return res
	}

	check := func(k int) bool {
		for i := 0; i+k <= n; i++ {
			a := get(arr2, i, i+k, 0, max)
			b := get(arr1, i, i+k, 1<<30, min)
			if a-b <= limit {
				return true
			}
		}
		return false
	}

	var left, right = 1, n + 1
	for left < right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right - 1
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
