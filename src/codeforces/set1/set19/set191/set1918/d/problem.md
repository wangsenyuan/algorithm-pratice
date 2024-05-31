You are given an array of numbers 𝑎1,𝑎2,…,𝑎𝑛
. Your task is to block some elements of the array in order to minimize its cost. Suppose you block the elements with indices 1≤𝑏1<𝑏2<…<𝑏𝑚≤𝑛
. Then the cost of the array is calculated as the maximum of:

the sum of the blocked elements, i.e., 𝑎𝑏1+𝑎𝑏2+…+𝑎𝑏𝑚
.
the maximum sum of the segments into which the array is divided when the blocked elements are removed. That is, the maximum sum of the following (𝑚+1
) subarrays: [1,𝑏1−1
], [𝑏1+1,𝑏2−1
], […
], [𝑏𝑚−1+1,𝑏𝑚−1
], [𝑏𝑚+1,𝑛
] (the sum of numbers in a subarray of the form [𝑥,𝑥−1
] is considered to be 0
).
For example, if 𝑛=6
, the original array is [1,4,5,3,3,2
], and you block the elements at positions 2
 and 5
, then the cost of the array will be the maximum of the sum of the blocked elements (4+3=7
) and the sums of the subarrays (1
, 5+3=8
, 2
), which is max(7,1,8,2)=8
.

You need to output the minimum cost of the array after blocking.

### ideas
1. 分出来的block越多，每段的sum就越小，但是blocked sum就越大
2. 能二分吗？比如说通过某种方式能够得到x。那么x+1也肯定是可行的吗？
3. 先假设可行，那在给定x的情况下，如何block，以使得cost <= x 呢？
4. 在这种情况下，对于r计算出f[r] 表示，最元的l, sum[l...r] <= x
5. 因为要尽量少的出现block的数字，假设选中了一段，且这段是从r结束的
6. 那么它肯定是从l开始的
7. 那这样子是最优的吗？比如a[l-1] > a[l], 如果选择 a[l-1]到block的数字中时，> x
8. 但是选择 a[l]不会，会出现这种情况吗？
9. 会的。因为选择a[l-1]的区间，可以保证它不超过x
10. 重新思考一下，就是要在i前，保证每段sum <= x的情况下，是的block的和最小
11. 如果a[i]被block掉, dp[i] = dp[i-1] + x
12. 如果a[i]没有被block, dp[i] = min(dp[f[l]], dp[f[l] + 1], ... dp[i-1]) 
13. 