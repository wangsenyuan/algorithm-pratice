This is an interactive problem!

Nastia has a hidden permutation 𝑝
 of length 𝑛
 consisting of integers from 1
 to 𝑛
. You, for some reason, want to figure out the permutation. To do that, you can give her an integer 𝑡
 (1≤𝑡≤2
), two different indices 𝑖
 and 𝑗
 (1≤𝑖,𝑗≤𝑛
, 𝑖≠𝑗
), and an integer 𝑥
 (1≤𝑥≤𝑛−1
).

Depending on 𝑡
, she will answer:

𝑡=1
: max(min(𝑥,𝑝𝑖),min(𝑥+1,𝑝𝑗))
;
𝑡=2
: min(max(𝑥,𝑝𝑖),max(𝑥+1,𝑝𝑗))
.
You can ask Nastia at most ⌊3⋅𝑛2⌋+30
 times. It is guaranteed that she will not change her permutation depending on your queries. Can you guess the permutation?

### ideas
1. 如果 x = 1
2. t = 1, max(min(1, pi), min(2, pj))
3.  = max(1, min(2, pj)) = 2, 如果 pj <= 2 那么就是 pj 
4.  似乎没啥用
5. 如果 x = n - 1
6. t = 1, max(min(n - 1, pi), min(n, pj))
7.  = max(min(n-1, pi), pj) 很有可能是 max(pi, pj)
8.  除非 pi > n - 1 (也就是 p[i] = n)
9. 如果 t = 1, 且 max(...) = n => p[j] = n 
10. 如果 max(min(n-1, pi), min(n, pj)) = n - 1 (这个pi = n - 1, or p[j] = n - 1 都有可能)
11. 知道找到了n, 那么就可以用 query 2, min(max(x, n), max(x+1, pj)) 来确定pj (x = 1)
12. 这样子至少要 2 * n 次查询
13. 如果两个一组，肯定在某组里面查到 = n 吗? 不一定
14. min(max(x, 1), max(x+1, pj)) = 1 如果 x = 1, 且p[i] = 1 时
15. min(max(1, pi), max(2, pj)) = min(pi, max(2, pj)) 
16. 两个一组，如果存在 ans = 1的，那么 p[i] = 1
17. 如果 p[j] = 1 =》 上面的那个 = 2， 但是如果 上面的等于 2, 还有一种可能性, pi = 2
18. 也就是说通过n/2次查询， 找到p[i] = 1, (ans = 1)
19. 或者 p[i] = 2 (p[j] > 2), 或者p[j] = 1 (p[i] >= 2) 的
20. 如果只有一个 ans = 2的,就知道了两个
21. 如果有两个，就需要确定到底是p[i] = 2, 还是 p[j] = 1?
22. min(max(1, p[j]), max(2, p[i])) = 1 =》 p[j] =  1
23. 如果 = 2 =》 p[i] = 2
24. 找到1以后，就简单了