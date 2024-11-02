Recall that the permutation is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in arbitrary order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears twice in the array) and [1,3,4]
 is also not a permutation (𝑛=3
 but there is 4
 in the array).

A sequence 𝑎
 is a subsegment of a sequence 𝑏
 if 𝑎
 can be obtained from 𝑏
 by deletion of several (possibly, zero or all) elements from the beginning and several (possibly, zero or all) elements from the end. We will denote the subsegments as [𝑙,𝑟]
, where 𝑙,𝑟
 are two integers with 1≤𝑙≤𝑟≤𝑛
. This indicates the subsegment where 𝑙−1
 elements from the beginning and 𝑛−𝑟
 elements from the end are deleted from the sequence.

For a permutation 𝑝1,𝑝2,…,𝑝𝑛
, we define a framed segment as a subsegment [𝑙,𝑟]
 where max{𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟}−min{𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟}=𝑟−𝑙
. For example, for the permutation (6,7,1,8,5,3,2,4)
 some of its framed segments are: [1,2],[5,8],[6,7],[3,3],[8,8]
. In particular, a subsegment [𝑖,𝑖]
 is always a framed segments for any 𝑖
 between 1
 and 𝑛
, inclusive.

We define the happiness of a permutation 𝑝
 as the number of pairs (𝑙,𝑟)
 such that 1≤𝑙≤𝑟≤𝑛
, and [𝑙,𝑟]
 is a framed segment. For example, the permutation [3,1,2]
 has happiness 5
: all segments except [1,2]
 are framed segments.

Given integers 𝑛
 and 𝑚
, Jongwon wants to compute the sum of happiness for all permutations of length 𝑛
, modulo the prime number 𝑚
. Note that there exist 𝑛!
 (factorial of 𝑛
) different permutations of length 𝑛
.

### ideas
1. 在一个给定的序列中，f(arr) = 怎么计算呢？
2. f(1) = 1
3. 假设所有的都是happy的， n! * n * (n + 1) / 2
4. f(2) = 2! * 2 * 3 / 2 = 6
5. 什么情况下是不happy的呢？ 