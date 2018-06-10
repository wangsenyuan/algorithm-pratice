package p850

func loudAndRich(richer [][]int, quiet []int) []int {
	n := len(quiet)

	graph := make([]map[int]bool, n)
	for i := 0; i < n; i++ {
		graph[i] = make(map[int]bool)
	}

	for _, r := range richer {
		// x is richer than y
		x, y := r[0], r[1]
		graph[y][x] = true
	}

	que := make([]int, n)
	seen := make([]int, n)

	bfs := func(x int, label int) int {
		front, end := 0, 0
		que[end] = x
		end++
		seen[x] = label

		ans := x

		for front < end {
			ee := end
			for front < ee {
				cur := que[front]
				front++
				if quiet[cur] < quiet[ans] {
					ans = cur
				}
				for y := range graph[cur] {
					if seen[y] == label {
						continue
					}
					seen[y] = label
					que[end] = y
					end++
				}
			}
		}
		return ans
	}

	res := make([]int, n)

	for i := 0; i < n; i++ {
		res[i] = bfs(i, i+1)
	}
	return res
}
