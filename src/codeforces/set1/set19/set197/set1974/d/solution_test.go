package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res == expect {
		return
	}
	if expect == "NO" || res == "NO" {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}

	pos := make([]int, 4)
	cnt := []int{0, 0}
	for i := 0; i < len(s); i++ {
		j := id(s[i])
		v := vals[j]
		if res[i] == 'R' {
			pos[0] += v[0]
			pos[1] += v[1]
			cnt[0]++
		} else {
			pos[2] += v[0]
			pos[3] += v[1]
			cnt[1]++
		}
	}

	if pos[0] != pos[2] || pos[1] != pos[3] || cnt[0] == 0 || cnt[1] == 0 {
		t.Fatalf("Sample result %s, not correct, getting %v", res, pos)
	}
}

func TestSample1(t *testing.T) {
	s := "NENSNE"
	expect := "RRHRRH"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "WWW"
	expect := "NO"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "NESSWS"
	expect := "HRRHRH"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "SN"
	expect := "NO"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "SN"
	expect := "NO"
	runSample(t, s, expect)
}

func TestSample6(t *testing.T) {
	s := "WE"
	expect := "NO"
	runSample(t, s, expect)
}

func TestSample7(t *testing.T) {
	s := "SSNN"
	expect := "RHRH"
	runSample(t, s, expect)
}

func TestSample8(t *testing.T) {
	s := "WESN"
	expect := "RRHH"
	runSample(t, s, expect)
}

func TestSample9(t *testing.T) {
	s := "SS"
	expect := "RH"
	runSample(t, s, expect)
}

func TestSample10(t *testing.T) {
	s := "EWNN"
	expect := "RRRH"
	runSample(t, s, expect)
}

func TestSample11(t *testing.T) {
	s := "WEWE"
	expect := "RRHH"
	runSample(t, s, expect)
}
