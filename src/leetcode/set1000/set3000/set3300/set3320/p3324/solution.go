package p3324

func stringSequence(target string) []string {
	var res []string

	var cur []byte

	for i := 0; i < len(target); i++ {
		// 始终从操作1开始
		var x byte = 'a'
		for x < target[i] {
			res = append(res, string(cur)+string(x))
			x++
		}
		cur = append(cur, x)
		res = append(res, string(cur))
	}

	return res
}
