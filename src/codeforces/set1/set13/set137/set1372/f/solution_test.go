package main

import (
	"reflect"
	"testing"
)

func runSample(t *testing.T, arr []int) {
	ask := func(l, r int) (int, int) {
		l--
		r--
		cnt := make(map[int]int)
		var mx int
		for i := l; i <= r; i++ {
			cnt[arr[i]]++
			if cnt[arr[i]] > mx {
				mx = cnt[arr[i]]
			}
		}

		for i := l; i <= r; i++ {
			if cnt[arr[i]] == mx {
				return arr[i], cnt[arr[i]]
			}
		}
		return -1, 0
	}

	res := solve(len(arr), ask)

	if !reflect.DeepEqual(res, arr) {
		t.Errorf("Sample expect %v, but got %v", arr, res)
	}
}

func TestSample1(t *testing.T) {
	arr := []int{1, 2, 2, 3, 3, 4}
	runSample(t, arr)
}
