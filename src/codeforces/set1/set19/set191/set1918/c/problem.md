You are given integers 𝑎
, 𝑏
, 𝑟
. Find the smallest value of |(𝑎⊕𝑥)−(𝑏⊕𝑥)|
 among all 0≤𝑥≤𝑟
.

⊕
 is the operation of bitwise XOR, and |𝑦|
 is absolute value of 𝑦
.

### ideas
1. (a ^ x) - (b ^ x)
2. let u = a ^ x
3. let v = b ^ x
4. 就是要让u和v的高位尽量的一致
5. 假设a和b的第一位不一致的h
6. a[h] = 1, b[h] = 0
7. x[h] = 1, a[h] ^ x[h] = 0, b[h] ^ x[h] = 1
8. 接下来应该就是麻烦的地方了
9. res |= 1 << h
10. 但是，为了让最后的结果变小，这时候就不能取相同的值，而应该取相反的值，使得这个差变小
11. 比如 x[h]取0，假设能够令 x ^ a = 1 << h 但是b ^ x 能够取值 (1 << h)  - 1
12. a + b = a | b + a & b
13. 