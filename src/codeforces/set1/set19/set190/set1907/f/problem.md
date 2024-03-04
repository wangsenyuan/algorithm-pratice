Given an array of integers 𝑎1,𝑎2,…,𝑎𝑛
. You can make two types of operations with this array:

Shift: move the last element of array to the first place, and shift all other elements to the right, so you get the
array 𝑎𝑛,𝑎1,𝑎2,…,𝑎𝑛−1
.
Reverse: reverse the whole array, so you get the array 𝑎𝑛,𝑎𝑛−1,…,𝑎1
.
Your task is to sort the array in non-decreasing order using the minimal number of operations, or say that it is
impossible.

### thoughts

1. 考虑最后一步，要们已经sorted
2. 要么shift， 满足 2341这样的模式
3. 要么reverse, 满足 4321这样的模式
4. 如果把