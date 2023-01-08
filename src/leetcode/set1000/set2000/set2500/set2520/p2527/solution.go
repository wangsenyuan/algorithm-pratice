package p2527

const H = 30

func xorBeauty(nums []int) int {
	//	n := len(nums)
	// 可以按位处理
	var res int
	for i := 0; i < H; i++ {
		var cnt int
		for _, num := range nums {
			cnt += (num >> i) & 1
		}
		if cnt > 0 {
			// 选择 cnt其中一个作为k, 选择cnt中任一个做为i, 任何其他一个作为j
			// 共有 cnt * cnt * n
			all := int64(cnt) * int64(cnt) * int64(cnt)
			//all += int64(cnt) * int64(cnt) * int64(n-cnt) * 2
			if all&1 == 1 {
				res |= 1 << i
			}
		}
	}

	return res
}
