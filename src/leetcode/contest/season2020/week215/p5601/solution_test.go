package p5601

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	os := Constructor(5)

	res := os.Insert(3, "ccccc")
	if len(res) != 0 {
		t.Errorf("Sample expect nil, but got %s", res)
	}
	res = os.Insert(1, "aaaaa")
	if !reflect.DeepEqual(res, []string{"aaaaa"}) {
		t.Errorf("fail at 1")
	}
	res = os.Insert(2, "bbbbb")
	if !reflect.DeepEqual(res, []string{"bbbbb", "ccccc"}) {
		t.Errorf("fail at 2")
	}
	res = os.Insert(5, "eeeee")
	if len(res) != 0 {
		t.Errorf("Sample expect nil, but got %s", res)
	}
	res = os.Insert(4, "ddddd")
	if !reflect.DeepEqual(res, []string{"ddddd", "eeeee"}) {
		t.Errorf("fail at 2")
	}
}
