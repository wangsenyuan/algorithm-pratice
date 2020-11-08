package p5563

import "sort"

const MOD = 1000000007

func maxProfit(inventory []int, orders int) int {
	sort.Ints(inventory)

	n := len(inventory)
	var res int
	var cnt int

	for i := n - 1; i >= 0 && cnt < orders; {
		j := i
		for i >= 0 && inventory[i] == inventory[j] {
			i--
		}
		left := orders - cnt
		var m = n - 1 - i
		cur := inventory[j]
		if i >= 0 {
			cur -= inventory[i]
		}
		a, b := left/m, left%m
		if cur <= a {
			a = cur
			b = 0
		}
		tmp := int((int64(inventory[j]+inventory[j]-a+1) * int64(a) / 2) % MOD)
		mul(&tmp, m)
		add(&res, tmp)
		tmp = inventory[j] - a
		mul(&tmp, b)
		add(&res, tmp)
		cnt += m*a + b
	}

	return res
}

func mul(a *int, b int) {
	res := int64(*a) * int64(b) % MOD
	*a = int(res)
}

func add(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}
