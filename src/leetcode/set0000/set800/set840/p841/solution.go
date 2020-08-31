package p841

func canVisitAllRooms(rooms [][]int) bool {
	n := len(rooms)
	open := make([]bool, n)
	open[0] = true

	stack := make([]int, n)
	var p int
	stack[p] = 0
	p++

	for p > 0 {
		cur := stack[p-1]
		p--
		for _, next := range rooms[cur] {
			if !open[next] {
				open[next] = true
				stack[p] = next
				p++
			}
		}
	}

	for i := 0; i < n; i++ {
		if !open[i] {
			return false
		}
	}
	return true
}
