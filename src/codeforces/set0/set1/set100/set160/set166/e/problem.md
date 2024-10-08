You are given a tetrahedron. Let's mark its vertices with letters A, B, C and D correspondingly.


An ant is standing in the vertex D of the tetrahedron. The ant is quite active and he wouldn't stay idle. At each moment of time he makes a step from one vertex to another one along some edge of the tetrahedron. The ant just can't stand on one place.

You do not have to do much to solve the problem: your task is to count the number of ways in which the ant can go from the initial vertex D to itself in exactly n steps. In other words, you are asked to find out the number of different cyclic paths with the length of n from vertex D to itself. As the number can be quite large, you should print it modulo 1000000007 (109 + 7).



### ideas
1. A, B, C, D的特点，是任何两点都相通
2. f(d) = f(a) + f(b) + f(c)
3. f(d) = 1
4. 构造一个矩阵，除了对角线是0外，其他位置都是1
5. 