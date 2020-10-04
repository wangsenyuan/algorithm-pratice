package p1518

const MOD = 1000000007

func numSub(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			continue
		}

		var j = i
		for j < len(s) && s[j] == '1' {
			j++
		}

		cnt := j - i

		res += cnt * (cnt + 1) / 2
		res %= MOD

		i = j - 1
	}

	return res
}
