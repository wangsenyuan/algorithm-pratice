package p933

type RecentCounter struct {
	que        []int
	front, end int
	size       int
}

func Constructor() RecentCounter {
	que := make([]int, 4000)
	return RecentCounter{que, 0, 0, 4000}
}

func (this *RecentCounter) Ping(t int) int {
	front, end := this.front, this.end

	size := this.size
	que := this.que

	que[end] = t
	end = (end + 1) % size
	this.end = end

	prev := t - 3000

	for front != end && que[front] < prev {
		front = (front + 1) % size
	}
	this.front = front

	if front <= end {
		return end - front
	}
	return (end + size - front)
}
