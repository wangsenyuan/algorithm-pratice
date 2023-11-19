package p2938

import "math/bits"

func maximumXorProduct(A, B int64, n int) int {
	const mod = 1_000_000_007
	a, b := int(A), int(B)
	if a < b {
		a, b = b, a // 保证 a >= b
	}

	mask := 1<<n - 1
	ax := a &^ mask // 第 n 位及其左边，无法被 x 影响，先算出来
	bx := b &^ mask
	a &= mask // 低于第 n 位，能被 x 影响
	b &= mask

	left := a ^ b      // 可分配：a XOR x 和 b XOR x 一个是 1 另一个是 0
	one := mask ^ left // 无需分配：a XOR x 和 b XOR x 均为 1
	ax |= one          // 先加到异或结果中
	bx |= one

	// 现在要把 left 分配到 ax 和 bx 中
	// 根据基本不等式（均值定理），分配后应当使 ax 和 bx 尽量接近，乘积才能尽量大
	if left > 0 && ax == bx { // a &^ mask = b &^ mask
		// 尽量均匀分配，例如把 1111 分成 1000 和 0111
		highBit := 1 << (bits.Len(uint(left)) - 1)
		ax |= highBit
		left ^= highBit
	}
	// 如果 a &^ mask 更大，则应当全部分给 bx（注意最上面保证了 a>=b）
	bx |= left
	return ax % mod * (bx % mod) % mod
}
