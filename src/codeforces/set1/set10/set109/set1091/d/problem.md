Let 𝑛
be an integer. Consider all permutations on integers 1
to 𝑛
in lexicographic order, and concatenate them into one big sequence 𝑝
. For example, if 𝑛=3
, then 𝑝=[1,2,3,1,3,2,2,1,3,2,3,1,3,1,2,3,2,1]
. The length of this sequence will be 𝑛⋅𝑛!
.

Let 1≤𝑖≤𝑗≤𝑛⋅𝑛!
be a pair of indices. We call the sequence (𝑝𝑖,𝑝𝑖+1,…,𝑝𝑗−1,𝑝𝑗)
a subarray of 𝑝
. Its length is defined as the number of its elements, i.e., 𝑗−𝑖+1
. Its sum is the sum of all its elements, i.e., ∑𝑗𝑘=𝑖𝑝𝑘
.

You are given 𝑛
. Find the number of subarrays of 𝑝
of length 𝑛
having sum 𝑛(𝑛+1)2
. Since this number may be large, output it modulo 998244353
(a prime number).

### ideas

1. 是一个很费脑的题目
2. 考虑上面1，2，3的例子， sum n * (n + 1) / 2 这个就是数组p的sum = 6
3. 每个独立的段，也就是从0, n, 2 * n ... i * n 的n长的都是满足条件的， 这个数量是 n!
4. 然后剩下的是，在覆盖两段的
5. 比如上面的 2,3,1
6. 这个有个简单的观察，就是对于后面一段是1开头的部分，那么2，3，1，或者 3,2,1
7. 也就是前面的2,3 的 个数的阶乘 (n-1)!, 但是这里有个问题，就是3, 2, 1 不会出现
8. n * (n - 1) ! - (n-1)! (后面一项表示的是每个block开始的部分)
9. 1，2，3，4，1，2，4，3，1，3，2，4，1，4，2，3，1，4，3，2
10. 有出现一个新的情况，上面 2,3,1,4
11. 以1开头的，然后1，2开头，然后1，3，然后1，4， 然后1，2，3，1，2，4， 1，3，4， 然后 1，2，3，4
12. 正难则反吗？
13. 以1开头的有(n-1)!,
14. 除了第一个和最后一个，剩下的 (n-1)! - 1 和1就可以组成一个满足条件的
15. 然后是1，2开头的，（n-2)! - 1
16. 1, 2, 3, 4, 1, 2, 4, 3, 1, ? 也是满足条件的
17. n! + n * ((n-1)! - 1) + n * (n - 1) * ((n - 2)! - 1) ..
18. 对于4， 24 + 4 * (6 - 1) + 4 * 3 * (2 - 1) = 24 + 20 + 12