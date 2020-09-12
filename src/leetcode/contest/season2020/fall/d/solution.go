package d

const MOD = 1000000007

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	// m == len(cost)
	//best := int64(target) * int64(inc)

	return int(loop(target, make(map[int]int64), inc, dec, jump, cost) % MOD)
}

func loop(cur int, mem map[int]int64, inc, dec int, jump, cost []int) int64 {
	if v, found := mem[cur]; found {
		return v
	}
	ans := int64(cur) * int64(inc)
	for i := 0; i < len(jump); i++ {
		// x * jump[i] <= cur
		x := cur / jump[i]
		if x > 0 {
			tmp := loop(x, mem, inc, dec, jump, cost) + int64(cost[i]) + int64(cur-x*jump[i])*int64(inc)
			if ans > tmp {
				ans = tmp
			}
		}

		y := (cur + jump[i] - 1) / jump[i]
		if y != x && y > 0 && y < cur {
			tmp := loop(y, mem, inc, dec, jump, cost) + int64(cost[i]) + int64(y*jump[i]-cur)*int64(dec)
			if ans > tmp {
				ans = tmp
			}
		}
	}

	// x * jump[i] > cur
	mem[cur] = ans
	return ans
}
