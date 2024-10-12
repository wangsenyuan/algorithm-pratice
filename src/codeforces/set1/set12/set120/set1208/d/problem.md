An array of integers 𝑝1,𝑝2,…,𝑝𝑛
 is called a permutation if it contains each number from 1
 to 𝑛
 exactly once. For example, the following arrays are permutations: [3,1,2],[1],[1,2,3,4,5]
 and [4,3,1,2]
. The following arrays are not permutations: [2],[1,1],[2,3,4]
.

There is a hidden permutation of length 𝑛
.

For each index 𝑖
, you are given 𝑠𝑖
, which equals to the sum of all 𝑝𝑗
 such that 𝑗<𝑖
 and 𝑝𝑗<𝑝𝑖
. In other words, 𝑠𝑖
 is the sum of elements before the 𝑖
-th element that are smaller than the 𝑖
-th element.

Your task is to restore the permutation.

### ideas
1. s[1] = 0 (这个地方目前没有可以用的信息)
2. 如果 s[i] = 0 且 i > 1, 那么 p[i] = 1 (然后可以它右边的s[?] - 1)
3. 考虑 6, 5, 4, 3, 2, 1 
4. 所以始终找最靠右边的s[i] = 0 的部分, 就可以了
5. range update / query
6. 要没有更简单的结构？
7. 题解和我的做法是一样的。那可能没有更优的结构了吧