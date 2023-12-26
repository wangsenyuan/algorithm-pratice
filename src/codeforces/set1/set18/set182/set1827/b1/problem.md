You are given an array 𝑎
, consisting of 𝑛
distinct integers 𝑎1,𝑎2,…,𝑎𝑛
.

Define the beauty of an array 𝑝1,𝑝2,…𝑝𝑘
as the minimum amount of time needed to sort this array using an arbitrary number of range-sort operations. In each
range-sort operation, you will do the following:

Choose two integers 𝑙
and 𝑟
(1≤𝑙<𝑟≤𝑘
).
Sort the subarray 𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟
in 𝑟−𝑙
seconds.
Please calculate the sum of beauty over all subarrays of array 𝑎
.

A subarray of an array is defined as a sequence of consecutive elements of the array.

### thoughts

1. 先考虑对于arr来说，需要的最少range-sort次数
2. 例子 4 8 7 2 =》 1, 3, 2, 0
    1. [3, 2] => 1
    2. [2, 0] => 1
    3. [1, 3, 2] => 1
    4. [3, 2, 0] => 2
    5. [1, 3, 2, 0] =>  3
3. f(arr) = 最少的数量
    1. 如果 arr[0] 存在正确的位置上 => f(arr) = f(arr[1:])
    2. 如果 last 在正确的位置上 => f(arr) = f(arr/last)
    3. 否则的话， f(arr) = len(arr) - 1
4. 考虑位置r，dp[r] = 以r为右边界的所有子数组的的和
5. 如果 b[r] > 目前的最大值, dp[r] = dp[r-1]
6. else if b[r] < 目前的最大值
7. 考虑一个下标l, l...r 如果 b[r]仍然是这个区间内的最大值，还是不需要交换的
8. 假设 b[l] > b[r] 就需要交换了，也就是右边那些 > b[r] 的下标都需要和r进行sort
9. 包含l...r 的区间数量 = (l - k) * dist=
10. 对于r，需要知道左边有多少个比它大的数 (0 + r) * (r - 1) / 2 - sum of 比它小的数
11. 还需要知道count
   