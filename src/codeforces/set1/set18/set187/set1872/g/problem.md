Given an array 𝑎
of 𝑛
positive integers. You need to perform the following operation exactly once:

Choose 2
integers 𝑙
and 𝑟
(1≤𝑙≤𝑟≤𝑛
) and replace the subarray 𝑎[𝑙…𝑟]
with the single element: the product of all elements in the subarray (𝑎𝑙⋅…⋅𝑎𝑟)
.
For example, if an operation with parameters 𝑙=2,𝑟=4
is applied to the array [5,4,3,2,1]
, the array will turn into [5,24,1]
.

Your task is to maximize the sum of the array after applying this operation. Find the optimal subarray to apply this
operation.

#### thoughts

1. 在给定r的时候， 是不是能够贪心的指定l，如果 pref[l-1] + suf[r+1] + prod[l...r] 更大
2. 只要 prod[l...r] > sum[l...r] 似乎就可以选择
3. 一个观察是，如果这个区间内，所有的数，都大于1, 那么选择它们是没有问题的
4. 只有当区间内包含1的数量比较多时，才可能造成差的选择
5. 假设 prod[l...r] < sum[l....r], 这里面至少有多少个1个？
6. 假设中间有x个1，sum[l...r] = x + y (y = 至少2以上的数相乘)
7. 假设l...r的长度为ln 那么这样的2以上的个数，不会超过20个， 所以，只要检查前面20个非1的位置即可