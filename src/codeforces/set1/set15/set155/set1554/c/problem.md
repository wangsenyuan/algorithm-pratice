You are given two integers 𝑛
 and 𝑚
. Find the MEX
 of the sequence 𝑛⊕0,𝑛⊕1,…,𝑛⊕𝑚
. Here, ⊕
 is the bitwise XOR operator.

MEX
 of the sequence of non-negative integers is the smallest non-negative integer that doesn't appear in this sequence. For example, MEX(0,1,2,4)=3
, and MEX(1,2021)=0
.


### ideas
1. n > m, 那么 => 0, 这是因为n ^ 0, ... n ^ m, n ^ (m + 1) 因为n > m ， 所以不会出现 n ^ n 的情况， 不会出现0
2. n <= m, n ^ 0, n ^ 1, ... n ^ n, n ^ (n + 1), n ^ (n + 2) ... n ^ m
3. 