You are given a permutation 𝑝1,𝑝2,…,𝑝𝑛
 (an array where each integer from 1
 to 𝑛
 appears exactly once). The weight of the 𝑖
-th element of this permutation is 𝑎𝑖
.

At first, you separate your permutation into two non-empty sets — prefix and suffix. More formally, the first set contains elements 𝑝1,𝑝2,…,𝑝𝑘
, the second — 𝑝𝑘+1,𝑝𝑘+2,…,𝑝𝑛
, where 1≤𝑘<𝑛
.

After that, you may move elements between sets. The operation you are allowed to do is to choose some element of the first set and move it to the second set, or vice versa (move from the second set to the first). You have to pay 𝑎𝑖
 dollars to move the element 𝑝𝑖
.

Your goal is to make it so that each element of the first set is less than each element of the second set. Note that if one of the sets is empty, this condition is met.

For example, if 𝑝=[3,1,2]
 and 𝑎=[7,1,4]
, then the optimal strategy is: separate 𝑝
 into two parts [3,1]
 and [2]
 and then move the 2
-element into first set (it costs 4
). And if 𝑝=[3,5,1,6,2,4]
, 𝑎=[9,1,9,9,1,9]
, then the optimal strategy is: separate 𝑝
 into two parts [3,5,1]
 and [6,2,4]
, and then move the 2
-element into first set (it costs 1
), and 5
-element into second set (it also costs 1
).

Calculate the minimum number of dollars you have to spend.


### ideas
1. 考虑最终第一个set的大小，假设是x
2. 那么就是把前x中 > x 的部分移动到后边 + 后面（n-x中) <= x 的部分移动到前面来