The prime factorization of a positive integer 𝑚
 is the unique way to write it as 𝑚=𝑝𝑒11⋅𝑝𝑒22⋅…⋅𝑝𝑒𝑘𝑘
, where 𝑝1,𝑝2,…,𝑝𝑘
 are prime numbers, 𝑝1<𝑝2<…<𝑝𝑘
 and 𝑒1,𝑒2,…,𝑒𝑘
 are positive integers.

For each positive integer 𝑚
, 𝑓(𝑚)
 is defined as the multiset of all numbers in its prime factorization, that is 𝑓(𝑚)={𝑝1,𝑒1,𝑝2,𝑒2,…,𝑝𝑘,𝑒𝑘}
.

For example, 𝑓(24)={2,3,3,1}
, 𝑓(5)={1,5}
 and 𝑓(1)={}
.

You are given a list consisting of 2𝑛
 integers 𝑎1,𝑎2,…,𝑎2𝑛
. Count how many positive integers 𝑚
 satisfy that 𝑓(𝑚)={𝑎1,𝑎2,…,𝑎2𝑛}
. Since this value may be large, print it modulo 998244353
.

### ideas
1. 显然pi都是质数，且互补相同
2. 假设有m个质数, 如果m < n => 0
3. else 选其中n个做底数 => C(m, n), 剩下的数做指数
4. 考虑这个例子 (2, 3, 2, 3) 
5. 要从distinct(prime numbers) 中选取n个做为底数
6. (2, 2, 2, 3, 3, 3, 5, 7)
7. 底数肯定是2, 3, 5, 7, 剩余(2, 2, 3, 3)
8. 但是对于（2， 2）表示成（a, b）
9. 那么 (2, a, 3, b) 和 (2, b, 3, a) 或者 (2, a, 5, b), (2, b, 5, a)都是相同的数 
10. 怎么区分呢？
11. 对于（{2, 2}, {3, 2}) 可以 C(4, 2) * C(2, 2), 先为2选择两个位置
12. 然后为3选择两个位置，确实是可以的
13. 但是这个没法和第一步连起来。是因为，第一步选走n个数后，剩余哪些数字，这个是不知道的，所以也就没法统计这个数字
14. 换个思路，把set当成一个列表，必须在奇数下标使用质数，在偶数位没有限制
15. dp[i][j] 表示到i为止，使用了j个质数的计数(j >= (i + 1) / 2)
16. dp[i][j] = 如果i是奇数, 必须使用质数 dp[i-1][j-1] * 能使用的质数的数量
17.    如果i是偶数 dp[i-1][j] * 非质数的数量 + dp[i-1][j-1] * 质数的数量
18.  不考虑i, dp[j] = dp[j-1] * 质数的数量， dp[j] * 非质数的数量 + dp[j-1] * 质数的数量
19.  good
20.  有问题，因为这个没有办法避免p[1] = p[2]的情况
21.  假设某个质数被作为了底数，那么它的贡献 C(n/2, m-1) * n/2 (n是当前剩余的位置)在偶数个位置(n/2) 选出m-1个位置让x作为指数，选出的x放到奇数位
22.  如果是非质数，那么就是 (C(n/2), m) 
23.  偶数位和奇数位要单独处理
24.  dp[i][j] 表示处理到第i个唯一的数目时，j表示底数的个数
25.  如果arr[i]是质数，那么它可以做底数，也可以不做
26.  那么它做底数的话，就是 dp[i-1][j-1] * (可用底数的个数 = n - (j - 1)) * (可用指数的个数)
27.   如果它不做底数的话，dp[i-1][j] * (可用指数的个数)
28.   到目前为止一共使用掉的位置 = sum(...i-1), 共有2 * n 个位置，其中j个是底数，那么指数使用了 sum - j 
29.   所以剩余 n - (sum - j)个位置