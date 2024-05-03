There are n shovels in Polycarp's shop. The i-th shovel costs i burles, that is, the first shovel costs 1 burle, the
second shovel costs 2 burles, the third shovel costs 3 burles, and so on. Polycarps wants to sell shovels in pairs.

Visitors are more likely to buy a pair of shovels if their total cost ends with several 9s. Because of this, Polycarp
wants to choose a pair of shovels to sell in such a way that the sum of their costs ends with maximum possible number of
nines. For example, if he chooses shovels with costs 12345 and 37454, their total cost is 49799, it ends with two nines.

You are to compute the number of pairs of shovels such that their total cost ends with maximum possible number of nines.
Two pairs are considered different if there is a shovel presented in one pair, but not in the other.

### ideas

1. 假设选择的是a, b, 那么a + b = ???999
2. 首先就是,小于n的最大的连续的9,总是可以达到的
3. 假设num有3位数，那么至少可以得到99
4. 如果它有4位数，那么就可以得到999
5. 依次类推，唯一变化的部分是最高位