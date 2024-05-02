Count the number of distinct sequences a1, a2, ..., an (1 ≤ ai) consisting of positive integers such that gcd(a1,
a2, ..., an)= x and . As this number could be large, print the answer modulo 109 + 7.

### ideas

1. gcd(a1, a2, ... ai) = x
2. => a[i] % x = 0, and some gcd(a[i], a[j]) = x
3. f[x] = 满足条件的能够整除x的计数
4. g[x] = 整除x，且gcd = x的计数
5. 能不能很容易的计算出f[z]?
6. z肯定时y的因数，且是x的倍数
7. 那么给定z时，y % z = 0, let m = y / z
8. let b[0] = a[0] / z, b[i] = a[i]/ z
9. 那么 sum of b = m
10. 也就是说存在一个序列，取每个数取自然数，和为m的个数
11. 那么就是在m个格子中间放置搁板，共有 m-1个位置，每个位置放或者不放 pow(2, m - 1)
12. f[x] = pow(2, m - 1) = pow(2, y / x - 1)
13. 然后计算g[x]
14. 