Anton and Dasha like to play different games during breaks on checkered paper. By the 11th grade they managed to play all the games of this type and asked Vova the programmer to come up with a new game. Vova suggested to them to play a game under the code name "dot" with the following rules:

On the checkered paper a coordinate system is drawn. A dot is initially put in the position (x, y).
A move is shifting a dot to one of the pre-selected vectors. Also each player can once per game symmetrically reflect a dot relatively to the line y = x.
Anton and Dasha take turns. Anton goes first.
The player after whose move the distance from the dot to the coordinates' origin exceeds d, loses.
Help them to determine the winner.

Input
The first line of the input file contains 4 integers x, y, n, d ( - 200 ≤ x, y ≤ 200, 1 ≤ d ≤ 200, 1 ≤ n ≤ 20) — the initial coordinates of the dot, the distance d and the number of vectors. It is guaranteed that the initial dot is at the distance less than d from the origin of the coordinates. The following n lines each contain two non-negative numbers xi and yi (0 ≤ xi, yi ≤ 200) — the coordinates of the i-th vector. It is guaranteed that all the vectors are nonzero and different.

Output
You should print "Anton", if the winner is Anton in case of both players play the game optimally, and "Dasha" otherwise.

### ideas
1. 有点奇怪，那个reflect，似乎不起作用啊
2. 如果reflect前会失败，refelect会获胜，但是另外一方也可以reflect回来哪
3. 那reflect就是干扰项。
4. dp[x][y] = -1, 表示[x, y]是失败点，任何人走到这步，都是会输
5. dp[x][y] = 1 表示这是一个胜利点，任何人走到这一步都会胜利
6. abs(x) + abs(y) > d => -1
7. dp[x][y] = 1， 所有的vector都指向了失败点， else 它也是一个失败点
8. 这样的点一共有 400 * 400 个