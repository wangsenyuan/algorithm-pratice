package p930

func numSubarraysWithSum(A []int, S int) int {
	n := len(A)
	f := make([]int, n+1)
	f[0] = 1
	var ans int
	var sum int
	for i := 0; i < n; i++ {
		sum += A[i]
		prev := sum - S
		if prev >= 0 {
			ans += f[prev]
		}
		f[sum]++
	}
	return ans
}
