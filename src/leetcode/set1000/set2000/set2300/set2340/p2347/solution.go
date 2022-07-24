package p2347

import "sort"

func bestHand(ranks []int, suits []byte) string {
	if flush(ranks, suits) {
		return "Flush"
	}
	if threeOfKind(ranks, suits) {
		return "Three of a Kind"
	}

	if pair(ranks, suits) {
		return "Pair"
	}
	return "High Card"
}

func flush(ranks []int, suites []byte) bool {
	for i := 1; i < 5; i++ {
		if suites[i] != suites[0] {
			return false
		}
	}
	return true
}

func threeOfKind(ranks []int, suites []byte) bool {
	arr := make([]int, 5)
	copy(arr, ranks)
	sort.Ints(arr)
	for i := 0; i+3 <= 5; i++ {
		if arr[i] == arr[i+1] && arr[i] == arr[i+2] {
			return true
		}
	}
	return false
}

func pair(ranks []int, suites []byte) bool {
	arr := make([]int, 5)
	copy(arr, ranks)
	sort.Ints(arr)
	for i := 0; i+2 <= 5; i++ {
		if arr[i] == arr[i+1] {
			return true
		}
	}
	return false
}
