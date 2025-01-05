package p3408

import "container/heap"

type TaskManager struct {
	pq    *PriorityQueue
	tasks map[int]*Task // taskId -> userId
}

func Constructor(tasks [][]int) TaskManager {
	mgr := new(TaskManager)
	mgr.pq = new(PriorityQueue)
	mgr.tasks = make(map[int]*Task)
	for _, task := range tasks {
		userId, taskId, priority := task[0], task[1], task[2]
		mgr.Add(userId, taskId, priority)
	}
	return *mgr
}

func (this *TaskManager) Add(userId int, taskId int, priority int) {
	task := &Task{taskId, userId, priority, -1}
	heap.Push(this.pq, task)
	this.tasks[taskId] = task
}

func (this *TaskManager) Edit(taskId int, newPriority int) {
	task := this.tasks[taskId]
	this.pq.update(task, newPriority)
}

func (this *TaskManager) Rmv(taskId int) {
	task := this.tasks[taskId]
	this.pq.remove(task)
	delete(this.tasks, taskId)
}

func (this *TaskManager) ExecTop() int {
	if len(*this.pq) == 0 {
		return -1
	}
	task := heap.Pop(this.pq).(*Task)
	return task.uid
}

type Task struct {
	id       int
	uid      int
	priority int
	index    int
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].id > pq[j].id
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	task := x.(*Task)
	task.index = n
	*pq = append(*pq, task)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	task := old[n-1]
	task.index = -1
	*pq = old[:n-1]
	return task
}

const inf = 1 << 60

func (pq *PriorityQueue) update(task *Task, priority int) {
	task.priority = priority
	heap.Fix(pq, task.index)
}

func (pq *PriorityQueue) remove(task *Task) {
	task.priority = inf
	heap.Fix(pq, task.index)
	heap.Pop(pq)
}

/**
 * Your TaskManager object will be instantiated and called as such:
 * obj := Constructor(tasks);
 * obj.Add(userId,taskId,priority);
 * obj.Edit(taskId,newPriority);
 * obj.Rmv(taskId);
 * param_4 := obj.ExecTop();
 */
