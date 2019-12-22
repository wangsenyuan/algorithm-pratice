package p1298

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)

	que1 := NewQueue(n + 1)
	que2 := NewQueue(n + 1)
	que3 := NewQueue(n + 1)

	for _, i := range initialBoxes {
		if status[i] == 1 {
			que1.Offer(i)
		} else {
			que2.Offer(i)
		}
	}

	hasKey := make([]bool, n)

	var got int

	for !que1.IsEmpty() {
		cur := que1.Pull()
		got += candies[cur]
		for _, k := range keys[cur] {
			hasKey[k] = true
		}

		for _, b := range containedBoxes[cur] {
			if status[b] == 1 {
				que1.Offer(b)
			} else {
				que2.Offer(b)
			}
		}

		for !que2.IsEmpty() {
			x := que2.Pull()
			if hasKey[x] {
				que1.Offer(x)
			} else {
				que3.Offer(x)
			}
		}

		for !que3.IsEmpty() {
			que2.Offer(que3.Pull())
		}
	}

	return got
}

type Queue struct {
	size  int
	front int
	end   int
	arr   []int
}

func NewQueue(n int) Queue {
	arr := make([]int, n)
	return Queue{n, 0, 0, arr}
}

func (q *Queue) Offer(x int) {
	q.arr[q.end] = x
	q.end = (q.end + 1) % q.size
}

func (q *Queue) Pull() int {
	res := q.arr[q.front]
	q.front = (q.front + 1) % q.size
	return res
}

func (q Queue) IsEmpty() bool {
	return q.front == q.end
}
