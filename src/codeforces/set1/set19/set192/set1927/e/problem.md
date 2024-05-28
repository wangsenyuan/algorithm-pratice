You are given two integers 𝑛
 and 𝑘
 (𝑘≤𝑛
), where 𝑘
 is even.

A permutation of length 𝑛
 is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in any order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (as 2
 appears twice in the array) and [0,1,2]
 is also not a permutation (as 𝑛=3
, but 3
 is not present in the array).

Your task is to construct a 𝑘
-level permutation of length 𝑛
.

A permutation is called 𝑘
-level if, among all the sums of continuous segments of length 𝑘
 (of which there are exactly 𝑛−𝑘+1
), any two sums differ by no more than 1
.

More formally, to determine if the permutation 𝑝
 is 𝑘
-level, first construct an array 𝑠
 of length 𝑛−𝑘+1
, where 𝑠𝑖=∑𝑖+𝑘−1𝑗=𝑖𝑝𝑗
, i.e., the 𝑖
-th element is equal to the sum of 𝑝𝑖,𝑝𝑖+1,…,𝑝𝑖+𝑘−1
.

A permutation is called 𝑘
-level if max(𝑠)−min(𝑠)≤1
.

Find any 𝑘
-level permutation of length 𝑛
.

### ideas
1. abs(p[k+1] - p[1]) <= 1
2. abs(p[k+2] - p[2]) <= 1
3. p[1] = 1, p[k+1] = 2, p[2 * k + 1] = 3
4. 理解错了，不是相邻的和差最多1，而是整个s的最大值、最小值差1
5. 上面那个式子还是成立的，是max(s) - min(s) <= 1 的必要条件，但不是充分条件
6. 考虑n = 4, k = 2
7. [1, 4, 2, 3]
8. n = 6, k = 3
9. [1,4,5,2,3,6] (10, 11, 10, 11)
10. n = 8, k = 3
11. [1,6,7 2,5,8,3,4] (14, 15, 14, 15, 16, 15)
12. 如果k是偶数，按规则, 先顺序，再降序，是没有问题的，因为它们正好可以抵消掉，除了第一个/最后一个
13. 但是k是奇数，那么中间的某个区间，会存在两个值（头、尾）比前面的大，造成差值为2
14. 比如上面的区间 [7, 2, 5, 8, 3] (8 - 7 = 1, 3 - 2 = 1) 它们的差值为2
15. k确定是偶数的～～
16. 