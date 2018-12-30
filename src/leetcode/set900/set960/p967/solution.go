package p967

func numsSameConsecDiff(N int, K int) []int {

	var loop func(i int, prev int, num int)
	res := make([]int, 0, 10)

	loop = func(i int, prev int, num int) {
		if i == N {
			res = append(res, num)
			return
		}
		if K > 0 {
			if prev-K >= 0 {
				loop(i+1, prev-K, num*10+prev-K)
			}
			if prev+K <= 9 {
				loop(i+1, prev+K, num*10+prev+K)
			}
		} else {
			loop(i+1, prev, num*10+prev)
		}
	}
	var start = 1
	if N == 1 {
		start = 0
	}
	for i := start; i <= 9; i++ {
		loop(1, i, i)
	}
	return res
}
