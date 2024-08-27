Let 𝑎
 and 𝑏
 be two arrays of lengths 𝑛
 and 𝑚
, respectively, with no elements in common. We can define a new array merge(𝑎,𝑏)
 of length 𝑛+𝑚
 recursively as follows:

If one of the arrays is empty, the result is the other array. That is, merge(∅,𝑏)=𝑏
 and merge(𝑎,∅)=𝑎
. In particular, merge(∅,∅)=∅
.
If both arrays are non-empty, and 𝑎1<𝑏1
, then merge(𝑎,𝑏)=[𝑎1]+merge([𝑎2,…,𝑎𝑛],𝑏)
. That is, we delete the first element 𝑎1
 of 𝑎
, merge the remaining arrays, then add 𝑎1
 to the beginning of the result.
If both arrays are non-empty, and 𝑎1>𝑏1
, then merge(𝑎,𝑏)=[𝑏1]+merge(𝑎,[𝑏2,…,𝑏𝑚])
. That is, we delete the first element 𝑏1
 of 𝑏
, merge the remaining arrays, then add 𝑏1
 to the beginning of the result.
This algorithm has the nice property that if 𝑎
 and 𝑏
 are sorted, then merge(𝑎,𝑏)
 will also be sorted. For example, it is used as a subroutine in merge-sort. For this problem, however, we will consider the same procedure acting on non-sorted arrays as well. For example, if 𝑎=[3,1]
 and 𝑏=[2,4]
, then merge(𝑎,𝑏)=[2,3,1,4]
.

A permutation is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in arbitrary order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears twice in the array) and [1,3,4]
 is also not a permutation (𝑛=3
 but there is 4
 in the array).

There is a permutation 𝑝
 of length 2𝑛
. Determine if there exist two arrays 𝑎
 and 𝑏
, each of length 𝑛
 and with no elements in common, so that 𝑝=merge(𝑎,𝑏)
.


### ideas
1. 考虑后缀, dp[i][j]表示在后缀有j个元素和p[i]在一个分组中
2. 如果 p[i] < p[i+1], dp[i][j] = dp[i+1][j-1] (p[i],和 p[i+1]在一个分组)
3.    或者 dp[i][j] = dp[i+1][n - i - j] (p[i] 和 p[i+1]在不同的分组)
4.  如果 p[i] > p[i+1], 那么p[i]和p[i+1]似乎只能在一个分组？
5.  如果它们在两个分组，那么p[i+1]就必须先出现
6.  dp[i][j] = dp[i+1][j-1]