Bertown has n junctions and m bidirectional roads. We know that one can get from any junction to any other one by the existing roads.

As there were more and more cars in the city, traffic jams started to pose real problems. To deal with them the government decided to make the traffic one-directional on all the roads, thus easing down the traffic. Your task is to determine whether there is a way to make the traffic one-directional so that there still is the possibility to get from any junction to any other one. If the answer is positive, you should also find one of the possible ways to orient the roads.

### ideas
1. 要找一个环（n条边）
2. 如果存在一个 deg = 1 的点 => false
3. 如果存在critical bridge => false
4. 否则肯定是有答案的
5. 那种8字形的要怎么处理？
6. 貌似是汉密尔顿回路的简化版，汉密尔顿回路要求所有的城市只能被访问一次，且回到原点
7. 这里没有这个要求，只是要求所有的节点都在某个回路上
8. 如果所有点的度数 >= 2, 似乎也不一定，比如例子2
9. 如果所有点的度数都是偶数, 且连通，那么肯定有答案（这是因为，进入一个点后，肯定能离开这个点）
10. 但是如果存在奇数的点（却不一定没有答案）
11. 是不是要用上面的和下面的一起？