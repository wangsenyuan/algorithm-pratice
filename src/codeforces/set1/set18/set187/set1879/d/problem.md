You are given an array 𝑎
of length 𝑛
consisting of non-negative integers.

You have to calculate the value of ∑𝑛𝑙=1∑𝑛𝑟=𝑙𝑓(𝑙,𝑟)⋅(𝑟−𝑙+1)
, where 𝑓(𝑙,𝑟)
is 𝑎𝑙⊕𝑎𝑙+1⊕⋯⊕𝑎𝑟−1⊕𝑎𝑟
(the character ⊕
denotes bitwise XOR).

Since the answer can be very large, print it modulo 998244353
.

### thoughts

1. f(l, r) = a[l] ^ ... ^ a[r] = x[r] ^ x[l-1]
2. 按位考虑
3. 对于第d位，给定l时，
4. 假设 x[l...r] = 1，那么它的贡献就是 1 << d * (r - l + 1)
5. 所以，这里要找到那些等于1的位置
6. 1001011,
7. 1 << d是固定的，考虑r，x[l-1]是确定的，就是找和x[l-1]不等的位置的sum和cnt