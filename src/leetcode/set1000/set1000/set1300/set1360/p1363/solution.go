package p1363

import (
	"bytes"
	"sort"
)

func largestMultipleOfThree(digits []int) string {
	// sum(digits) % 3 == 0
	sort.Ints(digits)

	nums := make([][]int, 3)
	for i := 0; i < 3; i++ {
		nums[i] = make([]int, 0, len(digits))
	}

	check := func(a, b *[]int) {
		// one 1, or two 2
		if len(*a) > 0 {
			if len(*b) > 1 {
				if (*a)[0] > (*b)[0]+(*b)[1] {
					*b = (*b)[2:]
				} else {
					*a = (*a)[1:]
				}
			} else {
				*a = (*a)[1:]
			}
		} else {
			// nums[2] must have at least two nums
			*b = (*b)[2:]
		}
	}

	var rem int
	n := len(digits)
	for i := 0; i < n; i++ {
		r := digits[i] % 3
		nums[r] = append(nums[r], digits[i])
		rem = (rem + r) % 3
	}

	if rem == 1 {
		check(&nums[1], &nums[2])
	} else if rem == 2 {
		check(&nums[2], &nums[1])
	}

	xx := make([]int, 0, n)

	for i := 0; i < 3; i++ {
		for j := 0; j < len(nums[i]); j++ {
			xx = append(xx, nums[i][j])
		}
	}

	sort.Ints(xx)

	var buf bytes.Buffer

	for i := len(xx) - 1; i >= 0; i-- {
		buf.WriteByte(byte(xx[i] + '0'))
	}

	if buf.Len() == 0 {
		return ""
	}

	s := buf.String()

	var i int

	for i < len(s) && s[i] == '0' {
		i++
	}

	if i == len(s) {
		return "0"
	}

	return buf.String()
}
