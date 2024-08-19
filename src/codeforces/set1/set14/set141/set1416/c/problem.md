You are given an array 𝑎
 consisting of 𝑛
 non-negative integers. You have to choose a non-negative integer 𝑥
 and form a new array 𝑏
 of size 𝑛
 according to the following rule: for all 𝑖
 from 1
 to 𝑛
, 𝑏𝑖=𝑎𝑖⊕𝑥
 (⊕
 denotes the operation bitwise XOR).

An inversion in the 𝑏
 array is a pair of integers 𝑖
 and 𝑗
 such that 1≤𝑖<𝑗≤𝑛
 and 𝑏𝑖>𝑏𝑗
.

You should choose 𝑥
 in such a way that the number of inversions in 𝑏
 is minimized. If there are several options for 𝑥
 — output the smallest one.

 ### ideas
 1. b[i] = a[i] ^ x
 2. b[i+1] = a[i+1] ^ x
 3. b[i] ^ b[i+1] = a[i] ^ a[i+1]
 4. 希望b[i]尽可能的递增
 5. 有没有可能得到inverse = 0 的结果？
 6. 这时候，希望b[n]最大
 7. 假设前面有个a[i] > a[n], a[i][d] = 1, and a[n][d] = 0
 8. 那么只要x[d] = 1, 就可以让b[i] < b[n],
 9. 但是 x[d] = 1, 有可能让 a[i] > a[n-1] 吗？  