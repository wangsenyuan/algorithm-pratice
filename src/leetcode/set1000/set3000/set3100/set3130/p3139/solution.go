package p3139

const mod = 1e9 + 7

func minCostToEqualizeArray(nums []int, cost1 int, cost2 int) int {
	x := max_of(nums)
	var res int
	for _, num := range nums {
		res += (x - num) * cost1
	}
	n := len(nums)
	if 2*cost1 <= cost2 {
		return res % mod
	}

	mx := 0
	var sum int
	for _, num := range nums {
		mx = max(mx, x-num)
		sum += x - num
	}

	for add := 0; ; add++ {
		s := sum + add*n
		xm := mx + add
		if xm > s-xm {
			res = min(res, (s-xm)*cost2+(2*xm-s)*cost1)
		} else {
			res = min(res, s/2*cost2+(s&1)*cost1)
			if n%2 == 0 || s&1 == 0 {
				break
			}
		}
		if n <= 2 {
			break
		}
	}

	return res % mod
}

func max_of(nums []int) int {
	res := nums[0]
	for _, num := range nums {
		res = max(res, num)
	}
	return res
}
