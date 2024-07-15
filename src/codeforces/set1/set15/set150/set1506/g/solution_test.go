package main

import "testing"

func runSample(t *testing.T, s string, expect string) {
	res := solve(s)

	if res != expect {
		t.Fatalf("Sample expect %s, but got %s", expect, res)
	}
}

func TestSample1(t *testing.T) {
	s := "codeforces"
	expect := "odfrces"
	runSample(t, s, expect)
}

func TestSample2(t *testing.T) {
	s := "aezakmi"
	expect := "ezakmi"
	runSample(t, s, expect)
}

func TestSample3(t *testing.T) {
	s := "abacaba"
	expect := "cba"
	runSample(t, s, expect)
}

func TestSample4(t *testing.T) {
	s := "convexhull"
	expect := "convexhul"
	runSample(t, s, expect)
}

func TestSample5(t *testing.T) {
	s := "swflldjgpaxs"
	expect := "wfldjgpaxs"
	runSample(t, s, expect)
}
func TestSample6(t *testing.T) {
	s := "myneeocktxpqjpz"
	expect := "myneocktxqjpz"
	runSample(t, s, expect)
}


func TestSample7(t *testing.T) {
	s := "cemznbgdzzbddd"
	expect := "cemzngdb"
	runSample(t, s, expect)
}
