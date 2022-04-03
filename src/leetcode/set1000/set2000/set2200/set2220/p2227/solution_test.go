package p2227

import "testing"

func TestSample1(t *testing.T) {
	keys := []byte("abcd")
	values := []string{"ei", "zf", "ei", "am"}
	dict := []string{"abcd", "acbd", "adbc", "badc", "dacb", "cadb", "cbda", "abad"}
	encrypter := Constructor(keys, values, dict)

	res := encrypter.Encrypt("abcd")

	if res != "eizfeiam" {
		t.Fatalf("Sample expect eizfeiam, but got %s", res)
	}

	cnt := encrypter.Decrypt("eizfeiam")

	if cnt != 2 {
		t.Fatalf("Sample expect 2, but got %d", cnt)
	}
}
