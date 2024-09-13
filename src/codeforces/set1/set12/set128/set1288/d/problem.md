You are given 𝑛
 arrays 𝑎1
, 𝑎2
, ..., 𝑎𝑛
; each array consists of exactly 𝑚
 integers. We denote the 𝑦
-th element of the 𝑥
-th array as 𝑎𝑥,𝑦
.

You have to choose two arrays 𝑎𝑖
 and 𝑎𝑗
 (1≤𝑖,𝑗≤𝑛
, it is possible that 𝑖=𝑗
). After that, you will obtain a new array 𝑏
 consisting of 𝑚
 integers, such that for every 𝑘∈[1,𝑚]
 𝑏𝑘=max(𝑎𝑖,𝑘,𝑎𝑗,𝑘)
.

Your goal is to choose 𝑖
 and 𝑗
 so that the value of min𝑘=1𝑚𝑏𝑘
 is maximum possible.

 ### ideas
 1. 这里m比较小 (8)，
 2. 一头雾水
 3. 在选中j的情况下，如何找到i使得，b[k]最大呢？
 4. 如果设定，a[i][k] 和 a[j][k]的关系（根据state[k] = 0/1), 是否有用
 5. 假设现在b[k] = a[j][k], 也就是a[i][k] <= b[j][k]
 6. 如果 b[k] = a[i][k], 也就是 a[i][k] >= b[j][k]
 7. 也就是说需要确定在固定j，且满足state时的最大的值是什么
 8. 感觉不大行
 9. 在确定a[i], a[j]时，f(i, j) = min(max(a[i][k], b[j][k]))
 10. 假设结果在x时成立，显然在x-1时成立
 11. 假设现在希望结果时x能够成立，那么必须选择这样的(i, j), 它们对应位置至少有一个数 >= x
 12. 那就好说了
 13. 按照 大于，或者小于 x 表示状态，然后去检查，对于状态x,是否存在它的补的超集