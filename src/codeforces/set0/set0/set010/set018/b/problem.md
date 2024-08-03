In one one-dimensional world there are n platforms. Platform with index k (platforms are numbered from 1) is a segment with coordinates [(k - 1)m, (k - 1)m + l], and l < m. Grasshopper Bob starts to jump along the platforms from point 0, with each jump he moves exactly d units right. Find out the coordinate of the point, where Bob will fall down. The grasshopper falls down, if he finds himself not on the platform, but if he finds himself on the edge of the platform, he doesn't fall down.

Input
The first input line contains 4 integer numbers n, d, m, l (1 ≤ n, d, m, l ≤ 106, l < m) — respectively: amount of platforms, length of the grasshopper Bob's jump, and numbers m and l needed to find coordinates of the k-th platform: [(k - 1)m, (k - 1)m + l].

Output
Output the coordinates of the point, where the grosshopper will fall down. Don't forget that if Bob finds himself on the platform edge, he doesn't fall down.

### ideas
1. n * d > k * m + l 
2. n * d < (k + 1) * m
3. 也就是在某个k * m + l 到 (k + 1) * m 中间存在d的一个乘数
4. 可以假设这个数x % m > l
5. 那么就有 (k * m + x) % d = 0
6. 这里k还是未知的x % d = 这个可以算出来， 
7. k * m % d = (d - x % d) = y 这个也是已知的
8. k * m = n * d + y
9. 求出最小的k 满足 k * m % d = y
10. m % d 是知道的 但是如何求出k呢？
11. m % d = z
12. k * m % d = ((k % d) * (m % d)) % d = y
13. (k % d) * z % d = y