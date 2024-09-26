Once upon a time in a kingdom far, far away… Okay, let’s start at the point where Ivan the Fool met Gorynych the Dragon. Ivan took out his magic sword and the battle began. First Gorynych had h heads and t tails. With each strike of the sword Ivan can either cut off several heads (from 1 to n, but not more than Gorynych has at the moment), or several tails (from 1 to m, but not more than Gorynych has at the moment). At the same time, horrible though it seems, Gorynych the Dragon can also grow new heads and tails. And the number of growing heads and tails is determined uniquely by the number of heads or tails cut by the current strike. When the total number of heads and tails exceeds R, Gorynych the Dragon strikes its final blow and destroys Ivan the Fool. That’s why Ivan aims to cut off all the dragon’s heads and tails as quickly as possible and win. The events can also develop in a third way: neither of the opponents can win over the other one and they will continue fighting forever.

The tale goes like this; easy to say, hard to do. Your task is to write a program that will determine the battle’s outcome. Consider that Ivan strikes consecutively. After each blow Gorynych grows a number of new heads and tails depending on the number of cut ones. Gorynych the Dragon is defeated if after the blow he loses all his heads and tails and can’t grow new ones. Ivan fights in the optimal way (fools are lucky), i.e.

if Ivan can win, he wins having struck the least number of blows;
if it is impossible to defeat Gorynych, but is possible to resist him for an infinitely long period of time, then that’s the strategy Ivan chooses;
if Gorynych wins in any case, Ivan aims to resist him for as long as possible.

### ideas
1. 这个题目好混乱～
2. 先考虑终止状态，就是假如h + t >= R => lose (但也需要知道花了最多的时间)
3. h = 0 && t = 0 => win, 也需要知道花了最少的时间 (这里必须存在一个状态，砍掉头以后，头不能在涨回来，砍掉尾巴后，尾巴也不能长回来。否则肯定是win不了的)
4. 要draw的话，必须进入一个平衡状态，就是经过一圈的攻击后，又回到了一个之前存在的状态，且在这个过程中，没有失败
5. dp[x][y]表示从状态(x, y)到达终止状态的最短路径
6. dp[0][0] = 0,
7. var dfs(x, y) int = 从状态(0, 0) 到(x, y)的最短路径
8.  x + y >= R => inf
9.  x = 0 && y = 0 => 0
10. 然后使用砍掉head/tail 的方式（进入下一个状态）
11. 好费劲的题目