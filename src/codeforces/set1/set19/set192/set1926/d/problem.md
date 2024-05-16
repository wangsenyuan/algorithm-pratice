Vladislav has 𝑛
 non-negative integers, and he wants to divide all of them into several groups so that in any group, any pair of numbers does not have matching bit values among bits from 1
-st to 31
-st bit (i.e., considering the 31
 least significant bits of the binary representation).

For an integer 𝑘
, let 𝑘2(𝑖)
 denote the 𝑖
-th bit in its binary representation (from right to left, indexing from 1). For example, if 𝑘=43
, since 43=1010112
, then 432(1)=1
, 432(2)=1
, 432(3)=0
, 432(4)=1
, 432(5)=0
, 432(6)=1
, 432(7)=0
, 432(8)=0,…,432(31)=0
.

Formally, for any two numbers 𝑥
 and 𝑦
 in the same group, the condition 𝑥2(𝑖)≠𝑦2(𝑖)
 must hold for all 1≤𝑖<32
.

What is the minimum number of groups Vlad needs to achieve his goal? Each number must fall into exactly one group.


### ideas

1. 不仅仅是1，0也不能一样， 所以只有 x 和 mask ^ x可以分为一组