package main

import "testing"

func isPalindrome(str string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func runSample(t *testing.T, sample string, expect string, expectCnt int, expectCan bool) {
	res, cnt, can := solve(sample)
	if can != expectCan {
		t.Errorf("sample %s could be lucky? %v, but get %v", sample, expectCan, can)
		return
	}
	if !expectCan {
		return
	}

	if cnt != expectCnt {
		t.Errorf("sample %s could be lucky after %d changes, but got %d", sample, expectCnt, cnt)
		return
	}

	if res != expect {
		t.Errorf("sample %s could be lucky %s, but get %s\n", sample, expect, res)
	}

	if !isPalindrome(res) {
		t.Errorf("sample %s could be luck %s, get %s, which is not palindrome", sample, expect, res)
		return
	}

	diff := 0
	for i := 0; i < len(res); i++ {
		if sample[i] != res[i] {
			diff++
		}
	}

	if diff != cnt {
		t.Errorf("sample %s change to %s should use %d changes, but actually is %d", sample, res, cnt, diff)
	}
}

func TestSample1(t *testing.T) {
	runSample(t, "laubcdkey", "luckykcul", 8, true)
}

func TestSample2(t *testing.T) {
	runSample(t, "luckycodechef", "luckycocykcul", 6, true)
}

func TestSample3(t *testing.T) {
	runSample(t, "aaaaaaaa", "", -1, false)
}

func TestSample4(t *testing.T) {
	runSample(t, "aluckycodechefb", "aluckycocykcula", 7, true)
}

func TestSample5(t *testing.T) {
	runSample(t, "aluckycoodechefb", "aluckycoocykcula", 7, true)
}
