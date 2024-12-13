There are 𝑛
 children in a class, 𝑚
 pairs among them are friends. The 𝑖
-th pair who are friends have a friendship value of 𝑓𝑖
.

The teacher has to go for 𝑘
 excursions, and for each of the excursions she chooses a pair of children randomly, equiprobably and independently. If a pair of children who are friends is chosen, their friendship value increases by 1
 for all subsequent excursions (the teacher can choose a pair of children more than once). The friendship value of a pair who are not friends is considered 0
, and it does not change for subsequent excursions.

Find the expected value of the sum of friendship values of all 𝑘
 pairs chosen for the excursions (at the time of being chosen). It can be shown that this answer can always be expressed as a fraction 𝑝𝑞
 where 𝑝
 and 𝑞
 are coprime integers. Calculate 𝑝⋅𝑞−1mod(109+7)
.

### ideas
1. 一共可以选的pair数量是 n * (n - 1) / 2 = X
2. 所以，一对friends被选中（一次）的概率为 1/X
3. 一共m对，每次选中其中一对的概率 m / X * k? (选中任何一对，都是增加1) + sum(f)
4. 