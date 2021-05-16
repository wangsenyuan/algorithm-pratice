package p1861

func memLeak(memory1 int, memory2 int) []int {
	for i := 1; ; i++ {
		// need i bits
		if memory1 < memory2 {
			// use 2
			if memory2 < i {
				return []int{i, memory1, memory2}
			}
			memory2 -= i
		} else {
			if memory1 < i {
				return []int{i, memory1, memory2}
			}
			memory1 -= i
		}
	}
}
