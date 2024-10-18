Alan decided to get in shape for the summer, so he created a precise workout plan to follow. His plan is to go to a different gym every day during the next N days and lift 𝑋[𝑖]
 grams on day 𝑖
. In order to improve his workout performance at the gym, he can buy exactly one pre-workout drink at the gym he is currently in and it will improve his performance by 𝐴
 grams permanently and immediately. In different gyms these pre-workout drinks can cost different amounts 𝐶[𝑖]
 because of the taste and the gym's location but its permanent workout gains are the same. Before the first day of starting his workout plan, Alan knows he can lift a maximum of 𝐾
 grams. Help Alan spend a minimum total amount of money in order to reach his workout plan. If there is no way for him to complete his workout plan successfully output −1
.

### ideas
1. 反悔dp吗？
2. 假设现在储备了一堆可以提升能力的饮料
3. 如果当前的gym要求太高了，加上储备的也不够，=》-1
4. 否则的话，就消耗储备（或者不消耗），至少到达目前的要求
5. 但是消耗哪些呢？是价格最便宜的吗？
6. 还是单位价格最便宜的？
7. 价格最便宜的，因为A是固定的
8. 