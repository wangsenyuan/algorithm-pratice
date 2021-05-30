package p1882

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return getAsNum(firstWord)+getAsNum(secondWord) == getAsNum(targetWord)
}

func getAsNum(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		res = res*10 + int(s[i]-'a')
	}
	return res
}
