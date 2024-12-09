You are given an integer 𝑛
and an array 𝑎1,𝑎2…,𝑎𝑛
of integers in the range [0,𝑛]
.

A permutation 𝑝1,𝑝2,…,𝑝𝑛
of [1,2,…,𝑛]
is good if, for each 𝑖
, the following condition is true:

the number of values ≤𝑖
in [𝑝1,𝑝2,…,𝑝𝑖]
is exactly 𝑎𝑖
.
Count the good permutations of [1,2,…,𝑛]
, modulo 998244353
.

### ideas

1. 从后往前吗？
2. 如果a[n] = x, 那么表示在p[1....n] 中小于等于n的数，恰好为a[n]
3. a[n] = n (只有这样)
4. a[n-1] = x 表示 ... <= n - 1
5. 如果 a[n-1] = n - 1, 那么 p[n-1] = n可以吗？不行，这样子a[n-1] = n - 2
6. p[n-1] = n - 1 可以吗？可以，但是前提是n不在[1....n-1]范围内, 那么只能是p[n] = n (f(n) = f(n-1))
7. p[n-1] = n - 2，n必须在[1....n-1]范围内，那么p[n]可以选择 1...n-1中的任意一个数吗？至少目前看是可以的
8. 但是，还有一个考虑，比如a[1] = 1, 那么p[1] = 1(那么此时)在 p[n]就不能选择1
9. 混乱。不大对
10. 