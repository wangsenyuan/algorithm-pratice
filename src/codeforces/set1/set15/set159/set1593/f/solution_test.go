package main

import "testing"

func runSample(t *testing.T, nums string, a, b int, expect string) {
	res := solve(nums, a, b)

	if len(expect) == 0 {
		if len(res) > 0 {
			t.Fatalf("Sample expect %s, but got %s", expect, res)
		}
		return
	}

	if len(res) == 0 {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	check := func(ans string) int {
		cnt := []int{0, 0}
		r := []int{0, 0}
		for i := 0; i < len(nums); i++ {
			num := int(nums[i] - '0')
			if ans[i] == 'R' {
				r[0] = (r[0]*10 + num) % a
				cnt[0]++
			} else {
				r[1] = (r[1]*10 + num) % b
				cnt[1]++
			}
		}
		if r[0] != 0 || r[1] != 0 {
			return -1
		}
		return abs(cnt[0] - cnt[1])
	}

	x := check(expect)
	y := check(res)

	if x != y {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	num := "02165"
	a, b := 3, 13
	expect := "RBRBR"
	runSample(t, num, a, b, expect)
}

func TestSample2(t *testing.T) {
	num := "1357"
	a, b := 2, 1
	expect := ""
	runSample(t, num, a, b, expect)
}

func TestSample3(t *testing.T) {
	num := "12345678"
	a, b := 1, 1
	expect := "BBRRRRBB"
	runSample(t, num, a, b, expect)
}

func TestSample4(t *testing.T) {
	num := "90"
	a, b := 7, 9
	expect := "BR"
	runSample(t, num, a, b, expect)
}

func TestSample5(t *testing.T) {
	num := "1284326696960955811491101"
	a, b := 13, 11
	expect := "RRRRRRRRRRRBRBBBBBBBBBBBB"
	runSample(t, num, a, b, expect)
}
