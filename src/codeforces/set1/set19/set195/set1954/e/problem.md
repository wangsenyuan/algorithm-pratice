There are 𝑛
 monsters standing in a row. The 𝑖
-th monster has 𝑎𝑖
 health points.

Every second, you can choose one alive monster and launch a chain lightning at it. The lightning deals 𝑘
 damage to it, and also spreads to the left (towards decreasing 𝑖
) and to the right (towards increasing 𝑖
) to alive monsters, dealing 𝑘
 damage to each. When the lightning reaches a dead monster or the beginning/end of the row, it stops. A monster is considered alive if its health points are strictly greater than 0
.

For example, consider the following scenario: there are three monsters with health equal to [5,2,7]
, and 𝑘=3
. You can kill them all in 4
 seconds:

launch a chain lightning at the 3
-rd monster, then their health values are [2,−1,4]
;
launch a chain lightning at the 1
-st monster, then their health values are [−1,−1,4]
;
launch a chain lightning at the 3
-rd monster, then their health values are [−1,−1,1]
;
launch a chain lightning at the 3
-th monster, then their health values are [−1,−1,−2]
.
For each 𝑘
 from 1
 to max(𝑎1,𝑎2,…,𝑎𝑛)
, calculate the minimum number of seconds it takes to kill all the monsters.

### ideas
1. f(k) 是所需的最短时间， 那么 f(1) >= f(2) >= f(3)... >= f(m)
2. 考虑在给定k的时候，如何操作，以获得最小值？
3. 假设有3个连续的位置，那么比较合理的操作应该是在中间位置？似乎也不是的，貌似可以在相邻地方处理，可以减少浪费
4. 不管怎么处理，每个monster，能够承受的伤害次数=(a[i] + k - 1) / k
5. 每次操作，相当于在i-1处-1，在i+2处+1
6. 可以不考虑那些dead的状态，目标为在最短的时间内，让所有的数字 <= 0
7. 貌似处理逻辑只能是贪心的从头（从尾）开始，如果当前数字 > 0, 则进行一次操作
8. 但是比如 [2, 1, 3]
9. 但是在处理a[1]时，a[3]能够减少的数量 = min(a[1], a[2], a[3])

### solution
Let's solve the problem for a single 𝑘
. We'll start with 𝑘=1
 for simplicity.

The first lightning can be launched at any monster, as it will always spread to all of them. We will continue launching lightnings until a monster dies. When one or more monsters die, the problem breaks down into several independent subproblems, because no lightning will pass through dead monsters. This means that there is no concept of "minimum number of seconds" — the answer does not depend on the choice of monsters to launch the lightnings.

Great, so how do we calculate this answer? The idea is as follows. We will attack the first monster until it dies. This will take 𝑎1
 seconds. We then move on to the second monster. If it has more health than the first one, we need to launch an additional 𝑎2−𝑎1
 lightnings to kill it. Otherwise, it will already be dead. How much damage will the third monster receive in both cases? Let's say it has a lot of health. In the first case, it will receive 𝑎2
 damage, because all the lightnings will reach it. But in the second case, it will also receive 𝑎2
 damage, because the lightnings launched at the first monster after the death of the second one will not reach the third one. This means that we now need to compare the health of the second monster with the third one in the same way. And so on. This means that the 𝑖
-th monster needs to be hit with max(0,𝑎𝑖−𝑎𝑖−1)
 lightnings.

Then the answer for 𝑘=1
 is equal to 𝑎1+∑𝑖=2𝑛max(0,𝑎𝑖−𝑎𝑖−1)
.

How to calculate the answer for any 𝑘
? In fact, the difference is not very significant. It is sufficient to change the health of each monster from 𝑎𝑖
 to ⌈𝑎𝑖𝑘⌉
, and the entire process described earlier will remain the same. Therefore, the answer for any 𝑘
 is equal to ⌈𝑎1𝑘⌉+∑𝑖=2𝑛max(0,⌈𝑎𝑖𝑘⌉−⌈𝑎𝑖−1𝑘⌉)
.

To further optimize this solution, another transformation is needed. Ideally, we would like each 𝑎𝑖
 to contribute to the answer independently of other values. And this can almost be achieved. Notice that the maximum returns 0
 only if 𝑎𝑖<𝑎𝑖−1
 for any 𝑘
, not just for 𝑘=1
. This may require proof, but it is quite obvious.

This means that the coefficient for ⌈𝑎𝑖𝑘⌉
 in the answer depends on two conditions:

it is increased by 1
 if 𝑖=1
 or 𝑎𝑖≥𝑎𝑖−1
;
it is decreased by 1
 if 𝑖<𝑛
 and 𝑎𝑖<𝑎𝑖+1
.
Let's call this coefficient for the 𝑖
-th monster 𝑐𝑖
. Therefore, we need to calculate ∑𝑖=1𝑛𝑐𝑖⋅⌈𝑎𝑖𝑘⌉
.

There are two ways to optimize the solution further.

The first option is to notice that ⌈𝑎𝑖𝑘⌉
 doesn't take a lot of different values for different 𝑘
. More precisely, it is 𝑂(𝑎𝑖‾‾√)
. This can be shown as follows. Consider ⌈𝑎𝑖𝑘⌉=𝑥
. Either 𝑘≤𝑎𝑖‾‾√
, or 𝑥≤𝑎𝑖‾‾√
. Therefore, 𝑥
 takes no more than 2𝑎𝑖‾‾√
 different values.

Then the solution can be implemented as follows. For each 𝑎𝑖
, we will identify all possible values that the rounding result takes. For each of them, we will find the range of 𝑘
 for which the result is equal to that. And we will add the contribution of the 𝑖
-th monster within this range of values to the result. This can be done using a difference array to achieve a complexity of 𝑂(𝑛⋅𝐴‾‾√)
.

The second option is a bit smarter. Let's take another look at the formula for calculating the answer for a fixed 𝑘
: ∑𝑖=1𝑛𝑐𝑖⋅⌈𝑎𝑖𝑘⌉
. Let's group the terms by equal values of ⌈𝑎𝑖𝑘⌉
. What do they look like? Numbers from 1
 to 𝑘
 give the value 1
. Numbers from 𝑘+1
 to 2𝑘
 give the value 2
, and so on. This means that for a certain 𝑘
, there are 𝐴𝑘
 segments, on each of which we need to calculate the sum of 𝑐𝑖
 for those 𝑖
 for which 𝑎𝑖
 fall into this segment. The total number of segments for all 𝑘
 is 𝑂(𝐴log𝐴)
. The complexity of the solution will then be 𝑂(𝑛+𝐴log𝐴)
.