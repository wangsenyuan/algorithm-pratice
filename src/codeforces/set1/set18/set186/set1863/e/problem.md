You are playing a video game. The game has 𝑛
quests that need to be completed. However, the 𝑗
-th quest can only be completed at the beginning of hour ℎ𝑗
of a game day. The game day is 𝑘
hours long. The hours of each game day are numbered 0,1,…,𝑘−1
. After the first day ends, a new one starts, and so on.

Also, there are dependencies between the quests, that is, for some pairs (𝑎𝑖,𝑏𝑖)
the 𝑏𝑖 -th quest can only be completed after the 𝑎𝑖 -th quest. It is guaranteed that there are no circular dependencies,
as otherwise the game would be unbeatable and nobody would play it.

You are skilled enough to complete any number of quests in a negligible amount of time (i. e. you can complete any
number of quests at the beginning of the same hour, even if there are dependencies between them). You want to complete
all quests as fast as possible. To do this, you can complete the quests in any valid order. The completion time is equal
to the difference between the time of completing the last quest and the time of completing the first quest in this
order.

Find the least amount of time you need to complete the game.

### thoughts

1. 对于quest x, y，x -> y, 可以构造dag，然后从topo顺序依次处理
2. 每次处理最近的，且in deg = 0的quest
3. 计算的是最短的持续时间，所以并不是启动越早越好
4. dag分成了很多component，计算出每个component的[l, r]
5. 然后每个comp最多只能移动向后移动k或者不移动，
6. 然后计算这些区间似乎要用到dp，但怎么写呢
7. dp[i][r] = l 表示当右端点时r时，最大的l，可以覆盖前面的所有的区间
8. dp[i+1][r] = min(l[i+1], dp[i][r]) 如果r[i+1] <= r
9. 假设按照r进行升序排，那些完全被某个区间的部分，可以删除掉，只保留那些分离的，或者有部分重叠的部分
10. 似乎清晰起来了，得到上面的数组后，肯定是前面的部分往后整体移动，后面的部分整体，假设在i处，分别计算前面移动后的最大/最小值，以及后面不移动
11. 最大最小值
12. 搞不出来～～～

### solution

First of all, assume that the first quest we complete is at the hour 𝑥
. We can assume that 0≤𝑥<𝑘
, as increasing it by 𝑘
just shifts all the times by 𝑘
. In this case one can greedily find the completion times for all the quests: essentially, for every quest 𝑣
, if we know that the quests it depends on are completed at hours 𝑐1
, ..., 𝑐𝑑
, then the quest 𝑣
cannot be completed earlier than max(𝑐1,…,𝑐𝑑)
. So if we define 𝑓(𝑥,𝑦)
to be the smallest 𝑧≥𝑥
such that 𝑧≡𝑦(mod𝑘)
, then the completion time of 𝑣
is 𝑓(max(𝑐1,…,𝑐𝑑),ℎ𝑣)
. If the quest 𝑣
does not depend on anything, then, obviously, the time if simply 𝑓(𝑥,ℎ𝑣)
.

The problem is that we don't know which starting time 𝑥
is optimal. On the other hand, we know that it can be assumed to be from [0,𝑘)
. Also, there is no sense in 𝑥
not being ℎ𝑣
for any quest 𝑣
without incoming dependencies. So we can do the following: first assume that 𝑥=0
and find all the completion times for the quests, denote them by 𝑐𝑖
. Then consequently increase 𝑥
until it becomes 𝑘
.

Sometimes 𝑥
becomes equal to some ℎ𝑣
where 𝑣
is a quest that could theoretically be completed first. At these moments we know that the answer is no more than
max𝑐𝑖−𝑥
. When we increase 𝑥
again, such quests can no longer have 𝑐𝑖=𝑥
, and thus some values 𝑐𝑖
increase by some multiple of 𝑘
.

However, for 𝑥=𝑘
all 𝑐𝑖
are exactly 𝑘
more than when 𝑥=0
. Therefore, each value of 𝑐𝑖
in this process increases exactly once and exactly by 𝑘
. Now there are two ways to finish the solution.

Sort all events of type "some 𝑐𝑖
, where quest 𝑖
doesn't have incoming dependencies, increase by 𝑘
". For each such event, we run DFS from such vertices to see if some the quests depending on them must also be completed
𝑘
hours later. One can see that this DFS will be run from each vertex exactly once throughout the process, thus resulting
in 𝑂(𝑛+𝑚)
time complexity.
When building all 𝑐𝑖
the first time, also find for each vertex 𝑣
the first 𝑥
when its 𝑐𝑣
must change. If we denote it with 𝑤𝑣
then it can be expressed as max{𝑤𝑢:𝑢→𝑣,𝑐𝑢+𝑘>𝑐𝑣}
. Then sort all events of type "some 𝑐𝑖
, where quest 𝑖
may or may not have incoming dependencies, increase by 𝑘
".
In both approaches one can easily maintain max𝑐𝑖
as this value can only increase.