Toad Rash has a binary string 𝑠
. A binary string consists only of zeros and ones.

Let 𝑛
 be the length of 𝑠
.

Rash needs to find the number of such pairs of integers 𝑙
, 𝑟
 that 1≤𝑙≤𝑟≤𝑛
 and there is at least one pair of integers 𝑥
, 𝑘
 such that 1≤𝑥,𝑘≤𝑛
, 𝑙≤𝑥<𝑥+2𝑘≤𝑟
, and 𝑠𝑥=𝑠𝑥+𝑘=𝑠𝑥+2𝑘
.

Find this number of pairs for Rash.

#### ideas
1. 计算的是l...r, 如果l...r满足条件，那么l...r+1, r+2... 也是满足条件的
2. 所以l..r 要找到最小的r
3. 对于l如何找到这个r, 那么r = x + 2 * k (否则的话，r就不是最小的)
4. 也就是说要存在两个位置x, x + k 满足 s[x] = s[r], s[x+k] = s[r]
5. 如果s[l] = s[r],那么让x = l, 是否可行呢？
6. r - l = 2 * k, 验证k是否满足条件
7. 对于r来说，需要找到最大的x, s[x] = s[r], 且s[x + k] = s[r]
8. k = (r - x) / 2
9. 比如r = 5, x = 3, 那么 k = 1
10. 那么 l <= x即可
11. 如果固定 y = x + k, 那么  x = y - k, r = y + k
12. 好处理了