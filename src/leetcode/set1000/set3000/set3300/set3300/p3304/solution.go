package p3304

func kthCharacter(k int) byte {
	res := "a"
	for len(res) < k {
		tmp := []byte(res)
		for i := 0; i < len(tmp); i++ {
			x := tmp[i]
			if x != 'z' {
				tmp[i] = x + 1
			} else {
				tmp[i] = 'a'
			}
		}
		res += string(tmp)
	}

	return res[k-1]
}
