Consider a system of n water taps all pouring water into the same container. The i-th water tap can be set to deliver
any amount of water from 0 to ai ml per second (this amount may be a real number). The water delivered by i-th tap has
temperature ti.

If for every you set i-th tap to deliver exactly xi ml of water per second, then the resulting temperature of water will
be  (if , then to avoid division by zero we state that the resulting water temperature is 0).

You have to set all the water taps in such a way that the resulting temperature is exactly T. What is the maximum amount
of water you may get per second if its temperature has to be T?

### solution

1. let X = sum of xi
2. let Y = sum of xi * ti
3. Y = X * T
4. 这里T是固定的， 问最大的X是多少
5. 没啥思路
6. 假设最简单的一种情况，所有它t[i] = T, 那么此时所有的taps，加入的量，不影响温度，所以可以取a[i]
7. 如果所有的t[i] < T => 0
8. 所有的t[i] > T => 0
9. 假设最优的答案是W, 那么 W - delta 是否也是ok的？
10. 是ok的，因为在W的情况下成立，意味着有部分 > T, 部分 < T, 部分 = T,
11. 如果还有 = T的部分，就去掉delta部分的这种
12. 或者去掉混合后的delta部分
13. 所以可以二分
14. 对于给定的W，先将超过T的放在一起，然后混合出W的，如果温度 = T正好
15. 如果一开始就使用高温的，那么通过加入低温的，temp 就会走低
16. 然后不断加入低温
17. 假设目前的温度是tx, 且总容量是 wx,
18. 那么 wx * tx + b[i] * t[i] / (wx + b[i])
    19. 如果全部放入后 tx > T, 但是 ty < T， 且 wx >= W => good