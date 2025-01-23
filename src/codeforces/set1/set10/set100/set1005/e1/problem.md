You are given a permutation 𝑝1,𝑝2,…,𝑝𝑛
. A permutation of length 𝑛
 is a sequence such that each integer between 1
 and 𝑛
 occurs exactly once in the sequence.

Find the number of pairs of indices (𝑙,𝑟)
 (1≤𝑙≤𝑟≤𝑛
) such that the value of the median of 𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟
 is exactly the given number 𝑚
.

The median of a sequence is the value of the element which is in the middle of the sequence after sorting it in non-decreasing order. If the length of the sequence is even, the left of two middle elements is used.

For example, if 𝑎=[4,2,7,5]
 then its median is 4
 since after sorting the sequence, it will look like [2,4,5,7]
 and the left of two middle elements is equal to 4
. The median of [7,1,2,9,6]
 equals 6
 since after sorting, the value 6
 will be in the middle of the sequence.

Write a program to find the number of pairs of indices (𝑙,𝑟)
 (1≤𝑙≤𝑟≤𝑛
) such that the value of the median of 𝑝𝑙,𝑝𝑙+1,…,𝑝𝑟
 is exactly the given number 𝑚
.

### ideas
1. 如果m是在奇数长度的序列中，那么就是要选者x个<m 的数， >m个数，且这个2 * x + 1长度的区间，要包含m
2. 在选定l的时候， r是唯一确定的
3. 假设在l...i (m 所在的位置)中有x个数小于m，那么就有 y = (i - l)个数 > m
4. x + u = y + v (u+v是右边的长度)
5. u - v = y - x
6. 好搞了