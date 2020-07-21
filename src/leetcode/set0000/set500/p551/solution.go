package p551

func checkRecord(s string) bool {
	aCount := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			aCount++
		}
		if aCount > 1 {
			return false
		}

		if s[i] == 'L' && i > 1 && s[i-1] == 'L' && s[i-2] == 'L' {
			return false
		}
	}

	return true
}
