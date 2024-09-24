Programmers working on a large project have just received a task to write exactly m lines of code. There are n programmers working on a project, the i-th of them makes exactly ai bugs in every line of code that he writes.

Let's call a sequence of non-negative integers v1, v2, ..., vn a plan, if v1 + v2 + ... + vn = m. The programmers follow the plan like that: in the beginning the first programmer writes the first v1 lines of the given task, then the second programmer writes v2 more lines of the given task, and so on. In the end, the last programmer writes the remaining lines of the code. Let's call a plan good, if all the written lines of the task contain at most b bugs in total.

Your task is to determine how many distinct good plans are there. As the number of plans can be large, print the remainder of this number modulo given positive integer mod.

### ideas
1. 好像顺序没啥关系，每个程序员都需要写代码，假设写了v行，那么就产生了v * a[i]个bug
2. v1 + v2 + ... + vn = m
3. v1 * a1 + v2 * a2 + ... vn * an <= b
4. dp[i][j][x] = dp[i-1][j-v][x - v * a[i]]
5. 这个有4层循环， i, j, x, v所以不大行 
6. v的上限 = min(x / a[i], j)
7. 如果从矩形的角度看，dp[i][j][x] = dp[i-1]{j - v, j}{x - v * a[i], x}的和？
8. 哪些a[i] = 0 的可以先提取出来，它们不产生bug，假设其中m0行，由这些不产生bug的人员编写，C(m, m0) * pow(m0, cnt0)
9. 先写个4层循环的，把例子跑出来看看