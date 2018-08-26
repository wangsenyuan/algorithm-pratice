package p895

import "container/heap"

const MAX_N = 10000

type FreqStack struct {
	que    PriorityQueue
	mem    map[int]*Item
	insert int
}

func Constructor() FreqStack {
	que := make(PriorityQueue, 0, MAX_N)
	mem := make(map[int]*Item)
	return FreqStack{que, mem, 0}
}

func (this *FreqStack) Push(x int) {
	insert := this.insert
	if item, found := this.mem[x]; found {
		this.que.update(item, item.freq+1, append(item.insert, insert))
	} else {
		item = new(Item)
		item.val = x
		item.freq = 1
		item.insert = make([]int, 0, 10)
		item.insert = append(item.insert, insert)
		heap.Push(&this.que, item)
		this.mem[x] = item
	}

	this.insert++
}

func (this *FreqStack) Pop() int {

	//assume alwasy can pop
	head := this.que[0]
	if head.freq == 1 {
		heap.Pop(&this.que)
		delete(this.mem, head.val)
	} else {
		this.que.update(head, head.freq-1, head.insert[:len(head.insert)-1])
	}
	return head.val
}

type Item struct {
	val    int
	freq   int
	insert []int
	index  int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].freq > pq[j].freq {
		return true
	}

	if pq[i].freq < pq[j].freq {
		return false
	}
	a := pq[i]
	b := pq[j]
	ai := len(a.insert)
	bi := len(b.insert)
	return a.insert[ai-1] > b.insert[bi-1]
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, freq int, insert []int) {
	item.freq = freq
	item.insert = insert
	heap.Fix(pq, item.index)
}
