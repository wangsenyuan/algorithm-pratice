Recently, Olya received a magical square with the size of 2𝑛×2𝑛
.

It seems to her sister that one square is boring. Therefore, she asked Olya to perform exactly 𝑘
splitting operations.

A Splitting operation is an operation during which Olya takes a square with side 𝑎
and cuts it into 4 equal squares with side 𝑎2
. If the side of the square is equal to 1
, then it is impossible to apply a splitting operation to it (see examples for better understanding).

Olya is happy to fulfill her sister's request, but she also wants the condition of Olya's happiness to be satisfied
after all operations.

The condition of Olya's happiness will be satisfied if the following statement is fulfilled:

Let the length of the side of the lower left square be equal to 𝑎
, then the length of the side of the right upper square should also be equal to 𝑎
. There should also be a path between them that consists only of squares with the side of length 𝑎
. All consecutive squares on a path should have a common side.

Obviously, as long as we have one square, these conditions are met. So Olya is ready to fulfill her sister's request
only under the condition that she is satisfied too. Tell her: is it possible to perform exactly 𝑘
splitting operations in a certain order so that the condition of Olya's happiness is satisfied? If it is possible, tell
also the size of the side of squares of which the path from the lower left square to the upper right one will consist.

### ideas

1. 每次都是对一块区域进行4等分
2. 假设一整块，一直进行4等分，长度位n (1 << n), 那么一共可以进行多少次呢？
3. 如果 f(n) 表示次数 f(0) = 0 (长度为1)
4. f(1) = 1 (只能进行一次)
5. f(2) = 1 + 4 * f(1) = 5
6. f(3) = 1 + 4 * f(2) ..
7. f(i) = 1 + 4 * f(i-1)
8. 所以，如果 k > f(n) 显然是不行的，k <= f(n)
9. 然后就是考虑条件2，
10. 这个条件表明三块区域的边长，必须是a的倍数(n/2) 要能整除它
11. 然后剩下的那个区域，可以正常处理(1为单位)