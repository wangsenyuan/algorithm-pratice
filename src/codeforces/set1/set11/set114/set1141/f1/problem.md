You are given an array of integers 𝑎[1],𝑎[2],…,𝑎[𝑛].
 A block is a sequence of contiguous (consecutive) elements 𝑎[𝑙],𝑎[𝑙+1],…,𝑎[𝑟]
 (1≤𝑙≤𝑟≤𝑛
). Thus, a block is defined by a pair of indices (𝑙,𝑟)
.

Find a set of blocks (𝑙1,𝑟1),(𝑙2,𝑟2),…,(𝑙𝑘,𝑟𝑘)
 such that:

They do not intersect (i.e. they are disjoint). Formally, for each pair of blocks (𝑙𝑖,𝑟𝑖)
 and (𝑙𝑗,𝑟𝑗
) where 𝑖≠𝑗
 either 𝑟𝑖<𝑙𝑗
 or 𝑟𝑗<𝑙𝑖
.
For each block the sum of its elements is the same. Formally,
𝑎[𝑙1]+𝑎[𝑙1+1]+⋯+𝑎[𝑟1]=𝑎[𝑙2]+𝑎[𝑙2+1]+⋯+𝑎[𝑟2]=
⋯=
𝑎[𝑙𝑘]+𝑎[𝑙𝑘+1]+⋯+𝑎[𝑟𝑘].
The number of the blocks in the set is maximum. Formally, there does not exist a set of blocks (𝑙′1,𝑟′1),(𝑙′2,𝑟′2),…,(𝑙′𝑘′,𝑟′𝑘′)
 satisfying the above two requirements with 𝑘′>𝑘
.
The picture corresponds to the first example. Blue boxes illustrate blocks.
Write a program to find such a set of blocks.

### ideas
1. 应该可以二分。但二分的时候，不好检查
2. n <= 1500, 所以 n * n 的算法应该是可行的
3. 确定s后，在其中计算有多少个s，这个初步感觉是 n * n * n 的
4. 但是如果dp[i][s] 表示到i为止，且sum=s的segment的数量是多少
5. dp[i][s] = dp[i-1][s]
6. 然后从j...i = x, dp[i][x] = max(dp[i][x], dp[j-1][x] + 1)
7. 可以用lazy tree