You are given two positive integers 𝑛
, 𝑚
.

Calculate the number of ordered pairs (𝑎,𝑏)
 satisfying the following conditions:

1≤𝑎≤𝑛
, 1≤𝑏≤𝑚
;
𝑎+𝑏
 is a multiple of 𝑏⋅gcd(𝑎,𝑏)
.


### ideas
1. let x = b * gcd(a, b)
2. let y = a + b
3. y % x = 0
4. 如果 x = gcd(a, b), 那么 y = a + b 肯定整除 x
5. a = g * m
6. b = g * n
7. y = g * (m + n)
8. x = g * n * g
9. (m + n) % (g * n)
10. m和n是互质的数，m + n 要能整除n的倍数
11. 如果 m < n => 显然是不行的，m + n < 2 * n
12. m > n, m + n = g * n => m = (g - 1) * n, 这里不就gcd(m, n) = n 了吗？
13. 和m,n互质冲突
14. (a, b) = (6, 2)
15. gcd(a, b) = 2
16. a + b = 8, 8 % (2 * 2) = 0
17. 当n=1的时候是成立的
18. gcd(a, b) = b, a是b的倍数
19. a + b = (m + 1) * b % (b * b) = 0
20. 也就是 (m + 1) % b = 0
21. a是b的倍数，且a = x * b, x + 1 也是b的倍数
22. 如果 b的倍数
23. b <= m
24. x * b <= n
25. 且 (x + 1) % b == 0
26. x <= n / b
27. 然后在这个范围内，找到 b倍数
28. (n / b + 1) / b
29. 但是这里不能迭代b呐。咋搞呢？
30. 如果不考虑 (n / b + 1) % b = 0的情况
31. 那么 n / b / b 的结果
32. 当 b * b > n 的时候， 这个结果就是0了， 所以，可以只处理 前sqrt(n)个即可