Consider some square matrix A with side n consisting of zeros and ones. There are n rows numbered from 1 to n from top to bottom and n columns numbered from 1 to n from left to right in this matrix. We'll denote the element of the matrix which is located at the intersection of the i-row and the j-th column as Ai, j.

Let's call matrix A clear if no two cells containing ones have a common side.

Let's call matrix A symmetrical if it matches the matrices formed from it by a horizontal and/or a vertical reflection. Formally, for each pair (i, j) (1 ≤ i, j ≤ n) both of the following conditions must be met: Ai, j = An - i + 1, j and Ai, j = Ai, n - j + 1.

Let's define the sharpness of matrix A as the number of ones in it.

Given integer x, your task is to find the smallest positive integer n such that there exists a clear symmetrical matrix A with side n and sharpness x.


### ideas
如果n是奇数, (n * n + 1) / 2 个1可以放进去
如果n是偶数 ((n - 1) * (n - 1) + 1) + n - 1 个
如果 x % 4 == 1, 那么可以用 n * n + 1 / 2 这个模式去匹配
如果 x % 4 = 2, 那么必须有一个地方是上下，或者左右, 且必须是中轴线
如果 x % 4 = 3, 那么必须在中心，中轴线上，都要有
如果 x % 4 = 0, 和1的情况类似，但是不用中心
