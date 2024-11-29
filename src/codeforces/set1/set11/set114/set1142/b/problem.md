Recently Lynyrd and Skynyrd went to a shop where Lynyrd bought a permutation 𝑝
 of length 𝑛
, and Skynyrd bought an array 𝑎
 of length 𝑚
, consisting of integers from 1
 to 𝑛
.

Lynyrd and Skynyrd became bored, so they asked you 𝑞
 queries, each of which has the following form: "does the subsegment of 𝑎
 from the 𝑙
-th to the 𝑟
-th positions, inclusive, have a subsequence that is a cyclic shift of 𝑝
?" Please answer the queries.

A permutation of length 𝑛
 is a sequence of 𝑛
 integers such that each integer from 1
 to 𝑛
 appears exactly once in it.

A cyclic shift of a permutation (𝑝1,𝑝2,…,𝑝𝑛)
 is a permutation (𝑝𝑖,𝑝𝑖+1,…,𝑝𝑛,𝑝1,𝑝2,…,𝑝𝑖−1)
 for some 𝑖
 from 1
 to 𝑛
. For example, a permutation (2,1,3)
 has three distinct cyclic shifts: (2,1,3)
, (1,3,2)
, (3,2,1)
.

A subsequence of a subsegment of array 𝑎
 from the 𝑙
-th to the 𝑟
-th positions, inclusive, is a sequence 𝑎𝑖1,𝑎𝑖2,…,𝑎𝑖𝑘
 for some 𝑖1,𝑖2,…,𝑖𝑘
 such that 𝑙≤𝑖1<𝑖2<…<𝑖𝑘≤𝑟
.


### ideas
1. 可以对数字进行重新编排，将p按照1,2,3..., 然后将a重新编码
2. 那么查询a中是否存在某个cyclic shift就变成了， 是否存在两端，连续的子序列（或者一段）
3. 然后改写a数组
4. dp[i] = 最近的j，在区间i...j内包含了1，2。。。n的一个shift
5. dp[i] = dp[i+1] 
6.  或者因为 a[i]的加入，产生了新的一个shift,
7.  假设j = fp[a[i] + 1] = 从a[i] + 1 到n的位置
8.  dp[i] = gp[j+1] 在j+1后，能够从1到a[i] - 1 的位置
9.  所以，需要知道j+1能够到达的a[i]-1的位置  
10. 这个要怎么计算呢？