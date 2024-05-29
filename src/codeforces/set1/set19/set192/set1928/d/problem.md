Once, the people, elves, dwarves, and other inhabitants of Middle-earth gathered to reclaim the treasures stolen from them by Smaug. In the name of this great goal, they rallied around the powerful elf Timothy and began to plan the overthrow of the ruler of the Lonely Mountain.

The army of Middle-earth inhabitants will consist of several squads. It is known that each pair of creatures of the same race, which are in different squads, adds 𝑏
 units to the total strength of the army. But since it will be difficult for Timothy to lead an army consisting of a large number of squads, the total strength of an army consisting of 𝑘
 squads is reduced by (𝑘−1)⋅𝑥
 units. Note that the army always consists of at least one squad.

It is known that there are 𝑛
 races in Middle-earth, and the number of creatures of the 𝑖
-th race is equal to 𝑐𝑖
. Help the inhabitants of Middle-earth determine the maximum strength of the army they can assemble.


### ideas
1. 先理解这个题目
2. 假设最后有k个squad, 第i个种族的数量是a[i]， 
3. 如果它们被这样分布x[1], x[2], x[3], ... x[k]
4. 那么第i个种族的贡献= x[2] * x[1] + (x1 + x2) * x[3] + (x1 + x2 + x3) * x4, ....
5. 这个式子直观的感受，应该是分的越均匀越好
6. 假设a[i] = k, 那么每个squad分一个
7. (k - 1) * k 如果有一个分2个，其他的分1个，但有一个分0个
8. 2 * (k - 2) + (k - 2) * (k - 1)
9. (k + 1) * (k - 2) = k * k - k - 2 < k * k - k = (k - 1) * k
10. (k - 1) * k = k * k - k
11. 在给定k的情况下，
12. 迭代k?
13. 当 c[i] <= k 的时候，它的贡献是确定的，c[i] * (c[i] - 1) 
14. 当 c[i] > k的时候， 它的贡献 = x = c[i] / k, 其中 r = c[i] % k 个 = x + 1
15. x * (c[i] - x) * (k - r) + (x + 1) * (c[i] - (x + 1)) * r
16. 不大行
17. 这个式子里面k即在分母上(x = c[i] / k) 又在分子上，而且还有取整的问题
18. 但是可能可以证明的是，它在某个靠近中部取到最大值，而在两边取到最小值
19. 所以可以用多路查询