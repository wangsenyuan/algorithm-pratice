package p743

import "container/heap"

func networkDelayTime(times [][]int, N int, K int) int {
	K--

	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, N)
		for j := 0; j < N; j++ {
			grid[i][j] = -1
		}
	}

	for _, time := range times {
		u, v, w := time[0]-1, time[1]-1, time[2]
		if grid[u][v] == -1 || grid[u][v] > w {
			grid[u][v] = w
		}
	}
	dist := make([]int, N)

	for i := 0; i < N; i++ {
		dist[i] = -1
	}

	dist[K] = 0

	que := Items(make([]Item, 0, N))
	heap.Init(&que)
	heap.Push(&que, Item{K, 0})

	checked := make([]bool, N)

	for que.Len() > 0 {
		cur := heap.Pop(&que).(Item)
		v, d := cur.node, cur.dist
		if checked[v] {
			continue
		}
		checked[v] = true
		for i := 0; i < N; i++ {
			if grid[v][i] >= 0 && (dist[i] == -1 || dist[i] > d+grid[v][i]) {
				dist[i] = d + grid[v][i]
				heap.Push(&que, Item{i, dist[i]})
			}
		}
	}
	var ans int
	for i := 0; i < N; i++ {
		if dist[i] < 0 {
			return -1
		}
		if dist[i] > ans {
			ans = dist[i]
		}
	}
	return ans
}

type Item struct {
	node int
	dist int
}

type Items []Item

func (items Items) Len() int {
	return len(items)
}

func (items Items) Less(i, j int) bool {
	return items[i].dist < items[j].dist
}

func (items Items) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

func (items *Items) Push(x interface{}) {
	*items = append(*items, x.(Item))
}

func (items *Items) Pop() interface{} {
	n := len(*items)
	last := (*items)[n-1]
	*items = (*items)[:n-1]
	return last
}
