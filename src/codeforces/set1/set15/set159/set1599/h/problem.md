This is an interactive problem!

As part of your contribution in the Great Bubble War, you have been tasked with finding the newly built enemy fortress. The world you live in is a giant 109×109
 grid, with squares having both coordinates between 1
 and 109
.

You know that the enemy base has the shape of a rectangle, with the sides parallel to the sides of the grid. The people of your world are extremely scared of being at the edge of the world, so you know that the base doesn't contain any of the squares on the edges of the grid (the 𝑥
 or 𝑦
 coordinate being 1
 or 109
).

To help you locate the base, you have been given a device that you can place in any square of the grid, and it will tell you the manhattan distance to the closest square of the base. The manhattan distance from square (𝑎,𝑏)
 to square (𝑝,𝑞)
 is calculated as |𝑎−𝑝|+|𝑏−𝑞|
. If you try to place the device inside the enemy base, you will be captured by the enemy. Because of this, you need to make sure to never place the device inside the enemy base.

Unfortunately, the device is powered by a battery and you can't recharge it. This means that you can use the device at most 40
 times.

Input
The input contains the answers to your queries.

Interaction
Your code is allowed to place the device on any square in the grid by writing "? 𝑖
 𝑗
" (1≤𝑖,𝑗≤109)
. In return, it will recieve the manhattan distance to the closest square of the enemy base from square (𝑖,𝑗)
 or −1
 if the square you placed the device on is inside the enemy base or outside the grid.

If you recieve −1
 instead of a positive number, exit immidiately and you will see the wrong answer verdict. Otherwise, you can get an arbitrary verdict because your solution will continue to read from a closed stream.

Your solution should use no more than 40
 queries.

Once you are sure where the enemy base is located, you should print "! 𝑥
 𝑦
 𝑝
 𝑞
" (1≤𝑥≤𝑝≤109,1≤𝑦≤𝑞≤109)
, where (𝑥,𝑦)
 is the square inside the enemy base with the smallest 𝑥
 and 𝑦
 coordinates, and (𝑝,𝑞)
 is the square inside the enemy base with the largest 𝑥
 and 𝑦
 coordinates. Note that answering doesn't count as one of the 40 queries.


### ideas
1. (1, 1)这里问一下，返回的是(x1 - 1, y1 - 1) = a
2. 如果 a = 2 => x1 = 2, y1 = 2 else
3. (1, 2) (x1 - 1, y1 - 2) = b => 两个联立，貌似算不出x1, y1 （b = a - 1)
4. (1, 1) => x1 - 1 + y1 - 1 = a
5. (1, n) => x1 - 1 + n - y2 = b
6. (n, 1) => n - x2 + y1 - 1 = c
7. (n, n) => n - x2 + n - y2 = d
8. n - y2 - y1 + 1 = b - a
9. y2 + y1 = n + 1 - (b - a)
10. 所以本质上，还是4个参数，2个方程
11. y2 + y1 => 是左右两端的和
12. 那么它的中点就是 (y2 + y1) / 2 (两个或者1个)
13. 假设y1 = 2, y2 = 6 => mid = 4
14. 如果使用（1， 4） => x1 - 1 + 4 - y1 = u =>可以推论出来的
15. 