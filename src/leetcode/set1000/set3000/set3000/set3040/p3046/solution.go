package p3046

import (
	"container/heap"
	"sort"
)

func earliestSecondToMarkIndices(nums, changeIndices []int) int {
	n, m := len(nums), len(changeIndices)
	firstT := make([]int, n)
	for t := m - 1; t >= 0; t-- {
		firstT[changeIndices[t]-1] = t + 1
	}
	ans := sort.Search(m, func(mx int) bool {
		cnt := 0
		done := make([]bool, n)
		h := hp{}
		for t := mx; t >= 0; t-- {
			i := changeIndices[t] - 1
			v := nums[i]
			if v <= 1 || firstT[i]-1 != t {
				cnt++ // 留着给前面，用来减一/标记
				continue
			}
			if cnt == 0 {
				if len(h) == 0 || v <= h[0].v {
					cnt++ // 留着给前面，用来减一/标记
					continue
				}
				done[heap.Pop(&h).(pair).i] = false
				cnt += 2 // 反悔：一次置 0，一次标记
			}
			done[i] = true
			cnt-- // nums[i] 置 0，然后消耗一次标记的时间
			heap.Push(&h, pair{v, i})
		}
		for i, b := range done {
			if !b {
				cnt -= nums[i] + 1 // 减一的时间，以及标记的时间
			}
		}
		return cnt >= 0
	})
	if ans == m {
		return -1
	}
	return ans + 1
}

type pair struct{ v, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { a, b := h[i], h[j]; return a.v < b.v }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
