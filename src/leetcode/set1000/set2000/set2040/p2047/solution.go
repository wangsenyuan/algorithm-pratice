package p2047

func countValidWords(sentence string) int {

	var res int

	var j int
	for i := 0; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == ' ' {
			res += check(sentence[j:i])
			j = i + 1
		}
	}

	return res
}

func check(s string) int {
	if len(s) == 0 {
		return 0
	}
	var cnt int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] >= '0' && s[i] <= '9' {
			return 0
		}
		if s[i] == '-' {
			if cnt > 0 {
				return 0
			}
			if i == 0 || s[i-1] < 'a' || s[i-1] > 'z' || i+1 == len(s) || s[i+1] < 'a' || s[i+1] > 'z' {
				return 0
			}
			cnt++
		}
		if s[i] == '!' || s[i] == ',' || s[i] == '.' {
			if i+1 != len(s) {
				return 0
			}
		}
	}
	return 1
}
