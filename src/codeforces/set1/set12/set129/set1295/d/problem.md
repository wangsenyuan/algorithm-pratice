You are given two integers 𝑎
 and 𝑚
. Calculate the number of integers 𝑥
 such that 0≤𝑥<𝑚
 and gcd(𝑎,𝑚)=gcd(𝑎+𝑥,𝑚)
.

Note: gcd(𝑎,𝑏)
 is the greatest common divisor of 𝑎
 and 𝑏
.

### ideas
1. gcd(a + b, b) = gcd(a, b)
2. gcd(a + x, m) = gcd(a - m + x, m) = gcd(a, m)
3. x < m
4. gcd(a, m)是可以计算出来的。
5. gcd(a + ?, m) = g
6. gcd(a + x, m) = gcd(a + x, g)
7. 也就是说 a + x 可以整除g
8. 那么这里找g的倍数, g, 2 * g, 3 * g, ... i * g, 
9. 然后看 i * g - a 是否 >= 0 < m
10. i * g >= a => i >= a / g
11. i * g - a < m => i <= (m + a) / g
12. 这里不对。这里会把 2*g, 3 * g, 4 * g 等等，都计算进去
13. m一定能整除2 * g吗？不一定，
14. 找到k * g, m % (k * g) = 0
15. 那么需要从中减去这部分