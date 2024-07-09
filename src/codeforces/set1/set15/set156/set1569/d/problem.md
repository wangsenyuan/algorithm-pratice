There is a city that can be represented as a square grid with corner points in (0,0)
 and (106,106)
.

The city has 𝑛
 vertical and 𝑚
 horizontal streets that goes across the whole city, i. e. the 𝑖
-th vertical streets goes from (𝑥𝑖,0)
 to (𝑥𝑖,106)
 and the 𝑗
-th horizontal street goes from (0,𝑦𝑗)
 to (106,𝑦𝑗)
.

All streets are bidirectional. Borders of the city are streets as well.

There are 𝑘
 persons staying on the streets: the 𝑝
-th person at point (𝑥𝑝,𝑦𝑝)
 (so either 𝑥𝑝
 equal to some 𝑥𝑖
 or 𝑦𝑝
 equal to some 𝑦𝑗
, or both).

Let's say that a pair of persons form an inconvenient pair if the shortest path from one person to another going only by streets is strictly greater than the Manhattan distance between them.

Calculate the number of inconvenient pairs of persons (pairs (𝑥,𝑦)
 and (𝑦,𝑥)
 are the same pair).

Let's recall that Manhattan distance between points (𝑥1,𝑦1)
 and (𝑥2,𝑦2)
 is |𝑥1−𝑥2|+|𝑦1−𝑦2|
.

### ideas
1. 可以反过来计算有哪些是convenient 的pair吗？
2. 好像可以。假设从高到低，考虑当前person x
3. 如果x在横向的街上面，那么所以那些在垂直方向街上面的人，都和他是方便的。反之也是
4. 但是那些在路口的，就有点麻烦了
5. 因为路口的和所有人都convenient，是不是把他给忽略掉，反而容易些？好像是的