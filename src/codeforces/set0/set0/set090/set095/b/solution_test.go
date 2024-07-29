package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	cnt := make([]int, 2)
	for i := 0; i < len(res); i++ {
		if res[i] == '4' {
			cnt[0]++
		} else if res[i] == '7' {
			cnt[1]++
		} else {
			t.Fatalf("Sample result %s, not correct", res)
		}
	}
	if cnt[0] != cnt[1] {
		t.Fatalf("Sample result %s, not correct", res)
	}

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

}

func TestSample1(t *testing.T) {
	s := "4500"
	expect := "4747"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "47"
	expect := "47"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "4747"
	expect := "4747"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "4847"
	expect := "7447"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "7754"
	expect := "444777"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "47447774444477747744744477747744477774777774747474477744474447744447747777744777444474777477447777747477474774477444777777744774777474477744444474744777774744447747477747474447444444447444774744777447"
	expect := "47447774444477747744744477747744477774777774747474477744474447744447747777744777444474777477447777747477474774477444777777744774777474477744444474744777774744447747477747474447444444447444774747444444"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "774447"
	expect := "774447"
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "744447"
	expect := "744477"
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "99999999"
	expect := "4444477777"
	runSample(t, s, expect)
}
