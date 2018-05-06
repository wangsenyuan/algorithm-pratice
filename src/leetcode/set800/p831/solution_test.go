package p831

import "testing"

func TestMaskEmail(t *testing.T) {
	S := "JackAndJill@Gmail.Com"
	expect := "j*****l@gmail.com"
	res := maskPII(S)
	if res != expect {
		t.Errorf("error got %s", res)
	}
}

func TestMaskPhone(t *testing.T) {
	S := "+111 111 111 1111"
	expect := "+***-***-***-1111"
	res := maskPII(S)
	if res != expect {
		t.Errorf("error got %s", res)
	}
}

func TestMaskPhone2(t *testing.T) {
	S := "1(234)567-890"
	expect := "***-***-7890"
	res := maskPII(S)
	if res != expect {
		t.Errorf("error got %s", res)
	}
}
