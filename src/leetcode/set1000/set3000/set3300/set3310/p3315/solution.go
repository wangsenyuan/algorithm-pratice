package p3315

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, num := range nums {
		ans[i] = find(num)
	}

	return ans
}

func find(num int) int {
	if num == 2 {
		return -1
	}
	// a | (a + 1) = num 且num是质数
	// 其他的都是奇数, 3, 5, 7, ...
	// num - 1 是一个答案
	// 1001
	// 7 ， 101 + 1 =	> 110
	var p int
	for tmp := num; tmp&1 == 1; tmp >>= 1 {
		p++
	}
	// num & (1 << p) == 0
	p--
	return num ^ (1 << p)
}
