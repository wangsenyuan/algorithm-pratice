Vlad decided to go on a trip to the mountains. He plans to move between 𝑛
mountains, some of which are connected by roads. The 𝑖
-th mountain has a height of ℎ𝑖
.

If there is a road between mountains 𝑖
and 𝑗
, Vlad can move from mountain 𝑖
to mountain 𝑗
by spending ℎ𝑗−ℎ𝑖
units of energy. If his energy drops below zero during the transition, he will not be able to move from mountain 𝑖
to mountain 𝑗
. Note that ℎ𝑗−ℎ𝑖
can be negative and then the energy will be restored.

Vlad wants to consider different route options, so he asks you to answer the following queries: is it possible to
construct some route starting at mountain 𝑎
and ending at mountain 𝑏
, given that he initially has 𝑒
units of energy?

### thoughts

1. 考虑一条线路，从 h1, h2, h3 .... hn
2. 如果要能从h1 到达hn, 那么初始的energy至少是max(h1, ..., hn)
3. 换句话说，如果知道(u, v)中间路径的最大值, e >= h(u,v)即可
4. 我知道了, 通过加入山（已经路径），判断(a, b)是否相连即可