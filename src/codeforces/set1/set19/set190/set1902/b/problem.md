Monocarp is a student at Berland State University. Due to recent changes in the Berland education system, Monocarp has
to study only one subject — programming.

The academic term consists of 𝑛
days, and in order not to get expelled, Monocarp has to earn at least 𝑃
points during those 𝑛
days. There are two ways to earn points — completing practical tasks and attending lessons. For each practical task
Monocarp fulfills, he earns 𝑡
points, and for each lesson he attends, he earns 𝑙
points.

Practical tasks are unlocked "each week" as the term goes on: the first task is unlocked on day 1
(and can be completed on any day from 1
to 𝑛
), the second task is unlocked on day 8
(and can be completed on any day from 8
to 𝑛
), the third task is unlocked on day 15
, and so on.

Every day from 1
to 𝑛
, there is a lesson which can be attended by Monocarp. And every day, Monocarp chooses whether to study or to rest the
whole day. When Monocarp decides to study, he attends a lesson and can complete no more than 2
tasks, which are already unlocked and not completed yet. If Monocarp rests the whole day, he skips a lesson and ignores
tasks.

Monocarp wants to have as many days off as possible, i. e. he wants to maximize the number of days he rests. Help him
calculate the maximum number of days he can rest!

### thoughts

1. 假设mono一共使用了m天，那么他take lessions的应该放在最后的几天
2. 这是因为在take lesson的时候，可以完成2个task，而task是每隔7天出现一个，且不会消失
3. 越在后面，就有足够的task来完成，并获得分数
4. 所以，mono始终在最后的x天进行处理，即可
5. 那么要求 x * l + min(y, 2 * x) * t >= P (y = 激活的task的数量)
6. y = (n + 6) / 7 (只要在最后，肯定能用到)
7. x * l + y * t >= P (y <= 2 * x)
8. 