Fox loves permutations! She came up with the following problem and asked Cat to solve it:

You are given an even positive integer 𝑛
 and a permutation†
 𝑝
 of length 𝑛
.

The score of another permutation 𝑞
 of length 𝑛
 is the number of local maximums in the array 𝑎
 of length 𝑛
, where 𝑎𝑖=𝑝𝑖+𝑞𝑖
 for all 𝑖
 (1≤𝑖≤𝑛
). In other words, the score of 𝑞
 is the number of 𝑖
 such that 1<𝑖<𝑛
 (note the strict inequalities), 𝑎𝑖−1<𝑎𝑖
, and 𝑎𝑖>𝑎𝑖+1
 (once again, note the strict inequalities).

Find the permutation 𝑞
 that achieves the maximum score for given 𝑛
 and 𝑝
. If there exist multiple such permutations, you can pick any of them.

†
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
1. n is even, 这个信息估计是有用的
2. 在一个数组中，最大的score = (n-1) / 2
3. 比如 [1, 2, 3, 4] => [1, 3, 2, 4] 
4. 是不是一定能够得到这样的score呢？
5. [1, 2, 3, 4, 5, 6] + [5, 6, 2, 3, 4, 1]
6. [6, 8, 5, 7, 9, 7]
7. 在过程中，肯定是可以保证 a[i] != a[i+1], 也就是不会出现相等的两个数
8. 如果出现了 p[i] != p[i+1], 但是 a[i] = a[i+1], q[i] = a[i] - p[i], q[i+1] = a[i+1] - p[i+1]
9. 那么只要把 p[i]和 p[i+1]交换顺序，就能保证它们不相等
10. 先保证峰顶，且使用后面的一半数字来保证