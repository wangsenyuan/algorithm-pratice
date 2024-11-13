Alex decided to go on a touristic trip over the country.

For simplicity let's assume that the country has 𝑛
 cities and 𝑚
 bidirectional roads connecting them. Alex lives in city 𝑠
 and initially located in it. To compare different cities Alex assigned each city a score 𝑤𝑖
 which is as high as interesting city seems to Alex.

Alex believes that his trip will be interesting only if he will not use any road twice in a row. That is if Alex came to city 𝑣
 from city 𝑢
, he may choose as the next city in the trip any city connected with 𝑣
 by the road, except for the city 𝑢
.

Your task is to help Alex plan his city in a way that maximizes total score over all cities he visited. Note that for each city its score is counted at most once, even if Alex been there several times during his trip.


### ideas
1. 如果alex可以按照要求访问到所有的city，那么他的score就是所有city的sum
2. 所以，必然有部分city（在同一条路径不能访问两次的条件下）无法被访问到
3. 那么貌似就是那些bridge，就没法被访问到
4. 假设到了节点u，如果alex遇到了两个bridge，那么他只能选择其中的一个（这个是一个子问题）
5. 那些scc是可以被完全访问到，并回到进入的节点（或者任何节点）
6. 所以，先将图变成（scc组成的）树
7. 那么就是从s到某个叶子节点的最大值
8. 还不一样。
9. 如果进入了叶子节点（这个叶子节点只有一个节点组成，那确实是出不来的）
10. 否则其实它是可以出来的
11. ssc + dp