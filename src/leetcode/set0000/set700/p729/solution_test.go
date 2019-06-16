package p729

import "testing"

func TestSample1(t *testing.T) {
	calendar := Constructor()
	res := calendar.Book(10, 20)
	if res != true {
		t.Errorf("book 10 - 20 should success, but fail")
	}
	res = calendar.Book(15, 25)
	if res {
		t.Errorf("book 15 - 25 should fail, but success")
	}

	res = calendar.Book(20, 30)
	if !res {
		t.Errorf("book 20 - 30 should success, but fail")
	}
}

func TestSample2(t *testing.T) {
	calendar := Constructor()
	calendar.Book(7997, 9496)
	calendar.Book(4833, 6071)
	calendar.Book(9589, 10000)
	calendar.Book(7483, 8952)
}
