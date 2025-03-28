In the probability theory the following paradox called Benford's law is known: "In many lists of random numbers taken from real sources, numbers starting with digit 1 occur much more often than numbers starting with any other digit" (that's the simplest form of the law).

Having read about it on Codeforces, the Hedgehog got intrigued by the statement and wishes to thoroughly explore it. He finds the following similar problem interesting in particular: there are N random variables, the i-th of which can take any integer value from some segment [Li;Ri] (all numbers from this segment are equiprobable). It means that the value of the i-th quantity can be equal to any integer number from a given interval [Li;Ri] with probability 1 / (Ri - Li + 1).

The Hedgehog wants to know the probability of the event that the first digits of at least K% of those values will be equal to one. In other words, let us consider some set of fixed values of these random variables and leave only the first digit (the MSD — most significant digit) of each value. Then let's count how many times the digit 1 is encountered and if it is encountered in at least K per cent of those N values, than such set of values will be called a good one. You have to find the probability that a set of values of the given random variables will be a good one.


### ideas
1. 先考虑在一个区间内，随机取一个数字，它的第一个字符是1的出现概率
2. calc(num) 表示在区间0...num中出现首位为1的数字的个数
3. (calc(r) - calc(l-1)) / (r - l + 1) = 首位是1的数字的概率
4. 这个只是单次选择一个数字value，得到good数字的概率
5. 那么选择n个数字， 假设都是1的概率 = prod(fi)
6. 如果是k%, 如果恰好是100个区间，那么就是其中至少k个区间是1的概率
7. dp[i][j] 表示前i个区间，其中j个区间为1的概率
8. dp[i][j] = dp[i-1][j] * (1 - f[i]) + dp[i-1][j-1] * f[i]
9. dp[i][j] 如果 j * 100 >= k * i, sum起来就可以了