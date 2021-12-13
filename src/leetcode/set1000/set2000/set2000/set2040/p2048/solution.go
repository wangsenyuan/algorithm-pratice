package p2048

func nextBeautifulNumber(n int) int {
	arr := digits(n)

	// first try to find that length num > n
	var ans int
	var ans2 int
	cnt := make([]int, 10)

	for mask := 1; mask < 1<<10; mask++ {
		var sum int
		for i := 0; i < 10; i++ {
			if (mask>>i)&1 == 1 {
				sum += i
				cnt[i] = i
			} else {
				cnt[i] = 0
			}
		}
		if sum == len(arr) {
			tmp := find(cnt, arr, n)
			if tmp > n && (ans == 0 || tmp < ans) {
				ans = tmp
			}
		} else if sum == len(arr)+1 {
			var tmp int
			for i := 1; i < 10; i++ {
				for cnt[i] > 0 {
					tmp = tmp*10 + i
					cnt[i]--
				}
			}
			if ans2 == 0 || ans2 > tmp {
				ans2 = tmp
			}
		}
	}
	if ans > n {
		return ans
	}
	return ans2
}

func find(cnt []int, ds []int, num int) int {
	// from high to low
	n := len(ds)

	B := 1
	for i := 0; i < n; i++ {
		B *= 10
	}

	var res int
	var i int
	// find until we can stay same
	for i < n {
		B /= 10
		num %= B
		d := ds[i]
		var ok bool
		if cnt[d] > 0 {
			cnt[d]--
			// if we put d at this place
			// get max num from rest, and we need to make sure it is larger than num
			var tmp int
			for x := 9; x > 0; x-- {
				c := cnt[x]
				for c > 0 {
					tmp = tmp*10 + x
					c--
				}
			}
			cnt[d]++
			ok = tmp > num
		}
		if ok {
			res = res*10 + d
			cnt[d]--
			i++
			continue
		}
		//not ok
		for x := d + 1; x < 10; x++ {
			if cnt[x] > 0 {
				res = res*10 + x
				cnt[x]--
				ok = true
				break
			}
		}
		if ok {
			i++
			break
		}
		return 0
	}

	for x := 1; x < 10; x++ {
		for cnt[x] > 0 {
			res = res*10 + x
			cnt[x]--
		}
	}
	return res
}

func digits(num int) []int {
	var res []int
	for num > 0 {
		res = append(res, num%10)
		num /= 10
	}
	reverse(res)
	return res
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
