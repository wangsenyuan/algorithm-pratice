You are given an integer 𝑛
.

For each (𝑚,𝑘)
 such that 3≤𝑚≤𝑛+1
 and 0≤𝑘≤𝑛−1
, count the permutations of [1,2,...,𝑛]
 such that 𝑝𝑖+𝑝𝑖+1≥𝑚
 for exactly 𝑘
 indices 𝑖
, modulo 998244353
.

Input
The input consists of a single line, which contains two integers 𝑛
, 𝑥
 (2≤𝑛≤4000
, 1≤𝑥<1000000007
).

### ideas
1. 在给定m,k的情况下，如何计算f(m, k)呢？
2. f(m, k) = 排列的数量, p[i] + p[i+1] >= m and 这样的i恰好有k个；
3. 比如当n=3的时候， f(3, 2) = 6
4. [1, 2, 3], [2, 1, 3], [3, 2, 1], [3, 1, 2] ... 
5. f(3, 3) = 0 (最多有n-1个位置) 
6. f(4, 2) = 1 (1, 3, 2)
7. 在给定m的情况下
8. f(i, j) 表示在前i个数中，最后一个数是j，
9. 这样子不行，比如 上吗的 m = 4, 当前处理i = 3, 那么将i放在不同的位置，这个个数的贡献是不一样的，比如放在最前面, 它有可能+1
10. 放在中间某个位置，有可能+0/1/2 等
11. 所以，靠末尾的数字是没法确定的
12. 就是一个permuation，它最大的是n + n - 1, 最小的是 1 + 2 = 3
13. 当 m = 从最大到最小的过程中，这个k是怎么变化的？
14. 当 m = n + n - 1时，f(m, 1)是有效的, 这时这样的permation = 2 * (n - 1)个， （n 和 n - 1) 必须排在一起，和剩余的一起排列，
15. m = n + n - 1 - 1时， n, n - 1, n - 2,   f(m, 1) = n和n-2一起， n 和 n- 1一起 f(m, 1) 是有效的 + f(m, 2) (n在n-1, n-2中间时)
16. 假设一个permution满足 f(m, k) 是否能满足 f(m-1, k)?
17. 这个是成立的，因为p[i] + p[i+1] >= m > m - 1 
18. f(m-1, k) += f(m, k) 增加的部分 = p[i] + p[i-1] == m - 1 （最小值） 的部分
19. g(m, k) = 切好有一对 p[i] + p[i-1] = m, 且其他的 k-1对至少是 m+1的结果
20. f(m-1, k) = f(m, k) + g(m-1, k)
21. 如何计算g(m, k)呢？
22. 比如n=4, m = 6时，4 + 2 = 6， 也就是说在，所以4出现的地方， 它旁边必须出现2
23. 但是这里有个问题，比如 1, 4, 3, 如果把2放在4的前面， g(6, k) = f(7, k) (没有改变k的个数)
24. 但是如果把2放在2的后面, g(6, k) = f(7, k+1) 这是因为2在4， 3中间时，破坏了一个7
25. 似乎还是不大行
26. m = n + n - 1， f(m, 1) = n - 1  * 2
27. m = n + n - 2,  f(m, 2) = n - 2  * 2, (n - 1, n, n - 2) 组成一团进行排列
28.     但是 f(m, 1)就比较麻烦了，(n, n - 2)组成一队， 2 * (n - 1)中组合
29.       (n, n - 1) 组成一队, 2 * (n - 1)中组合
30.      还有就是， (n, n - 1, n - 2)组成一队 2 * (n-2)中组合， n必须在中间的组合
31.      f(m, 1) = 2 * 2 * (n - 1) - 2 * (n - 2)
32.      f(m, 1) + f(m, 2) = 2 * 2 * (n - 1) （和k没有关系了）？
33.  f(m, k) = 有多少pair x + y = m， 假设有w组， f(m, k) = 2 * nCr(w, 1) * (n - 1)
34.  w = 共有多少组x + y = m, 从中选择1组，然后排列出来
35.  pow(x, mn + k) = pow(x, mn) * pow(x, k)
36.  pow(x, 1) + pow(x, 2) + ... + pow(x, k) = 等比
37. 这个solution没有看懂。再等段时间看看。

### solution
Let's solve for a single 𝑚
. Suppose 𝑚
 is even. Start from an empty array, and insert the elements in the order [𝑚/2,𝑚/2−1,𝑚/2+1,𝑚/2−2,…]
. At any moment, all the elements are concatenated, and you can insert new elements either at the beginning, at the end or between two existing elements.

When you insert an element ≥𝑚/2
, the sum with any of the previous inserted elements is ≥𝑚
.
Otherwise, the sum is <𝑚
.
So you can calculate 𝑑𝑝𝑖,𝑗=
 number of ways to insert the first 𝑖
 elements (of [𝑚/2,𝑚/2−1,𝑚/2+1,𝑚/2−2,…]
) and make 𝑗
 "good" pairs (with sum ≥𝑚
).

You can split the ordering [𝑚/2,𝑚/2−1,𝑚/2+1,𝑚/2−2,…]
 into two parts:

small and big elements alternate;
there are only small elements.
For the second part, you don't need DP. Suppose you have already inserted 𝑖
 elements, and there are 𝑗
 good pairs, but when you will have inserted all elements you want 𝑘
 good pairs. The number of ways to insert the remaining elements can be computed with combinatorics in 𝑂(1)
 after precomputing factorials and inverses (you have to choose which pairs to break and use stars and bars; we skip exact formulas because they are relatively easy to find).

If you rearrange the factorials correctly, you can get that all the answers for a fixed 𝑚
 can be computed by multiplying two polynomials, one of which contains the 𝑑𝑝𝑖,𝑗
, where 𝑖
 is equal to the length of the "alternating" prefix. NTT is fast enough.

Complexity: 𝑂(𝑛2log𝑛)