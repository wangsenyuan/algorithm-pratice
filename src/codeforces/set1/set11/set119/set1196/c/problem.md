𝑛
  robots have escaped from your laboratory! You have to find them as soon as possible, because these robots are experimental, and their behavior is not tested yet, so they may be really dangerous!

Fortunately, even though your robots have escaped, you still have some control over them. First of all, you know the location of each robot: the world you live in can be modeled as an infinite coordinate plane, and the 𝑖
-th robot is currently located at the point having coordinates (𝑥𝑖
, 𝑦𝑖
). Furthermore, you may send exactly one command to all of the robots. The command should contain two integer numbers 𝑋
 and 𝑌
, and when each robot receives this command, it starts moving towards the point having coordinates (𝑋
, 𝑌
). The robot stops its movement in two cases:

either it reaches (𝑋
, 𝑌
);
or it cannot get any closer to (𝑋
, 𝑌
).
Normally, all robots should be able to get from any point of the coordinate plane to any other point. Each robot usually can perform four actions to move. Let's denote the current coordinates of the robot as (𝑥𝑐
, 𝑦𝑐
). Then the movement system allows it to move to any of the four adjacent points:

the first action allows it to move from (𝑥𝑐
, 𝑦𝑐
) to (𝑥𝑐−1
, 𝑦𝑐
);
the second action allows it to move from (𝑥𝑐
, 𝑦𝑐
) to (𝑥𝑐
, 𝑦𝑐+1
);
the third action allows it to move from (𝑥𝑐
, 𝑦𝑐
) to (𝑥𝑐+1
, 𝑦𝑐
);
the fourth action allows it to move from (𝑥𝑐
, 𝑦𝑐
) to (𝑥𝑐
, 𝑦𝑐−1
).
Unfortunately, it seems that some movement systems of some robots are malfunctioning. For each robot you know which actions it can perform, and which it cannot perform.

You want to send a command so all robots gather at the same point. To do so, you have to choose a pair of integer numbers 𝑋
 and 𝑌
 so that each robot can reach the point (𝑋
, 𝑌
). Is it possible to find such a point?

### ideas
1. 假设存在这样的一个点(x, y), 这个点，应该是这些robot的一个位置？
2. 考虑有两个机器人，在某个不是它们位置的地方(x, y)能相遇
3. 如果其中给一个机器人只能向下，和向右运动，
4. 另外一个机器人能向上、向右运动
5. 所以，它们肯定可以在y-1的位置遇到
6. 所以，可以设定相遇的位置就是某个机器人的位置
7. 那就可以排序了