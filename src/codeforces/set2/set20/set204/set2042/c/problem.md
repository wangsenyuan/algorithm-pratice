Alice and Bob participate in a fishing contest! In total, they caught 𝑛
 fishes, numbered from 1
 to 𝑛
 (the bigger the fish, the greater its index). Some of these fishes were caught by Alice, others — by Bob.

Their performance will be evaluated as follows. First, an integer 𝑚
 will be chosen, and all fish will be split into 𝑚
 non-empty groups. The first group should contain several (at least one) smallest fishes, the second group — several (at least one) next smallest fishes, and so on. Each fish should belong to exactly one group, and each group should be a contiguous subsegment of fishes. Note that the groups are numbered in exactly that order; for example, the fishes from the second group cannot be smaller than the fishes from the first group, since the first group contains the smallest fishes.

Then, each fish will be assigned a value according to its group index: each fish in the first group gets value equal to 0
, each fish in the second group gets value equal to 1
, and so on. So, each fish in the 𝑖
-th group gets value equal to (𝑖−1)
.

The score of each contestant is simply the total value of all fishes that contestant caught.

You want Bob's score to exceed Alice's score by at least 𝑘
 points. What is the minimum number of groups (𝑚
) you have to split the fishes into? If it is impossible, you should report that.

### ideas
1. 求前缀和,假设分成了m段，那么有
2. 0 * (p[i1]) + 1 * (p[i2] - p[i1]) + 2 * (p[i3] - p[i2]) + .. + (m - 1) * (p[im] - p[im-1])
3. -p[i1] - p[i2] - p[i3] - ... - p[i(m-1)] + (m - 1) * p[im]
4. p[im] = p[n]
5. 也就是说 p[n] * (m - 1) - (p[i1] + p[i2] + ... p[im-1]) >= k
6. 也就是 p[i1] + ... + p[im-1] 要最小值
7. 