There are 𝑛
 points on a coordinate axis 𝑂𝑋
. The 𝑖
-th point is located at the integer point 𝑥𝑖
 and has a speed 𝑣𝑖
. It is guaranteed that no two points occupy the same coordinate. All 𝑛
 points move with the constant speed, the coordinate of the 𝑖
-th point at the moment 𝑡
 (𝑡
 can be non-integer) is calculated as 𝑥𝑖+𝑡⋅𝑣𝑖
.

Consider two points 𝑖
 and 𝑗
. Let 𝑑(𝑖,𝑗)
 be the minimum possible distance between these two points over any possible moments of time (even non-integer). It means that if two points 𝑖
 and 𝑗
 coincide at some moment, the value 𝑑(𝑖,𝑗)
 will be 0
.

Your task is to calculate the value ∑1≤𝑖<𝑗≤𝑛
 𝑑(𝑖,𝑗)
 (the sum of minimum distances over all pairs of points).

 ### ideas
 1. 显然不会存在碰撞的问题
 2. 速度相反的那些，他们的距离，最近肯定是0，可以忽略掉（相同符号的速度不用计算）
 3. 只考虑速度为正的情况，如果 x[i] < x[j], 但是 v[i] > v[j], 那么他们的最小距离肯定是0， 也可以忽略掉
 4. 那么就剩下 x[i] < x[j], 但是 v[i] <= v[j]的情况
 5. 他们的最短距离 = x[j] - x[i]
 6. 那就简单了
 7. 按照x升序处理， 对于j来说，要计算出所有v[i] <= v[j] 的 x的和（及数量）
 8. cnt * x[j] - sum
 9. 前面的分析有个很大的问题，假设i的速度v[i] <= v[j]
 10. 且x[i] <= x[j], 那么v[i] < 0 也是有贡献的
 11. 如果 v[i] > v[j], 那么它们肯定会相遇的
 12. 但是如果 v[i] < 0, v[j] < 0， 但是 abs(v[i]) < abs(v[j]), 也是会相遇的，