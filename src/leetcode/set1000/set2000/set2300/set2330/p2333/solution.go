package p2333

import "sort"

func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	k := k1 + k2
	n := len(nums1)
	// to min the result,
	// need to make a & b, closer
	// a, b, c, d
	// to make a, b closer, and c, d
	// (a + 1 - b) ** 2

	// to make every pair small enough
	num := make([]Pair, n)
	var sum int64
	for i := 0; i < n; i++ {
		num[i] = Pair{nums1[i], nums2[i]}
		sum += square(nums1[i] - nums2[i])
	}
	sort.Slice(num, func(i, j int) bool {
		return num[i].Less(num[j])
	})

	// < 0 的部qwert77890分让他们增加到接近0
	// 每a增加1或者b减少1 减少量是 2 * b - 2 * a -1
	for i := 0; i < n && k > 0; {
		// x * x + 2 * (b - a) * x
		d := abs(num[i].first - num[i].second)
		for i < n && abs(num[i].first-num[i].second) == d {
			i++
		}
		x := d
		if i < n {
			x = d - abs(num[i].first-num[i].second)
		}
		// 可以将d减少到floor
		if int64(x)*int64(i) <= int64(k) {
			sum += int64(i) * (square(x) - 2*int64(d)*int64(x))
			k -= x * i
		} else {
			// x * (i - j + 1) > K
			for k > 0 {
				y := k / i
				if y > 0 {
					sum += int64(i) * (square(y) - 2*int64(d)*int64(y))
					k -= y * (i)
					d -= y
				} else {
					// y == 0
					l := k
					// y = 1
					sum += int64(l) * (1 - 2*int64(d))
					k = 0
				}
			}
		}
	}

	return sum
}

func square(num int) int64 {
	return int64(num) * int64(num)
}

type Pair struct {
	first  int
	second int
}

func (this Pair) Less(that Pair) bool {
	x := abs(this.first - this.second)
	y := abs(that.first - that.second)
	return x > y
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
