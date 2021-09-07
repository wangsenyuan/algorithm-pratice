package p1994

// 2 * 7 * 3 * 5 * 29 * 23 * 29

const M = 1000000007

var primes = []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

var all = []int{1, 2, 3, 5, 6, 7, 10, 11, 13, 14, 15, 17, 19, 21, 22, 23, 26, 29, 30}

func numberOfGoodSubsets(nums []int) int {

	good := make([]bool, 31)

	for _, num := range all {
		good[num] = true
	}

	pf := make([]int, 31)

	for x := 2; x < 31; x++ {
		if pf[x] == 0 {
			for y := x; y < 31; y += x {
				if pf[y] == 0 {
					pf[y] = x
				}
			}
		}
	}

	pos := make([]int, 31)
	for i, p := range primes {
		pos[p] = i
	}

	cnt := make(map[int]int)
	var cnt1 int
	for _, num := range nums {
		if good[num] && num > 1 {
			cnt[num]++
		}
		if num == 1 {
			cnt1++
		}
	}
	arr := make([]Pair, 0, len(cnt))
	for k, v := range cnt {
		arr = append(arr, Pair{k, v})
	}

	n := len(arr)
	TOT := 1 << n

	var res int64

	cnt2 := make([]int, len(primes))

	for state := 1; state < TOT; state++ {
		// we want to get dp[state]
		var product int64 = 1
		ok := true
		for i := 0; i < n && ok; i++ {
			if (state>>i)&1 == 1 {
				num := arr[i].first
				for num > 1 && ok {
					cnt2[pos[pf[num]]]++
					if cnt2[pos[pf[num]]] > 1 {
						ok = false
					}
					num /= pf[num]
				}
				product *= int64(arr[i].second)
				product %= M
			}
		}

		for i := 0; i < len(primes); i++ {
			cnt2[i] = 0
		}

		if ok {
			res += product
			res %= M
		}
	}

	res = pow(2, cnt1) * res % M

	return int(res)
}

type Pair struct {
	first, second int
}

func numberOfGoodSubsets1(nums []int) int {

	cnt := make([]int, 31)
	for _, num := range nums {
		if num%4 == 0 || num%9 == 0 || num == 25 {
			continue
		}
		cnt[num]++
	}
	var ret int64

	for i := 2; i < 1<<19; i++ {
		valid := true
		if i&(1<<1) > 0 {
			// 2
			if (i&(1<<4)) > 0 || (i&(1<<6)) > 0 || (i&(1<<9)) > 0 || (i&(1<<14)) > 0 || (i&(1<<16)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<2) > 0 {
			// 3
			if (i&(1<<4)) > 0 || (i&(1<<10)) > 0 || (i&(1<<13)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<3) > 0 {
			// 5
			if (i&(1<<6)) > 0 || (i&(1<<10)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<4) > 0 {
			// 6
			if (i&(1<<6)) > 0 || (i&(1<<9)) > 0 || (i&(1<<10)) > 0 || (i&(1<<13)) > 0 || (i&(1<<14)) > 0 || (i&(1<<16)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<5) > 0 {
			// 7
			if (i&(1<<9)) > 0 || (i&(1<<13)) > 0 {
				valid = false
			}
		}
		if i&(1<<6) > 0 {
			// 10
			if (i&(1<<9)) > 0 || (i&(1<<10)) > 0 || (i&(1<<14)) > 0 || (i&(1<<16)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<7) > 0 {
			// 11
			if i&(1<<14) > 0 {
				valid = false
			}
		}
		if i&(1<<8) > 0 {
			// 13
			if i&(1<<16) > 0 {
				valid = false
			}
		}
		if i&(1<<9) > 0 {
			// 14
			if (i&(1<<13)) > 0 || (i&(1<<14)) > 0 || (i&(1<<16)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<10) > 0 {
			// 15
			if (i&(1<<13)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<13) > 0 {
			// 21
			if i&(1<<18) > 0 {
				valid = false
			}
		}
		if i&(1<<14) > 0 {
			// 22
			if (i&(1<<16)) > 0 || (i&(1<<18)) > 0 {
				valid = false
			}
		}
		if i&(1<<16) > 0 {
			// 26
			if i&(1<<18) > 0 {
				valid = false
			}
		}

		if !valid {
			continue
		}

		var ans int64 = 1
		for j := 0; j < 19; j++ {
			if i&(1<<j) > 0 {
				if j == 0 {
					ans = ans * ((pow(2, cnt[all[j]]) - 1 + M) % M) % M
				} else {
					ans = ans * int64(cnt[all[j]]) % M
				}
			}
		}
		ret = (ret + ans) % M
	}
	return int(ret)
}

func pow(a, b int) int64 {
	A := int64(a)
	R := int64(1)
	for b > 0 {
		if b&1 == 1 {
			R = (R * A) % M
		}
		A = (A * A) % M
		b >>= 1
	}
	return R
}
