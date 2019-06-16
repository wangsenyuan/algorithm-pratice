package p841

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	que := make([]int, n)
	front, rear := 0, 0

	que[0] = 0
	rear++
	d := make([]int, n)
	for i := 0; i < n; i++ {
		d[i] = -1
	}
	d[0] = 0

	for front < rear {
		tt := rear

		for front < tt {
			u := que[front]
			front++
			for _, v := range rooms[u] {
				if d[v] < 0 {
					d[v] = d[u] + 1
					que[rear] = v
					rear++
				}
			}
		}
	}
	return rear == n
}
