You are given a 2×𝑛
 grid filled with zeros and ones. Let the number at the intersection of the 𝑖
-th row and the 𝑗
-th column be 𝑎𝑖𝑗
.

There is a grasshopper at the top-left cell (1,1)
 that can only jump one cell right or downwards. It wants to reach the bottom-right cell (2,𝑛)
. Consider the binary string of length 𝑛+1
 consisting of numbers written in cells of the path without changing their order.

Your goal is to:

Find the lexicographically smallest†
 string you can attain by choosing any available path;
Find the number of paths that yield this lexicographically smallest string.
†
 If two strings 𝑠
 and 𝑡
 have the same length, then 𝑠
 is lexicographically smaller than 𝑡
 if and only if in the first position where 𝑠
 and 𝑡
 differ, the string 𝑠
 has a smaller element than the corresponding element in 𝑡
.

### ideas
1. 任何一个时刻，要么在第一行，要么在第二行，如果到了第二行，那么就一直要到末尾
2. 分层，第一层是(1, 1), 第二层是(2, 1), (1, 2)， 且在每一层都取前面的，激活下一层