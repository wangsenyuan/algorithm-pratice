A team of furry rescue rangers was sitting idle in their hollow tree when suddenly they received a signal of distress. In a few moments they were ready, and the dirigible of the rescue chipmunks hit the road.

We assume that the action takes place on a Cartesian plane. The headquarters of the rescuers is located at point (x1, y1), and the distress signal came from the point (x2, y2).

Due to Gadget's engineering talent, the rescuers' dirigible can instantly change its current velocity and direction of movement at any moment and as many times as needed. The only limitation is: the speed of the aircraft relative to the air can not exceed  meters per second.

Of course, Gadget is a true rescuer and wants to reach the destination as soon as possible. The matter is complicated by the fact that the wind is blowing in the air and it affects the movement of the dirigible. According to the weather forecast, the wind will be defined by the vector (vx, vy) for the nearest t seconds, and then will change to (wx, wy). These vectors give both the direction and velocity of the wind. Formally, if a dirigible is located at the point (x, y), while its own velocity relative to the air is equal to zero and the wind (ux, uy) is blowing, then after  seconds the new position of the dirigible will be .

Gadget is busy piloting the aircraft, so she asked Chip to calculate how long will it take them to reach the destination if they fly optimally. He coped with the task easily, but Dale is convinced that Chip has given the random value, aiming only not to lose the face in front of Gadget. Dale has asked you to find the right answer.

It is guaranteed that the speed of the wind at any moment of time is strictly less than the maximum possible speed of the airship relative to the air.


### ideas
1. 在t秒的时候，应该利用(vx, vy)尽可能的离(x2, y2)近？
2. 似乎也不一定，这个还需要看(wx, wy)的方向
3. 假设在t秒的时候，到到了(x, y) , 然后看花了多久从(x, y)到达(x2, y2)
4. 假设共花费了t秒，t1秒到到了(x, y), 2秒从(x, y)到到了(x2, y2)
5. t1 + t2 = t
6. ux = (x - x1) / t1, uy = (y - y1) / t1
7. ux -= vx, uy -= vy
8. ux * ux + uy * uy <= vmax * vmax
9. (ux, uy)是实际的速度，然后减去风的速度，就是飞行的速度
10. 所以，(x, y)组成了一个区域（椭圆？）
11. 用(wx, wy)也可以找到一个以（x2, y2)的区域，这两个区域有交集时，就时有解的
12. 但是问题是t1是一个很大的范围，且感觉它不具备二分的性质。
13. 这里t1是给定的
14. 所以，只需要二分t2就可以了
15. dx := (x - x2) / t2, dy := (y - y2) / t2
16. dx -= wx, dy -= wy
17. 但是得到的这个范围，要怎么判断，能够(x1, y1) 到达？
18. 找到内部的一个点（x1, y1), 看看，在它和（上一个区域的最近的距离，是否在t秒内到达）？
19. 不对，因为有速度的问题，所以应该是找到两条切线（看这两条切线所需的时间，只要其中一个切线的时间 <= t)就可以了
20. 但这些公式，不清楚呐
21. 