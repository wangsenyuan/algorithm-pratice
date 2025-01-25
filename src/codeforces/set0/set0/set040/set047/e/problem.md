Bertown is under siege! The attackers have blocked all the ways out and their cannon is bombarding the city. Fortunately, Berland intelligence managed to intercept the enemies' shooting plan. Let's introduce the Cartesian system of coordinates, the origin of which coincides with the cannon's position, the Ox axis is directed rightwards in the city's direction, the Oy axis is directed upwards (to the sky). The cannon will make n more shots. The cannon balls' initial speeds are the same in all the shots and are equal to V, so that every shot is characterized by only one number alphai which represents the angle at which the cannon fires. Due to the cannon's technical peculiarities this angle does not exceed 45 angles (π / 4). We disregard the cannon sizes and consider the firing made from the point (0, 0).

The balls fly according to the known physical laws of a body thrown towards the horizon at an angle:

vx(t) = V·cos(alpha)
vy(t) = V·sin(alpha)  –  g·t
x(t) = V·cos(alpha)·t
y(t) = V·sin(alpha)·t  –  g·t2 / 2
Think of the acceleration of gravity g as equal to 9.8.

Bertown defends m walls. The i-th wall is represented as a vertical segment (xi, 0) - (xi, yi). When a ball hits a wall, it gets stuck in it and doesn't fly on. If a ball doesn't hit any wall it falls on the ground (y = 0) and stops. If the ball exactly hits the point (xi, yi), it is considered stuck.

Your task is to find for each ball the coordinates of the point where it will be located in the end.

### ideas
1. 先把角度按照从小到大排序
2. 因为它们都没有超过45度，所以它们能够到达的最高点和距离，应该按照角度是在增长的
3. 也就是说，如果ai,能够到达h, 那么 a[i+1]也可以到达
4. 可以分两个阶段，一个是上升期，一个是下降期，随着角度的增加，上升期和下降期的分隔点在往后移动
5. 可以为每个wall算一个角度范围，就是在什么情况下，会落在它上面
6. (xi, yi) 如果是在上升阶段 
7. xi = V * cos(a) * t
8. yi = V * sin(a) * t - g * t * t / 2
9. yi = V * sin(a) * x / (V * cos(a)) - g * (x`)
10. yi = tan(a) * xi - g ....