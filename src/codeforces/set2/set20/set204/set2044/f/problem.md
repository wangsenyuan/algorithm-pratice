For an arbitrary grid, Robot defines its beauty to be the sum of elements in the grid.

Robot gives you an array 𝑎
 of length 𝑛
 and an array 𝑏
 of length 𝑚
. You construct a 𝑛
 by 𝑚
 grid 𝑀
 such that 𝑀𝑖,𝑗=𝑎𝑖⋅𝑏𝑗
 for all 1≤𝑖≤𝑛
 and 1≤𝑗≤𝑚
.

Then, Robot gives you 𝑞
 queries, each consisting of a single integer 𝑥
. For each query, determine whether or not it is possible to perform the following operation exactly once so that 𝑀
 has a beauty of 𝑥
:

Choose integers 𝑟
 and 𝑐
 such that 1≤𝑟≤𝑛
 and 1≤𝑐≤𝑚
Set 𝑀𝑖,𝑗
 to be 0
 for all ordered pairs (𝑖,𝑗)
 such that 𝑖=𝑟
, 𝑗=𝑐
, or both.
Note that queries are not persistent, meaning that you do not actually set any elements to 0
 in the process — you are only required to output if it is possible to find 𝑟
 and 𝑐
 such that if the above operation is performed, the beauty of the grid will be 𝑥
. Also, note that you must perform the operation for each query, even if the beauty of the original grid is already 𝑥
.



### ideas
1. s = (sum of a) * (sum of b)
2. 假设选中了这样的一个(r, c)
3. s - a[r] * s2 - b[c] * s1 + a[r] * b[c] = x
4. 排序后，对结果没有影响
5. 假设a, b都排好序了
6. s - a[r] * s2 - b[c] * (s1 - a[r])
7. s = s1 * s2
