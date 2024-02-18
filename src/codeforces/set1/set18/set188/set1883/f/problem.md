You are given an array of integers 𝑎1,𝑎2,…,𝑎𝑛
. Calculate the number of subarrays of this array 1≤𝑙≤𝑟≤𝑛
, such that:

The array 𝑏=[𝑎𝑙,𝑎𝑙+1,…,𝑎𝑟]
occurs in the array 𝑎
as a subsequence exactly once. In other words, there is exactly one way to select a set of indices 1≤𝑖1<𝑖2<…<
𝑖𝑟−𝑙+1≤𝑛
, such that 𝑏𝑗=𝑎𝑖𝑗
for all 1≤𝑗≤𝑟−𝑙+1
.

### thoughts

1. 假设选中了一个sub array, [l...r]， 如何确定它是唯一的？
2. 如果在l的前面，还存在一个a[l], 或者r的后面，还存在一个a[r]，那么它就不是唯一的
3. 除此之外是不是一定是唯一的？
4. 对于任何一个数，如果它是这个数的最左边的位置，那么它可以和所有右端点的数，匹配