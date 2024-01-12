Tema is playing a very interesting computer game.

During the next mission, Tema's character found himself on an unfamiliar planet. Unlike Earth, this planet is flat and
can be represented as an 𝑛×𝑚
rectangle.

Tema's character is located at the point with coordinates (0,0)
. In order to successfully complete the mission, he needs to reach the point with coordinates (𝑛,𝑚)
alive.

Let the character of the computer game be located at the coordinate (𝑖,𝑗)
. Every second, starting from the first, Tema can:

either use vertical hyperjump technology, after which his character will end up at coordinate (𝑖+1,𝑗)
at the end of the second;
or use horizontal hyperjump technology, after which his character will end up at coordinate (𝑖,𝑗+1)
at the end of the second;
or Tema can choose not to make a hyperjump, in which case his character will not move during this second;
The aliens that inhabit this planet are very dangerous and hostile. Therefore, they will shoot from their railguns 𝑟
times.

Each shot completely penetrates one coordinate vertically or horizontally. If the character is in the line of its impact
at the time of the shot (at the end of the second), he dies.

Since Tema looked at the game's source code, he knows complete information about each shot — the time, the penetrated
coordinate, and the direction of the shot.

What is the minimum time for the character to reach the desired point? If he is doomed to die and cannot reach the point
with coordinates (𝑛,𝑚)
, output −1
.

### thoughts

1. 考虑一下什么情况下，他无法到达目的地？
2. 假设在时刻t，Tema不能够处在一个安全的位置上，那么就无法继续
3. 在时刻t，哪些位置是安全的呢？
4. dp[i][j][x] = 如果 i + j <= x， 那么在时刻x时，Tema可以到达位置(i, j)
5. 且dp[i-1][j][x-1] 或者 dp[i][j-1][x-1] 也是安全的
6. 但显然无法真的按照时间去模拟，是因为x的范围太大
7. 但是r的范围比较小，只有100，所以如何做个映射呢？
8. 假设Tema当前位置在(x, y), 且躲过了前k个攻击
9. 那么在第k+1次攻击时，安全的区域能够包括哪些呢？
10. 如果时间差是dt, 那么就是(u, v) (u - x) + (v - y) <= dt 的区域，然后减去那些在下次攻击时的危险区域
11. 在时刻t，正常情况下，所有(x + y) <= t的区域都可以到达
12. 然后排除掉在当前时刻被攻击的区域
13. 居然直接把 x换成r就可以了～。。。。

### solution

Let's first solve it in (𝑛𝑚𝑡)
.

This can be done using dynamic programming. 𝑑𝑝[𝑖][𝑗][𝑘]=𝑡𝑟𝑢𝑒
if the character can be at coordinates (𝑖,𝑗)
at time 𝑡
, otherwise 𝑑𝑝[𝑖][𝑗][𝑘]=𝑓𝑎𝑙𝑠𝑒
. Such dynamics can be easily recalculated: 𝑑𝑝[𝑖][𝑗][𝑘]=𝑑𝑝[𝑖−1][𝑗][𝑘−1]|𝑑𝑝[𝑖][𝑗−1][𝑘−1]|𝑑𝑝[𝑖][𝑗][𝑘−1]
. If the cell is shot by one of the railguns at time 𝑡
, then 𝑑𝑝[𝑖][𝑗][𝑘]=𝑓𝑎𝑙𝑠𝑒
.

Now let's notice that if the character can reach the final point (𝑛,𝑚)
, then he will have to stand still no more than 𝑟
times. To prove this, we can prove another statement: if the character can reach the final point along some trajectory,
then for any such trajectory the character can stand still no more than 𝑟
times. And this statement can already be proven by mathematical induction.

Thus, instead of the 𝑑𝑝[𝑛][𝑚][𝑡]
dynamics, we can calculate the 𝑑𝑝[𝑛][𝑚][𝑟]
dynamics, where the third parameter is the number of times the character stood still. The transitions here are made
similarly.