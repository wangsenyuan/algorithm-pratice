You are given an array 𝑎
consisting of 𝑛
integers. You can perform the following operations with it:

Choose some positions 𝑖
and 𝑗
(1≤𝑖,𝑗≤𝑛,𝑖≠𝑗
), write the value of 𝑎𝑖⋅𝑎𝑗
into the 𝑗
-th cell and remove the number from the 𝑖
-th cell;
Choose some position 𝑖
and remove the number from the 𝑖
-th cell (this operation can be performed no more than once and at any point of time, not necessarily in the beginning).
The number of elements decreases by one after each operation. However, the indexing of positions stays the same. Deleted
numbers can't be used in the later operations.

Your task is to perform exactly 𝑛−1
operations with the array in such a way that the only number that remains in the array is maximum possible. This number
can be rather large, so instead of printing it you need to print any sequence of operations which leads to this maximum
number. Read the output format to understand what exactly you need to print.

### ideas

1. 先统计个数，如果0的个数超过1，那么随便搞，因为最终的结果都是0
2. 然后看是否能是正数，所以要知道负数的个数
3. 如果负数的个数是奇数个
4. 在有0的情况下 =》 0
5. 在没有0的情况下，删掉一个负数
6. 如果要删掉负数，删掉最大的负数