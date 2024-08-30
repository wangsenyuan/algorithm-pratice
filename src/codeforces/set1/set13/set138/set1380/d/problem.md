There are 𝑛
 warriors in a row. The power of the 𝑖
-th warrior is 𝑎𝑖
. All powers are pairwise distinct.

You have two types of spells which you may cast:

Fireball: you spend 𝑥
 mana and destroy exactly 𝑘
 consecutive warriors;
Berserk: you spend 𝑦
 mana, choose two consecutive warriors, and the warrior with greater power destroys the warrior with smaller power.
For example, let the powers of warriors be [2,3,7,8,11,5,4]
, and 𝑘=3
. If you cast Berserk on warriors with powers 8
 and 11
, the resulting sequence of powers becomes [2,3,7,11,5,4]
. Then, for example, if you cast Fireball on consecutive warriors with powers [7,11,5]
, the resulting sequence of powers becomes [2,3,4]
.

You want to turn the current sequence of warriors powers 𝑎1,𝑎2,…,𝑎𝑛
 into 𝑏1,𝑏2,…,𝑏𝑚
. Calculate the minimum amount of mana you need to spend on it.

#ideas
1. 考虑只能用方式2，那么此时a[i] = b[1], 那么a[i]必须是前缀[1...i]的最大值
2. let p[i]是a中位b[i]保留的下标
3. 在只使用操作2的情况下，p[i] ~ p[i+1] 中间的，都必须小于a[p[i]]或者 a[p[i+1]]
4. 这个可以用贪心的办法去判断。且结果 = (n - m) * y
5. 那么操作1要咋搞呢？
6. 一类是在p[i] ~ p[i+1]中间，如果使用操作1，更有利（k * y >= x)
7. 那么就使用操作1，还有一种是，如果 p[i] ~ p[i+1]中间出现了超过两端的数，那么就必须用操作1
8. 但是这里还有第三种情况，仍然贪心的选择p[i+1]是否是最优的？
9. 感觉仍然是最优的。贪心的算法，都比较难证明
10. 