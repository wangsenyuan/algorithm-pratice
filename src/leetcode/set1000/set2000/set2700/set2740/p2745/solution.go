package p2745

func maximumNumberOfStringPairs(words []string) int {
	cnt := make(map[int]int)
	for _, w := range words {
		cnt[index(w)]++
	}

	var res int
	for k := range cnt {
		var r int
		tmp := k
		for tmp > 0 {
			r = r*26 + tmp%26
			tmp /= 26
		}
		if r != k && cnt[r] > 0 {
			res++
			delete(cnt, r)
		}
	}

	return res
}

func index(w string) int {
	var res int
	for i := 0; i < len(w); i++ {
		res = res*26 + int(w[i]-'a')
	}
	return res
}
