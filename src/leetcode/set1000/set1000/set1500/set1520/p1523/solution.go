package p1523

const MOD = 1000000007

func numOfSubarrays(arr []int) int {
	cnt := make([]int, 2)

	var sum int
	cnt[0] = 1
	var res int
	for i := 0; i < len(arr); i++ {
		//if arr[i]&1 == 1 {
		//	res++
		//}
		sum += arr[i]
		sum = sum & 1

		if sum == 1 {
			res += cnt[0]
			cnt[1]++
		} else {
			res += cnt[1]
			cnt[0]++
		}
		res %= MOD
	}

	return res
}
