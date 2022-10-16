package p2441

func countDistinctIntegers(nums []int) int {
	cnt := make(map[int64]bool)

	for _, num := range nums {
		cnt[int64(num)] = true
		cnt[reverse(int64(num))] = true
	}

	return len(cnt)
}

func reverse(num int64) int64 {
	var res int64

	for num > 0 {
		x := num % 10
		res = res*10 + x
		num /= 10
	}
	return res
}
