You are given 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 and an integer 𝑘
. Find the maximum value of 𝑖⋅𝑗−𝑘⋅(𝑎𝑖|𝑎𝑗)
 over all pairs (𝑖,𝑗)
 of integers with 1≤𝑖<𝑗≤𝑛
. Here, |
 is the bitwise OR operator.

Input
The first line contains a single integer 𝑡
 (1≤𝑡≤10000
)  — the number of test cases.

The first line of each test case contains two integers 𝑛
 (2≤𝑛≤105
) and 𝑘
 (1≤𝑘≤min(𝑛,100)
).

The second line of each test case contains 𝑛
 integers 𝑎1,𝑎2,…,𝑎𝑛
 (0≤𝑎𝑖≤𝑛
).

It is guaranteed that the sum of 𝑛
 over all test cases doesn't exceed 3⋅105
.

Output
For each test case, print a single integer  — the maximum possible value of 𝑖⋅𝑗−𝑘⋅(𝑎𝑖|𝑎𝑗)
.

### ideas
1. i * j - k * (a[i] | a[j]) 最大的数 (k <= 100)
2. 在fix j的情况下，希望i越大越好，（前面就会足够大),且a[i]是 a[j]的子集最好（这样子不会贡献额外的位数）
3. a[i] | a[j] = a[i] + a[j] - a[i] & a[j]
4. i * j - k * (a[i] + a[j] - a[i] & a[j])
5. i * j - k * (a[i] + a[j]) + k * a[i] & a[j]
6. 假设有两位x, y, 它们与a[j]的or，都贡献了额外的一位
7. 且，a[x]贡献u, x[y]贡献v, 其他的都一样 
8. x * j - k * (a[x] | a[j])
9. y * j - k * (a[y] | a[j])
10. (x - y) * j - k * (1 << u - 1 << v)
11. 希望x更优， 那么有 (x - y) * j > k * (1 << u - 1 << v)
12. 上面可以看出来， j > 0, 所以它对符号是没贡献的，可以直接忽略掉, 还有k，也是没贡献的
13. x - y > pow(2, u) - pow(2, v)
14. 假设x - pow(2, u) > y - pow(2, v)
15. 也就是说，对于给定的j来说，就是要找到位置x, 且 x - (a[x] 那些不在a[j]中的的位)要最大
16. i - a[i] | a[j] 最大
17. i - (a[i] + a[j] - a[i] & a[j])
18. i - a[i] - a[i] & a[j] 最大 (i - a[i])是确定的, a[i] & a[j]时， i + a[i]的最大值
19. 也就是a[j]的子集中 i + a[i]的最大值