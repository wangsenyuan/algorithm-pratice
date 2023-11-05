package p187

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return nil
	}
	records := make(map[int64]int)

	n := len(s)
	var num int64
	var base int64 = 1

	getValue := func(x byte) int64 {
		if x == 'A' {
			return 0
		}
		if x == 'T' {
			return 1
		}
		if x == 'C' {
			return 2
		}
		return 3
	}

	for i := 0; i < 10; i++ {
		// ATCG
		num = num*4 + getValue(s[i])
		base *= 4
	}
	base /= 4
	var ans []string
	records[num]++
	for i := 1; i+10 <= n; i++ {
		num -= getValue(s[i-1]) * base
		num = num*4 + getValue(s[i+9])
		if records[num] == 1 {
			ans = append(ans, s[i:i+10])
		}
		records[num]++
	}

	return ans
}
