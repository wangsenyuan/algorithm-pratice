package p5608

import "sort"

func minimumEffort(tasks [][]int) int {
	n := len(tasks)
	tt := make([]Task, n)

	for i := 0; i < n; i++ {
		tt[i] = Task{tasks[i][0], tasks[i][1]}
	}
	sort.Sort(Tasks(tt))

	var res int
	var have int
	for i := 0; i < n; i++ {
		cur := tt[i]
		if have < cur.need {
			res += cur.need - have
			have = cur.need
		}
		have -= cur.cost
	}

	return res
}

type Task struct {
	cost int
	need int
}

func (task Task) save() int {
	return task.need - task.cost
}

type Tasks []Task

func (tasks Tasks) Len() int {
	return len(tasks)
}

func (tasks Tasks) Less(i, j int) bool {
	return tasks[i].save() > tasks[j].save()
}

func (tasks Tasks) Swap(i, j int) {
	tasks[i], tasks[j] = tasks[j], tasks[i]
}
