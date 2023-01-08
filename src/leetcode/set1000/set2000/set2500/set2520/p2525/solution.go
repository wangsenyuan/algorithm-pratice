package p2525

func categorizeBox(length int, width int, height int, mass int) string {
	L := int64(length)
	W := int64(width)
	H := int64(height)
	M := int64(mass)

	var x string

	if L >= 1e4 || W >= 1e4 || H >= 1e4 || L*W*H >= 1e9 {
		x = "Bulky"
	}

	var y string

	if M >= 100 {
		y = "Heavy"
	}

	if x == "Bulky" && y == "Heavy" {
		return "Both"
	}

	if x != "Bulky" && y != "Heavy" {
		return "Neither"
	}

	if len(x) > 0 {
		return x
	}
	return y
}
