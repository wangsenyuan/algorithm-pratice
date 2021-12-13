package p2035

import (
	"math"
	"sort"
)

func minimumDifference(nums []int) int {
	N := len(nums)

	if N < 15 {
		return solve1(nums)
	}

	n := N / 2

	cnt := make([]map[int]bool, n+1)

	for state := 0; state < (1 << n); state++ {
		dc := digitCount(state)
		if cnt[dc] == nil {
			cnt[dc] = make(map[int]bool)
		}
		var sum int
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				sum += nums[i]
			} else {
				sum -= nums[i]
			}
		}
		cnt[dc][sum] = true
	}

	arr := make([][]int, n+1)

	for i := 0; i <= n; i++ {
		arr[i] = make([]int, 0, len(cnt[i]))
		for k := range cnt[i] {
			arr[i] = append(arr[i], k)
		}
		sort.Ints(arr[i])
	}

	best := math.MaxInt32

	for state := 0; state < (1 << n); state++ {
		dc := digitCount(state)
		cd := n - dc
		var sum int
		for i := 0; i < n; i++ {
			if (state>>i)&1 == 1 {
				sum += nums[i+n]
			} else {
				sum -= nums[i+n]
			}
		}
		// find the close number to -sum in ar
		i := sort.SearchInts(arr[cd], -sum)
		if i < len(arr[cd]) && abs(arr[cd][i]+sum) < best {
			best = abs(arr[cd][i] + sum)
			i--
		}
		if i-1 >= 0 && abs(arr[cd][i-1]+sum) < best {
			best = abs(arr[cd][i-1] + sum)
		}
	}

	return best
}

func pick(nums []int, n int, sum int) int {
	// pick n from nums, and put them into arr
	var loop func(i int, j int, cur int) int

	loop = func(i int, j int, cur int) int {

		if i == len(nums) {
			if j == 0 {
				return abs(cur)
			}
			return math.MaxInt32
		}
		// no pick i
		a := loop(i+1, j, cur-nums[i])
		b := loop(i+1, j-1, cur+nums[i])
		return min(a, b)
	}

	return loop(0, n, sum)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func digitCount(num int) int {
	var res int
	for num > 0 {
		res++
		num -= num & -num
	}
	return res
}

func solve1(nums []int) int {
	N := len(nums)
	// N <= 30
	n := N / 2
	// n <= 15
	start := (1 << n) - 1
	end := start << n

	calc := func(mask int) int {
		var ans int
		for i := 0; i < N; i++ {
			if (mask>>i)&1 == 1 {
				ans += nums[i]
			} else {
				ans -= nums[i]
			}
		}
		return abs(ans)
	}

	best := math.MaxInt32

	for mask := start; mask < end; mask = nextNum(mask) {
		tmp := calc(mask)

		if best > tmp {
			best = tmp
		}
		if mask == end-1 {
			break
		}
	}

	return best
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func nextNum(x int) int {
	// right most set bit
	rightOne := x & -x

	// reset the pattern and set next higher bit
	// left part of x will be here
	nextHigherOneBit := x + rightOne

	// nextHigherOneBit is now part [D] of the above explanation.

	// isolate the pattern
	rightOnesPattern := x ^ nextHigherOneBit

	// right adjust pattern
	rightOnesPattern = (rightOnesPattern) / rightOne

	// correction factor
	rightOnesPattern >>= 2

	// rightOnesPattern is now part [A] of the above explanation.

	// integrate new pattern (Add [D] and [A])
	next := nextHigherOneBit | rightOnesPattern

	return next
}
