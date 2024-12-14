Let us define a magic grid to be a square matrix of integers of size 𝑛×𝑛
, satisfying the following conditions.

All integers from 0
 to (𝑛2−1)
 inclusive appear in the matrix exactly once.
Bitwise XOR of all elements in a row or a column must be the same for each row and column.
You are given an integer 𝑛
 which is a multiple of 4
. Construct a magic grid of size 𝑛×𝑛
.

### ideas
1. 00  01  10, 11 这些放在一行里面，是可以xor 掉的
2. 感觉是4个一组
3. 0, 1, 2, 3 
4. 4, 5, 6, 7 （100, 101, 110, 111) 
5. 但是问题出在col方向 （需要把 5 和 4交换）
6. 也就是在一个4 * 4个格子中，要能都有那些就可以了
7. 0, 1, 2, 3
8. 1, 0, 3, 2
9. 2, 3, 0, 1
10. 3, 2, 1, 0