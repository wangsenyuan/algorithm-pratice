package p1095

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 */
type MountainArray struct {
	arr []int
}

func (this *MountainArray) get(index int) int {
	return this.arr[index]
}
func (this *MountainArray) length() int {
	return len(this.arr)
}

func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()

	if n < 3 {
		return -1
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = -1
	}
	left, right := 0, n-1

	getValue := func(i int) int {
		if arr[i] >= 0 {
			return arr[i]
		}
		arr[i] = mountainArr.get(i)
		return arr[i]
	}

	for left < right {
		mid := (left + right) / 2

		if getValue(mid) < getValue(mid+1) {
			// pivot at right
			left = mid + 1
		} else {
			right = mid
		}
	}
	pivot := left

	if target > getValue(pivot) {
		return -1
	}

	if target == getValue(pivot) {
		return pivot
	}

	left = 0
	right = pivot
	for left < right {
		mid := (left + right) / 2
		if getValue(mid) == target {
			return mid
		}
		if getValue(mid) > target {
			right = mid
		} else {
			left = mid + 1
		}
	}

	left = pivot + 1
	right = n
	for left < right {
		mid := (left + right) / 2
		if getValue(mid) == target {
			return mid
		}
		if getValue(mid) < target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return -1
}
