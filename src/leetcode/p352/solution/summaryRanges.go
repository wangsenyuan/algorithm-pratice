package solution

import "fmt"

//SummaryRanges is used to get the disjoint set
type SummaryRanges struct {
	bounds map[int]int
	nums   []int
}

// Constructor is used Initialize your data structure here
func Constructor() SummaryRanges {
	return SummaryRanges{bounds: make(map[int]int), nums: make([]int, 0, 0)}
}

//Addnum add one more element
func (sr *SummaryRanges) Addnum(val int) {
	if len(sr.nums) == 0 {
		sr.nums = append(sr.nums, val)
		sr.bounds[val] = val
		return
	}

	i := lowerBound(sr.nums, val)

	if i >= 0 && sr.bounds[sr.nums[i]] >= val {
		return
	}

	if i >= 0 && sr.bounds[sr.nums[i]] + 1 == val {
		if i < len(sr.nums) - 1 && val + 1 == sr.nums[i + 1] {
			sr.bounds[sr.nums[i]] = sr.bounds[sr.nums[i + 1]]
			delete(sr.bounds, sr.nums[i + 1])
			sr.nums = deleteAt(sr.nums, i + 1)
		} else {
			sr.bounds[sr.nums[i]] = val
		}
		return
	}

	j := i + 1
	if j < len(sr.nums) && val + 1 == sr.nums[j] {
		sr.bounds[val] = sr.bounds[sr.nums[j]]
		delete(sr.bounds, sr.nums[j])
		sr.nums = replaceAt(sr.nums, j, val)
		return
	}

	sr.nums = insertAt(sr.nums, i + 1, val)
	sr.bounds[val] = val
}

func insertAt(nums []int, i int, val int) []int {
	xs := make([]int, i, i + 1)
	copy(xs, nums[0:i])
	xs = append(xs, val)
	return append(xs, nums[i:]...)
}

func replaceAt(nums []int, i int, val int) []int {
	xs := make([]int, i, i + 1)
	copy(xs, nums[0:i])
	xs = append(xs, val)
	return append(xs, nums[i + 1:]...)
}

func deleteAt(nums []int, i int) []int {
	xs := make([]int, i, i + 1)
	copy(xs, nums[0:i])
	return append(xs, nums[i + 1:]...)
}

func lowerBound(nums []int, num int) int {
	i, j := -1, len(nums)

	for i + 1 < j {
		mid := (i + j) / 2
		if nums[mid] == num {
			return mid
		}
		if nums[mid] > num {
			j = mid
		} else {
			i = mid
		}
	}

	return i
}

//Getintervals get the interverl result
func (sr *SummaryRanges) Getintervals() []Interval {
	result := make([]Interval, 0, len(sr.nums))
	for _, lo := range sr.nums {
		hi := sr.bounds[lo]
		result = append(result, Interval{lo, hi})
	}

	return result
}

//Interval is from leetcode
type Interval struct {
	Start int
	End   int
}

func (x Interval) String() string {
	return fmt.Sprintf("[%d, %d]", x.Start, x.End)
}
