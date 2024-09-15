package p3289

func getSneakyNumbers(nums []int) []int {
	n := len(nums)
	cnt := make([]int, n)
	var res []int
	for _, num := range nums {
		cnt[num]++
		if cnt[num] == 2 {
			res = append(res, num)
		}
	}
	return res
}
