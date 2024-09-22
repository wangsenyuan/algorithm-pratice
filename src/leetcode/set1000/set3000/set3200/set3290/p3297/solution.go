package p3297

func validSubstringCount(word1 string, word2 string) int64 {
	cnt2 := make([]int32, 26)

	for _, x := range []byte(word2) {
		cnt2[int(x-'a')]++
	}
	type pair struct {
		first  int32
		second int32
	}

	var arr []pair

	cnt1 := make([]int32, 26)

	var expect int32
	for i := 0; i < 26; i++ {
		if cnt2[i] > 0 {
			expect |= 1 << i
		}
	}

	var flag int32

	add := func(v byte) {
		i := int(v - 'a')
		cnt1[i]++
		if cnt1[i] == cnt2[i] {
			flag |= 1 << i
		}
	}

	rem := func(v byte) {
		i := int(v - 'a')
		cnt1[i]--
		if cnt1[i] == cnt2[i]-1 {
			flag ^= 1 << i
		}
	}
	n := len(word1)
	for l, r := 0, 0; r < n; r++ {
		add(word1[r])
		for l <= r && flag&expect == expect {
			rem(word1[l])
			if flag&expect != expect {
				// make sure s[l...r] is good
				add(word1[l])
				arr = append(arr, pair{int32(l), int32(r)})
				break
			}
			l++
		}
	}

	if len(arr) == 0 {
		return 0
	}

	var res int
	l := arr[0].first
	r := arr[0].second

	for i := 1; i < len(arr); i++ {
		if arr[i].first == l {
			// 这是一个更长的good的序列
			continue
		}
		// arr[i].first < l
		nr := arr[i].second
		res += int(l+1) * int(nr-r)
		l = arr[i].first
		r = nr
	}

	res += int(l+1) * (n - int(r))

	return int64(res)
}
