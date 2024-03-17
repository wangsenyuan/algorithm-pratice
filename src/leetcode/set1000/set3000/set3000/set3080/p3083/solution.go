package p3083

import (
	"container/heap"
	"sort"
)

func minimizeStringValue(s string) string {
	n := len(s)
	// 前后的字符的计数，都会影响到当前?的选择
	// 且在当前位置的选择，也会影响到那些后续要选择的位置
	// 假设当前位置选择了一个a，后面也选择了一个a，那么会产生额外的一次1
	// 但是，假设要选择最终要产生x个a，那么这x个a肯定是在前面的？
	// 是不是可以用dp了，
	// dp[i][x]表示前i个？，最后一个用x
	// 不行的，因为不确定前面用了多少个x，如果当前再使用x的时候，就无法计算当前的cost
	// 突然发现，这个cost其实和位置没关系，只会个数有关系，那就是尽量的使得个数相等
	// 且使用不存在的部分
	cnt := make([]int, 26)
	var mark []int
	for i := 0; i < n; i++ {
		if s[i] == '?' {
			mark = append(mark, i)
		} else {
			cnt[int(s[i]-'a')]++
		}
	}
	if len(mark) == 0 {
		return s
	}

	// 然后让最小值尽量的多
	pq := make(PriorityQueue, 26)
	for i := 0; i < 26; i++ {
		it := new(Item)
		it.id = i
		it.value = cnt[i]
		it.index = i
		pq[i] = it
	}
	heap.Init(&pq)

	var arr []int

	for i := 0; i < len(mark); i++ {
		it := pq[0]
		arr = append(arr, it.id)
		pq.update(it, it.value+1)
	}

	sort.Ints(arr)

	buf := []byte(s)

	for i := 0; i < len(mark); i++ {
		buf[mark[i]] = byte('a' + arr[i])
	}

	return string(buf)
}

type Pair struct {
	first  int
	second int
}

// An Item is something we manage in a priority queue.
type Item struct {
	id    int
	value int
	index int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value || pq[i].value == pq[j].value && pq[i].id < pq[j].id
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int) {
	item.value = value
	heap.Fix(pq, item.index)
}
