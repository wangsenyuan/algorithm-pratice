The mayor of the Central Town wants to modernize Central Street, represented in this problem by the (𝑂𝑥)
 axis.

On this street, there are 𝑛
 antennas, numbered from 1
 to 𝑛
. The 𝑖
-th antenna lies on the position 𝑥𝑖
 and has an initial scope of 𝑠𝑖
: it covers all integer positions inside the interval [𝑥𝑖−𝑠𝑖;𝑥𝑖+𝑠𝑖]
.

It is possible to increment the scope of any antenna by 1
, this operation costs 1
 coin. We can do this operation as much as we want (multiple times on the same antenna if we want).

To modernize the street, we need to make all integer positions from 1
 to 𝑚
 inclusive covered by at least one antenna. Note that it is authorized to cover positions outside [1;𝑚]
, even if it's not required.

What is the minimum amount of coins needed to achieve this modernization?

### ideas
1. 需要把[1...m]的区间都cover住，如果只有一个的话，是很好处理的
2. 但是问题有n个
3. dp[i][j] 表示使用前j个天线将位置[1...i]都cover住的情况
4. dp[i][j] = dp[i-1][j] + 1 仍然是用第j个cover, 并将其升高（如果需要，也有可能不需要）
5.           当 x[j] + s[j] >= i 的时候，不需要
6.           或者使用 dp[x][j-1] + 使用第j个，把 x....i给cover住
7.           这里考虑两种情况, 
8.           一种是 (x, i)都在 j的范围内, x >= x[j] - s[j]， 且 i <= x[j] + s[j]时
9.           也就是找出x[j] - s[j] 到i范围内的最小值(for j-1)
10.          如果是 (x[j] - x) > min(s[j], i - x[j])(此时肯定要增加高度了)
11.          假设位置是y， 那么要增加的高度 = max(x[j] - y - s[j], i - x[j] - s[j])
12.          上面那个式子还是不好处理，再分两种情况 
13.          x[j] - y - s[j] >= i - x[j] - s[j] (前面的间隔更大)
14.          y <= x[j] - s[j], y >= i - (x[j] - s[j]) 这个区间内的最小值
15.          dp[y][j-1] + x[j] - s[j] - y 的最小值
16.          如果 x[j] - y - s[j] <= i - x[j] - s[j]
17.          y <= min(x[j] - s[j], i - x[j] - s[j])
18.          dp[y][j-1] + i - x[j] - s[j] （单纯就是dp[y][j-1])的最小值