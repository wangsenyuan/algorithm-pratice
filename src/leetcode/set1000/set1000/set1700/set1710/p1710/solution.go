package p1710

const MOD = 1000000007

func modAdd(a *int, b int) {
	*a += b
	if *a >= MOD {
		*a -= MOD
	}
}

func countPairs(deliciousness []int) int {
	//sort.Ints(deliciousness)
	cnt := make(map[int]int)

	for i := 0; i < len(deliciousness); i++ {
		cnt[deliciousness[i]]++
	}

	var res int

	for k, v := range cnt {
		for j := 0; j <= 21; j++ {
			if 1<<j >= k {
				k1 := (1 << j) - k
				if k1 < k {
					modAdd(&res, cnt[k1]*v%MOD)
				} else if k1 == k {
					modAdd(&res, (v*(v-1)/2)%MOD)
				}
			}
		}
	}

	return res
}
