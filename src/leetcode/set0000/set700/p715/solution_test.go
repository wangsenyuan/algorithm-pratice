package p715

import "testing"

func TestSample1(t *testing.T) {
	target := Constructor()
	target.AddRange(10, 20)
	target.RemoveRange(14, 16)
	ok := target.QueryRange(10, 14)

	if !ok {
		t.Fatalf("range [10, 14) should be there")
	}

	ok = target.QueryRange(13, 15)

	if ok {
		t.Fatalf("range [13, 15) should not be there")
	}

	ok = target.QueryRange(16, 17)

	if !ok {
		t.Fatalf("range [16, 17) should be there")
	}
}

func TestSample2(t *testing.T) {
	target := Constructor()
	target.AddRange(10, 180)
	target.AddRange(150, 200)
	target.AddRange(250, 500)
	ok := target.QueryRange(50, 100)

	if !ok {
		t.Fatalf("range [50,100) should be there")
	}

	ok = target.QueryRange(180, 300)

	if ok {
		t.Fatalf("range [180,300) should not be there")
	}

	ok = target.QueryRange(600, 1000)

	if ok {
		t.Fatalf("range [600,1000) should not be there")
	}

	target.RemoveRange(50, 150)
	ok = target.QueryRange(50, 100)

	if ok {
		t.Fatalf("range [50,100) should not be there")
	}
}
