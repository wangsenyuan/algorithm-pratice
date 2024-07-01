You are playing a famous computer game (that just works) where you have various skills you can level up. Today, you focused on the "Smithing" skill. Your tactic is obvious: forging weapons from ingots and then melting them back to return the materials partially. For simplicity, every time you create an item, you get 1
 experience point, and every time you melt an item, you also get 1
 experience point.

There are 𝑛
 classes of weapons you can forge and 𝑚
 types of metal ingots.

You can create one weapon of the 𝑖
-th class, spending 𝑎𝑖
 ingots of metal of the same type. Melting a weapon of the 𝑖
-th class (which you crafted earlier) returns you 𝑏𝑖
 ingots of the type of metal it was made of.

You have 𝑐𝑗
 metal ingots of the 𝑗
-th type, and you know that you can craft a weapon of any class from any metal type. Each combination of a weapon class and a metal type can be used any number of times.

What is the maximum total amount of experience you can earn by crafting and melting weapons?

### ideas
1. 这个题目的描述有点废～～
2. c[j]表示有c[j]个类型位j的金属锭
3. a[i], b[i]表示使用a[i]个任意（同一种）类型的金属锭，然后将其溶化后获得b[i]个该类型的金属锭，并获得2个经验值
4. 如果制作了一个武器，肯定是要溶化，并获得部分的金属锭更优，所以这里可以把制作->溶化当作一个操作
5. 假设有x个金属锭，那么应该选择制作哪种武器呢？很明显应该选择那些返还更多的，这样子才能继续下去。当然前提是c[i] >= a[i]
6. 感觉花费同样的(a[i] - b[i])可以同时处理，且越少越好。在同样的花费下，应该使用越大的a[i]， 这是因为后面可能没有机会使用它

### solution

It's quite obvious that we should melt every weapon we forge, and we can do it as soon as we forge it. So, let's do these actions in pairs: forge a weapon, then instantly melt it.

Since you can't use two types of metal while forging one weapon, we can solve our task independently for each metal type. Suppose you have 𝑥
 ingots of the chosen metal.

You can act greedily: you can forge only weapons of classes with 𝑎𝑖≤𝑥
, but among them, it's optimal to chose one that makes us lose the minimum number of ingots by forging it and melting it; so, we need to minimize the value of 𝑎𝑖−𝑏𝑖
. However, this greedy solution is too slow when implemented naively.

Let's define 𝐴=max𝑎𝑖
 and look at two cases: 𝑥≤𝐴
 or 𝑥>𝐴
. If 𝑥≤𝐴
, let's just precalculate answers ans[𝑥]
 for all 𝑥
 from 0
 to 𝐴
. To do so, let's precalculate another value best[𝑥]
 as "the minimum 𝑎𝑖−𝑏𝑖
 among all classes 𝑖
 where 𝑎𝑖≤𝑥
". In other words, best[𝑥]
 will be equal to the minimum amount of ingots you'll lose if you have 𝑥
 ingots, and you need to forge-melt one weapon.

We can precalculate best[𝑥]
 in two steps:

for each class 𝑖
, we can set best[𝑎𝑖]=min(best[𝑎𝑖],𝑎𝑖−𝑏𝑖)
;
using approach similar to prefix minima or dynamic programming, we can also update best[𝑥]=min(best[𝑥],best[𝑥−1])
 for all 𝑥
 from 1
 to 𝐴
.
Using the array best
, we can calculate ans
 easily with dynamic programming: ans[𝑥]=2+ans[𝑥−best[𝑥]]
 for all 𝑥
 from 1
 to 𝐴
.
In case 𝑥>𝐴
, we can forge-melt a weapon with minimum 𝑎𝑖−𝑏𝑖
 as many times as we can until 𝑥≤𝐴
. Once 𝑥
 becomes less or equal to 𝐴
, we can take the value ans[𝑥]
 and finish processing that metal type.

Knowing that the minimum 𝑎𝑖−𝑏𝑖
 is just best[𝐴]
, we can reformulate our first step as finding minimum 𝑘
 such that 𝑥−𝑘⋅best[𝐴]≤𝐴
 or 𝑘≥𝑥−𝐴best[𝐴]
. In other words, we'll make the first step exactly 𝑘=⌈𝑥−𝐴best[𝐴]⌉
 times. So we can calculate the final answer as 2𝑘+ans[𝑥−𝑘⋅best[𝐴]]
.

As a result, we run precalculating in 𝑂(𝑛+𝐴)
 time and process each metal type in 𝑂(1)
 time. The total complexity is 𝑂(𝑛+𝑚+𝐴)
 time and space.