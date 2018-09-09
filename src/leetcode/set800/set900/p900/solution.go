package p900

type RLEIterator struct {
	arr   []int
	pos   []int
	index int
	used  int
}

func Constructor(A []int) RLEIterator {
	var j int
	for i := 0; i < len(A); i += 2 {
		if A[i] == 0 {
			continue
		}
		A[j] = A[i]
		A[j+1] = A[i+1]
		j += 2
	}

	pos := make([]int, j)
	for i := 2; i < j; i += 2 {
		pos[i] = A[i-2] + pos[i-2]
	}

	return RLEIterator{A[:j], pos, 0, 0}
}

func (this *RLEIterator) Next(n int) int {
	i := this.index
	arr := this.arr
	pos := this.pos
	if i >= len(arr) {
		return -1
	}
	this.used += n

	for i < len(arr) && pos[i]+arr[i] < this.used {
		i += 2
	}

	if i >= len(arr) {
		return -1
	}

	return arr[i+1]
}
