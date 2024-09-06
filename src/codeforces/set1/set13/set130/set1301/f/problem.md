Jaber is a superhero in a large country that can be described as a grid with 𝑛
 rows and 𝑚
 columns, where every cell in that grid contains a different city.

Jaber gave every city in that country a specific color between 1
 and 𝑘
. In one second he can go from the current city to any of the cities adjacent by the side or to any city with the same color as the current city color.

Jaber has to do 𝑞
 missions. In every mission he will be in the city at row 𝑟1
 and column 𝑐1
, and he should help someone in the city at row 𝑟2
 and column 𝑐2
.


### ideas
1. 对于每次query，如果两个格子的颜色一致，答案为1
2. 如果不一致，一个答案是直线距离
3. 还有一个就是通过穿越+移动；假设我们知道从颜色i到颜色j的穿越的最短距离
4. 那么在除了第一个和最后一个颜色中间，每次穿越，肯定要额外增加一次移动
5. 但是第一次和最后一次，有点麻烦。有可能需要移动（比如r1, c1在一个颜色的中间）有可能不需要
6. 所以get了。只需要再检查一下四周的颜色就可以了