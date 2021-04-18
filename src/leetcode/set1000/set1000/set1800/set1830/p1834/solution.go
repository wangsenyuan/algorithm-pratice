package p1834

func getXORSum(arr1 []int, arr2 []int) int {
	var xor int
	for _, num := range arr1 {
		xor ^= num
	}
	var res int
	for _, num := range arr2 {
		tmp := xor & num
		res ^= tmp
	}

	return res
}
func getXORSum1(arr1 []int, arr2 []int) int {

	cnt := make([]int, 30)

	for _, x := range arr2 {
		for i := 0; i < 30; i++ {
			if (x>>i)&1 == 1 {
				cnt[i]++
			}
		}
	}
	var ans int

	for _, num := range arr1 {
		var tmp int
		for i := 0; i < 30; i++ {
			a := (num >> i) & 1
			if a == 1 && cnt[i]&1 == 1 {
				tmp |= 1 << i
			}
		}
		ans ^= tmp
	}

	return ans
}
