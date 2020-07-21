package p346

type MovingAverage struct {
	windowSize int
	nums       []int
	sum        int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{size, nil, 0}
}

func (this *MovingAverage) Next(val int) float64 {
	this.nums = append(this.nums, val)
	this.sum += val
	if len(this.nums) > this.windowSize {
		this.sum -= this.nums[0]
		this.nums = this.nums[1:]
	}
	return float64(this.sum) / float64(len(this.nums))
}
