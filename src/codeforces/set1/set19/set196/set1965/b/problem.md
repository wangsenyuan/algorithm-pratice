You are given two integers 𝑛
 and 𝑘
. Find a sequence 𝑎
 of non-negative integers of size at most 25
 such that the following conditions hold.

There is no subsequence of 𝑎
 with a sum of 𝑘
.
For all 1≤𝑣≤𝑛
 where 𝑣≠𝑘
, there is a subsequence of 𝑎
 with a sum of 𝑣
.
A sequence 𝑏
 is a subsequence of 𝑎
 if 𝑏
 can be obtained from 𝑎
 by the deletion of several (possibly, zero or all) elements, without changing the order of the remaining elements. For example, [5,2,3]
 is a subsequence of [1,5,7,8,2,4,3]
.

It can be shown that under the given constraints, a solution always exists.

### ideas

1. 先考虑如何得到1....n的数
2. 假设x = 1111 (全1) <= n
3. 把x按位分解后，得到一组数，这样可以得到1...x的所有的数
4. 直接放置一个 n-x即可
5. 如果k > n => done
6. 这个有点麻烦呐
7. k <= n, let y = k - (n - x)
8. 如果y < 0, k < n - x
9. 那么k只能出现在前面的组合中。但是去掉任何一位，比如去掉 1 << d (d是k中的一位)那么会影响到很多位，
10. 假设 d-1也在k中出现了，那么就去掉d，但是增加 d-1 ？
11. 这是因为 1 << d + 1 << (d - 1) 不见了
12. 假设 k = 1 << d, 那么把 k去掉后， 增加 k + 1即可，其他的位数都仍然可以保证
13. 如果是普通的，那么把 1 << d (k的最高位去掉)， 然后增加 (1 << d) + 1

### solution
Notice that for a fixed 𝑘
, a solution for 𝑛=𝑐
 is also a solution for all 𝑛<𝑐
. So we can ignore the value of 𝑛
 and just assume it's always 106
.

If we didn't have the restriction that no subsequence can add up to 𝑘
, the most natural solution would be [1,2,4,8,⋯219]
. Every value from 1
 to 106
 appears as the sum of the subsequence given by its binary representation. We will use a modified version of this array to solve the problem.

Let 𝑖
 be the largest integer such that 2𝑖≤𝑘
. We will use this array (of size 22
):

𝑎=[𝑘−2𝑖,𝑘+1,𝑘+1+2𝑖,1,2,4,...2𝑖−1,2𝑖+1,...219]
To prove that no subsequence of 𝑎
 adds up to 𝑘
, consider the list of all elements in the array that are at most 𝑘
, since these are the only ones that could be present in a subsequence adding to 𝑘
. These are

𝑘−2𝑖,1,2,4,...2𝑖−1
Since these add up to 𝑘−1
, no subsequence can add up to 𝑘
.

To prove that for all 1≤𝑣≤𝑛
 where 𝑣≠𝑘
, there is a subsequence adding up to 𝑣
, we consider several cases:

If 𝑣<2𝑖
, we can simply use the binary representation of 𝑣
.
If 2𝑖≤𝑣<𝑘
, we can first take all of the elements that are at most 𝑘
 as part of our subsequence. We then need to remove elements with a sum equal to 𝑘−1−𝑣
. Because 2𝑖≤𝑣<𝑘<2𝑖+1
, 𝑘−1−𝑣
 is less than 2𝑖
, so we can simply remove its binary representation.
If 𝑣>𝑘
, we can take 𝑘+1
 along with the binary representation of 𝑣−𝑘−1
. The one edge case is when the 2𝑖
 bit is set in 𝑣−𝑘−1
. In this case, we replace 𝑘+1
 with 𝑘+1+2𝑖
.
So in all cases, we can form a subsequence adding up to 𝑣
.

