This is an interactive problem.

There are two hidden non-negative integers 𝑥
 and 𝑦
 (0≤𝑥,𝑦<230
). You can ask no more than 2
 queries of the following form.

Pick a non-negative integer 𝑛
 (0≤𝑛<230
). The judge will respond with the value of (𝑛|𝑥)+(𝑛|𝑦)
, where |
 denotes the bitwise OR operation.
After this, the judge will give you another non-negative integer 𝑚
 (0≤𝑚<230
). You must answer the correct value of (𝑚|𝑥)+(𝑚|𝑦)
.

### ideas
1. 按位考虑, (0, 0), (1, 1), (1, 0), (0, 1)
2. 如果n[i] = 0  00 10, 01, 01
3. 如果n[i] = 1, 10 10, 10, 10
4. 要有两位i/i+1, 才能判断当前位是哪种组合
5. 10  和 x[i+1..i], y[i+1..i]
6. 如果x[i]或者y[i]只有一个1，那么用n[i] = 0, 那么s[i] = 1
7. x|0 + y|0 = x + y = s
8. 考虑最低位，如果s[0] = 0, 还无法判断
9. 1111 + 1 => 10000
10. 但也可能是 1110 + 10 => 10000
11. a | b = a + b - a & b
12. x | n + y | n = x + n - x & n + y + n - y & n
13. = x + y + 2 * n - (x & n + y & n)
14. s1 = f(n1), s2 = f(n2)
15. s1 + s2 = 2 * x + 2 * y +2 * n1 + 2 *n2 - (x & n1 + x & n2 + y & n1 + y & n2)
16. 如果n1 & n2 = 0,  且n1 | n2 = 1 << 31 - 
17. s1 + s2 = x + y + 2 * (n1 + n2)
18. x + y = s (就可以计算出来)
19. n2 = 2 * n1
20. s2 - s1 = 2 * (n2 - n1) - (x & n2 - x & n1) - (y & n2 - y & n1)
21.    = 2 * n1 - x / 2 - y / 2 ?
22. 