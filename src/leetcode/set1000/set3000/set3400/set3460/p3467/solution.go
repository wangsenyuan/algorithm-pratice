package p3467

func transformArray(nums []int) []int {
	cnt := make([]int, 2)
	for _, num := range nums {
		cnt[num&1]++
	}
	n := len(nums)
	res := make([]int, n)
	for i := cnt[0]; i < n; i++ {
		res[i] = 1
	}
	return res
}
