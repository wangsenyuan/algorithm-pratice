package p2601

const N = 1001

func primeSubOperation(nums []int) bool {
	set := make([]bool, N)
	var primes []int
	for x := 2; x < N; x++ {
		if !set[x] {
			primes = append(primes, x)
			for y := x * x; y < N; y += x {
				set[y] = true
			}
		}
	}

	n := len(nums)
	cur := nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < cur {
			cur = nums[i]
			continue
		}
		j := search(len(primes), func(j int) bool {
			return nums[i]-primes[j] < cur
		})
		if j == len(primes) {
			return false
		}
		// nums[i] - primes[j] <= cur
		cur = nums[i] - primes[j]
		if cur <= 0 {
			return false
		}
	}
	return true
}

func search(n int, f func(int) bool) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return r
}
