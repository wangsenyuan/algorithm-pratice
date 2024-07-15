The student council is preparing for the relay race at the sports festival.

The council consists of 𝑛
 members. They will run one after the other in the race, the speed of member 𝑖
 is 𝑠𝑖
. The discrepancy 𝑑𝑖
 of the 𝑖
-th stage is the difference between the maximum and the minimum running speed among the first 𝑖
 members who ran. Formally, if 𝑎𝑖
 denotes the speed of the 𝑖
-th member who participated in the race, then 𝑑𝑖=max(𝑎1,𝑎2,…,𝑎𝑖)−min(𝑎1,𝑎2,…,𝑎𝑖)
.

You want to minimize the sum of the discrepancies 𝑑1+𝑑2+⋯+𝑑𝑛
. To do this, you are allowed to change the order in which the members run. What is the minimum possible sum that can be achieved?

### ideas
1. d1 = 0 always
2. di = max(ai) - min(ai)
3. 假设第一个是最大值，到了某个时候， a[0] - a[i], 但是如果这个时候， a[0]不是最大值，那么显然会得到更好的结果
4. 所以，可以看出来，就是最大值似乎在最后阶段会更优，这个对最小值也是成立的
5. 所以，可以大体得到这样一个结论，就是a[0], a[n-1]是最后两个
6. 然后是a[1], a[n-2], ...
7. 但是哪个更好呢？似乎是确定的。先用a[0], 那么 d[n-1] = a[n-2] - a[0] 或者 d[n-1] = a[n-1] - a[1] 
8. [104 943872923 6589 889921234 1000000000 69] = 2833800505
9. 好像又不成立
10. 假设最大值在x前面已经出现了，现在出现了y，y是第二大的数，那么用y替换x似乎是更优的
11. 在x之前的，没有影响，但是x之后的，x始终是最大值，但是用y替换后，就会出现一部分diff
12. dp[l..r]是先处理[l...r]的最优解 dp[l-1...r] = dp[l..r] + a[r] - a[l-1]