package p2366

func taskSchedulerII(tasks []int, space int) int64 {
	prev := make(map[int]int64)
	n := len(tasks)
	S := int64(space) + 1
	var day int64
	for i := 1; i <= n; i++ {
		task := tasks[i-1]
		if j, ok := prev[task]; ok {
			if day+1-j >= S {
				day++
			} else {
				day = S + j
			}
			prev[task] = day
		} else {
			day++
			prev[task] = day
		}
	}

	return day
}
