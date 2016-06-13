package recursive

func Permutations(str string) []string {
	if len(str) == 0 {
		return []string{""}
	}

	if len(str) == 1 {
		return []string{str}
	}

	cs := make([]string, 0, 100)

	for mid := 0; mid < len(str); mid++ {
		s := []rune(str)
		x := string(s[mid])
		as := copySliceBut(s, mid)
		ps := Permutations(string(as))
		for _, p := range ps {
			cs = append(cs, x + p)
			//cs = append(cs, p + x)
		}
	}
	return cs
}

func copySliceBut(slice []rune, i int) []rune {
	rs := make([]rune, i, len(slice) - 1)
	copy(rs, slice[0:i])
	return append(rs, slice[i+1:]...)
}
