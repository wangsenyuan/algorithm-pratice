package p3306

const voewls = "aeiou"

func countOfSubstrings(word string, k int) int64 {

	is := make([]bool, 26)
	var T int

	for _, x := range []byte(voewls) {
		is[int(x-'a')] = true
		T |= 1 << (x - 'a')
	}

	n := len(word)
	// 必须知道任何一个位置前面(后面)的辅音的位置
	dp := make([]int, n)
	prev := -1

	for i, x := range []byte(word) {
		dp[i] = prev
		v := int(x - 'a')
		if !is[v] {
			prev = i
		}
	}

	freq := make([]int, 26)

	var flag int
	var cnt int

	add := func(x byte) {
		i := int(x - 'a')
		freq[i]++
		if freq[i] == 1 {
			flag |= 1 << i
		}
		if !is[i] {
			cnt++
		}
	}

	rem := func(x byte) {
		i := int(x - 'a')
		freq[i]--
		if freq[i] == 0 {
			flag ^= 1 << i
		}
		if !is[i] {
			cnt--
		}
	}

	var res int
	for l, r := 0, 0; r < n; r++ {
		add(word[r])
		for cnt >= k && flag&T == T {
			rem(word[l])
			if cnt < k || flag&T != T {
				add(word[l])
				break
			}
			l++
		}

		if cnt == k && flag&T == T {
			res += (l - dp[l])
		}

	}
	return int64(res)
}
