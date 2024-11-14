You are given an integer sequence 𝑎1,𝑎2,…,𝑎𝑛
. Let 𝑆
 be the set of all possible non-empty subsequences of 𝑎
 without duplicate elements. Your goal is to find the longest sequence in 𝑆
. If there are multiple of them, find the one that minimizes lexicographical order after multiplying terms at odd positions by −1
.

For example, given 𝑎=[3,2,3,1]
, 𝑆={[1],[2],[3],[2,1],[2,3],[3,1],[3,2],[2,3,1],[3,2,1]}
. Then [2,3,1]
 and [3,2,1]
 would be the longest, and [3,2,1]
 would be the answer since [−3,2,−1]
 is lexicographically smaller than [−2,3,−1]
.

A sequence 𝑐
 is a subsequence of a sequence 𝑑
 if 𝑐
 can be obtained from 𝑑
 by the deletion of several (possibly, zero or all) elements.

A sequence 𝑐
 is lexicographically smaller than a sequence 𝑑
 if and only if one of the following holds:

𝑐
 is a prefix of 𝑑
, but 𝑐≠𝑑
;
in the first position where 𝑐
 and 𝑑
 differ, the sequence 𝑐
 has a smaller element than the corresponding element in 𝑑
.

### ideas
1. 最长的很容易找出来，就是说有distinct的数字组成的那个序列
2. 问题是那个满足条件的最小的序列
3. 奇数位要尽可能的大，偶数位要尽可能的小
4. 想到了