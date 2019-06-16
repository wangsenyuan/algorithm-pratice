package main

import "fmt"

func main() {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)
}

func wiggleSort(nums []int) {
	i, j, n := 0, len(nums)-1, len(nums)
	half := n / 2

	median := 0
	for i <= j {
		p, a, b := i, i+1, j
		for a <= b {
			if nums[a] >= nums[p] {
				nums[a], nums[b] = nums[b], nums[a]
				b--
			} else {
				a++
			}
		}
		nums[p], nums[b] = nums[b], nums[p]
		if b == half {
			median = nums[b]
			break
		} else if b > half {
			j = b - 1
		} else {
			i = b + 1
		}
	}

	m := func(idx int) int {
		return (2*idx + 1) % (n | 1)
	}

	first, mid, last := 0, 0, n-1

	for mid <= last {
		if nums[m(mid)] > median {
			nums[m(mid)], nums[m(first)] = nums[m(first)], nums[m(mid)]
			first++
			mid++
		} else if nums[m(mid)] < median {
			nums[m(mid)], nums[m(last)] = nums[m(last)], nums[m(mid)]
			last--
		} else {
			mid++
		}
	}
}
