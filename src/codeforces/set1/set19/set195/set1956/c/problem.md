The magical girl Nene has an 𝑛×𝑛
 matrix 𝑎
 filled with zeroes. The 𝑗
-th element of the 𝑖
-th row of matrix 𝑎
 is denoted as 𝑎𝑖,𝑗
.

She can perform operations of the following two types with this matrix:

Type 1
 operation: choose an integer 𝑖
 between 1
 and 𝑛
 and a permutation 𝑝1,𝑝2,…,𝑝𝑛
 of integers from 1
 to 𝑛
. Assign 𝑎𝑖,𝑗:=𝑝𝑗
 for all 1≤𝑗≤𝑛
 simultaneously.
Type 2
 operation: choose an integer 𝑖
 between 1
 and 𝑛
 and a permutation 𝑝1,𝑝2,…,𝑝𝑛
 of integers from 1
 to 𝑛
. Assign 𝑎𝑗,𝑖:=𝑝𝑗
 for all 1≤𝑗≤𝑛
 simultaneously.
Nene wants to maximize the sum of all the numbers in the matrix ∑𝑖=1𝑛∑𝑗=1𝑛𝑎𝑖,𝑗
. She asks you to find the way to perform the operations so that this sum is maximized. As she doesn't want to make too many operations, you should provide a solution with no more than 2𝑛
 operations.

A permutation of length 𝑛
 is an array consisting of 𝑛
 distinct integers from 1
 to 𝑛
 in arbitrary order. For example, [2,3,1,5,4]
 is a permutation, but [1,2,2]
 is not a permutation (2
 appears twice in the array), and [1,3,4]
 is also not a permutation (𝑛=3
 but there is 4
 in the array).

### ideas
1. 如果 m < n， 假设m = 2, 是否应该操作1和操作2都执行
2. 假设都执行了，那么它们必然在某个(i, j)处相交，单次操作的收益时 (1 + 2 + .. + n)
3. 但是有重叠的格子，造成(i, j)处的浪费了（比如操作1是x，操作2是y）那么x就丢失了。
4. 所以，如果 m <= n只应该进行一种操作
5. m > n, 此时最好把最小的那些给替换掉。 n个1给替换掉，2 * n, 然后是 3 * n
6. 假设 i * n < n * (n + 1) / 2， 也就是i < (n + 1) / 2 的时候，
7. 那么此时收益都是正的。但是如果i >= (n + 1) / 2时，此时收益就是负的
8. 反而要往前处理了。
9. 但是这样一定能得到最大的sum吗？
10. 