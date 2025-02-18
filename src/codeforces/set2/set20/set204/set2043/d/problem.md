Given three integers 𝑙
, 𝑟
, and 𝐺
, find two integers 𝐴
 and 𝐵
 (𝑙≤𝐴≤𝐵≤𝑟
) such that their greatest common divisor (GCD) equals 𝐺
 and the distance |𝐴−𝐵|
 is maximized.

If there are multiple such pairs, choose the one where 𝐴
 is minimized. If no such pairs exist, output "-1 -1".

 ### ideas
 1. gcd(a, b) = g
 2. a = n * g
 3. b = m * g
 4. gcd(n, m) = 1
 5. a <= b => b -a = (m - n) * g
 6. a = n * g >= l => n 是最小的满足 n * g >= l
 7. 然后在 l....r 中间，找出g的倍数， 找出最大的和n互质的数（至少存在n + 1)
 8. 如果存在的话，是不是从最大的m开始 m * g <= r 开始，是不是可以很快的找到一个数和n是互质的？
 9. 因为质数的分布是lgn的，
 10. 但是 a = n * g 一定是最优的那个吗？