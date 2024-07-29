You are given two binary strings 𝑎
 and 𝑏
 of the same length. You can perform the following two operations on the string 𝑎
:

Swap any two bits at indices 𝑖
 and 𝑗
 respectively (1≤𝑖,𝑗≤𝑛
), the cost of this operation is |𝑖−𝑗|
, that is, the absolute difference between 𝑖
 and 𝑗
.
Select any arbitrary index 𝑖
 (1≤𝑖≤𝑛
) and flip (change 0
 to 1
 or 1
 to 0
) the bit at this index. The cost of this operation is 1
.
Find the minimum cost to make the string 𝑎
 equal to 𝑏
. It is not allowed to modify string 𝑏
.

### ideas
1. 操作1相当于是交换临近的位置
2. 假设要交换(i, j)且它们的距离大于1了，比如(0, 2)，那么还不如使用操作2呢， 直接flip两次
3. 所以只需要考虑临近的两个是否，通过swap后，可以使得a/b相同，然后把剩余位置上的进行flip操作