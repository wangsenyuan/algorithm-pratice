There are 𝑛
 cities located on a straight line. The cities are numbered from 1
 to 𝑛
.

Portals are used to move between cities. There are 4
 colors of portals: blue, green, red, and yellow. Each city has portals of two different colors. You can move from city 𝑖
 to city 𝑗
 if they have portals of the same color (for example, you can move between a "blue-red" city and a "blue-green" city). This movement costs |𝑖−𝑗|
 coins.

Your task is to answer 𝑞
 independent queries: calculate the minimum cost to move from city 𝑥
 to city 𝑦
.

### ideas
1. 一共有4条高速路，每个城市相当于连接到了两条高速路
2. 都需要进入高速路，然后再移动到目的地
3. 考虑给定的一组(x, y)如何实现？
4. 如果(x, y)有相同颜色的portal，那么答案就是它们之间的距离
5. 如果没有，那么它们必须移动到某个具有它们共同颜色portal的城市，
6. 假设x具有(a, b), y具有(c, d) 
7. 那么必须到某个city z, 具有，要们(b, c), 要么（a, d), 要么（a, c), 要么(b, d)
8. 