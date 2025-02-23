You are given an array 𝑎
 of 𝑛
 integers, where all elements except for at most one are equal to −1
 or 1
. The remaining element 𝑥
 satisfies −109≤𝑥≤109
.

Find all possible sums of subarrays of 𝑎
, including the empty subarray, whose sum is defined as 0
. In other words, find all integers 𝑥
 such that the array 𝑎
 has at least one subarray (possibly empty) with sum equal to 𝑥
. A subarray is a contiguous subsegment of an array.

Output these sums in ascending order. Each sum should be printed only once, even if it is achieved by multiple subarrays.

### ideas
1. 如果没有x，由[-1, 1]组成的数组要怎么处理呢？
2. 似乎可以使用一个区间[lo, hi]来处理？
3. 就是要计算出最大的sum/和最小的sum，在这个区间内的，都可以得到
4. 但是如何x呢？x的问题是，使的这个区间不连续了
5. 