package p1699

func countStudents(students []int, sandwiches []int) int {
	que := NewQueue(students)
	var cnt int
	for i := 0; i < len(sandwiches); i++ {
		cur := que.PopFront()
		if cur == sandwiches[i] {
			cnt = 0
			continue
		}
		i--
		que.OfferEnd(cur)
		cnt++
		if cnt == que.Size() {
			break
		}
	}
	return que.Size()
}

type Queue struct {
	arr        []int
	front, end int
	size       int
}

func NewQueue(arr []int) *Queue {
	tmp := make([]int, len(arr)+1)
	copy(tmp, arr)
	return &Queue{tmp, 0, len(arr), len(arr)}
}

func (que Queue) GetFront() int {
	return que.arr[que.front]
}

func (que *Queue) PopFront() int {
	res := que.arr[que.front]
	que.front++
	if que.front == len(que.arr) {
		que.front = 0
	}
	que.size--
	return res
}

func (que *Queue) OfferEnd(x int) {
	que.arr[que.end] = x
	que.end++
	if que.end == len(que.arr) {
		que.end = 0
	}
	que.size++
}

func (que Queue) Size() int {
	return que.size
}
