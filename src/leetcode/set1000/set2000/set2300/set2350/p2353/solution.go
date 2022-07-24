package p2353

import "container/heap"

type FoodRatings struct {
	roots map[string]*PriorityQueue
	index map[string]*Item
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	roots := make(map[string]*PriorityQueue)
	index := make(map[string]*Item)

	for i := 0; i < len(foods); i++ {
		food := foods[i]
		cui := cuisines[i]
		rating := ratings[i]
		if _, ok := roots[cui]; !ok {
			pq := make(PriorityQueue, 0, 1)
			roots[cui] = &pq
		}
		pq := roots[cui]
		item := &Item{food, cui, rating, len(*pq)}
		heap.Push(pq, item)

		index[food] = item
	}

	return FoodRatings{roots, index}
}

func (this *FoodRatings) ChangeRating(food string, newRating int) {
	item := this.index[food]
	pq := this.roots[item.cuisine]
	item.priority = newRating
	heap.Fix(pq, item.index)
}

func (this *FoodRatings) HighestRated(cuisine string) string {
	pq := this.roots[cuisine]
	res := (*pq)[0]
	return res.food
}

// An Item is something we manage in a priority queue.
type Item struct {
	food     string
	cuisine  string
	priority int
	index    int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].food < pq[j].food
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

/**
 * Your FoodRatings object will be instantiated and called as such:
 * obj := Constructor(foods, cuisines, ratings);
 * obj.ChangeRating(food,newRating);
 * param_2 := obj.HighestRated(cuisine);
 */
