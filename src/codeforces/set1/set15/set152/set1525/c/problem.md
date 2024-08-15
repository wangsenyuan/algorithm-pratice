There are 𝑛
 robots driving along an OX axis. There are also two walls: one is at coordinate 0
 and one is at coordinate 𝑚
.

The 𝑖
-th robot starts at an integer coordinate 𝑥𝑖 (0<𝑥𝑖<𝑚)
 and moves either left (towards the 0
) or right with the speed of 1
 unit per second. No two robots start at the same coordinate.

Whenever a robot reaches a wall, it turns around instantly and continues his ride in the opposite direction with the same speed.

Whenever several robots meet at the same integer coordinate, they collide and explode into dust. Once a robot has exploded, it doesn't collide with any other robot. Note that if several robots meet at a non-integer coordinate, nothing happens.

For each robot find out if it ever explodes and print the time of explosion if it happens and −1
 otherwise.

 ### ideas
 1. 先不考虑回头的问题，考虑第一次，要在整数位置相撞的话，两个robot的距离，必须是偶数
 2. 所以，理论上最后最多只能剩下2个（一个在偶数位置，一个在奇数位置）
 3. 即使有回头的情况，上面的结论仍然是成立的
 4. 然后考虑robot i, 如果它是相左运动，正常它会和它左边的相撞
 5. 除非：1 左边没有robot（和它距离为偶数的）2，左边的在和它相撞前，已经撞完了
 6. 否则，它肯定和最靠近它的左边的那个相撞
 7. 感觉是个stack就可以了