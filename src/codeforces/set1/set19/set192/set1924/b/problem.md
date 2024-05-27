There are 𝑛
 points numbered 1
 to 𝑛
 on a straight line. Initially, there are 𝑚
 harbours. The 𝑖
-th harbour is at point 𝑋𝑖
 and has a value 𝑉𝑖
. It is guaranteed that there are harbours at the points 1
 and 𝑛
. There is exactly one ship on each of the 𝑛
 points. The cost of moving a ship from its current location to the next harbour is the product of the value of the nearest harbour to its left and the distance from the nearest harbour to its right. Specifically, if a ship is already at a harbour, the cost of moving it to the next harbour is 0
.

Additionally, there are 𝑞
 queries, each of which is either of the following 2
 types:

1
 𝑥
 𝑣
 — Add a harbour at point 𝑥
 with value 𝑣
. It is guaranteed that before adding the harbour, there is no harbour at point 𝑥
.
2
 𝑙
 𝑟
 — Print the sum of the cost of moving all ships at points from 𝑙
 to 𝑟
 to their next harbours. Note that you just need to calculate the cost of moving the ships but not actually move them.

 ### ideas
 1. 一个ship移动的cost，在港口确定的情况下，也是确定的
 2. 如果这个ship已经在港口了，那么移动它到下一个港口（或者更后面的港口)cost = 0, 也就是说一旦ship进入港口，它就不再移动了
 3. 如果它还不在港口，那么cost = V[i] * (X[i+1] - pos)
 4. pos = 1, 2, 3, ... n
 5. 也就是每个pos对移动的贡献是确定的 V[i] * X[i+1] - V[i] * pos
 6. 考虑没有query 1的情况下，只是简单的区间查询
 7. 考虑query 1的情况，假设在位置i增加了一个港口，那么对于假设 X[j] < i, X[j+1] > i 
 8. 在X[j]前面的ship没有影响， X[j+1]后面的也没有影响
 9. 只有它们中间的会有影响
10. 考虑位置 i-1, i, i + 1
11. 在之前，i-1的cost = V[j] * (X[j+1] - pos)
12.       i  的 cost = V[j] * (X[j+1] - pos)
13.       i+1 的cost = V[j] * (X[j+1] - pos)
14.   之后  i-1的cost = V[j] * (i - pos) 
15.        i   的 cost = 0
16.        i+1的从身体 v * (X[j+1] - (i + 1))
17.   在 X[j]+1 到 i的变化是 -V[j] * X[j+1] + i * V[j] = V[j] * (-X[j+1] + i)
18.   在 i+1后面的变化是 （X[j+1] - pos) * (-V[j] + v)
19.   对于X[j] + 1 到i，是一个固定的变化 V[j] * (-X[j+1] + i) （全部加上这个）
20.   对于i+1到X[j+1] - 1，这个要怎么更新呢？
21.   (X[j+1]) * (v - V[j]) 这个是固定值
22.   -pos * (v - V[j]) 这个每个位置不同， (但是 v-V[j])是固定的
23.   -(l + l+1 + ... r) * (v - V[j])
24.   (l + r) * (r - l + 1) / 2 * (v - V[j])