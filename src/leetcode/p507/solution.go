package main

func checkPerfectNumber(num int) bool {
	var perfectNumbers = []int{6, 28, 496, 8128, 33550336, 8589869056}

	for _, perfect := range perfectNumbers {
		if perfect == num {
			return true
		}
	}

	return false;
}
