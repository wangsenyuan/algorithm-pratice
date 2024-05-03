In Python, code blocks don't have explicit begin/end or curly braces to mark beginning and end of the block. Instead,
code blocks are defined by indentation.

We will consider an extremely simplified subset of Python with only two types of statements.

Simple statements are written in a single line, one per line. An example of a simple statement is assignment.

For statements are compound statements: they contain one or several other statements. For statement consists of a header
written in a separate line which starts with "for" prefix, and loop body. Loop body is a block of statements indented
one level further than the header of the loop. Loop body can contain both types of statements. Loop body can't be empty.

You are given a sequence of statements without indentation. Find the number of ways in which the statements can be
indented to form a valid Python program.

### ideas

1. s for statement and f for loop
2. dp[i][0] 表示到目前为止indent = x的计数
3. 如果遇到s, indent是不变的，但是如果遇到一个f，那么有两种情况，如果前面的是f， 那么只能indent
4. 但是如果前面的是s，那么可以不变
5. 遇到s，其实也有可能减小intent
6. dp[x][0/1]表示当前intent是多少，0表示当前的loop还没有body，1表示已经有body了
7. dp[x][0] + s => dp[x][1]
8. dp[x][0] + f => dp[x+1][0]
9. dp[x][1] + s => dp[x-1][1], dp[x][1]
10. dp[x][1] + f => dp[x+1][0] 或者 dp[x-1][0]
11. 上面的不大对
12. 首先注意到一点是f, 才会产生intend
13. 考虑一个状态机，有2个状态0, 1分别表示，正常状态， loop开始状态
14. 0 -> s -> 0 (0遇到一个s只能还是s)
15. 0 -> f -> 1 (0遇到一个f到达状态1)
16. 1 -> s -> 0/1 （两种可能性, 回到0，或者继续保持1）
17. 1 -> f -> 1 (只有一种可能性，继续流程1)
18. 感觉不大对，
19. 遇到一个f，这个f肯定需要被一个s个终结掉
20. 所以fffs 是一体的
21. 然后计算目前最大的indent是多少，那么新的，就是可以在这个范围内选择，或者增大intent

### solution

This problem can be solved using dynamic programming.

Let's consider all possible programs which end with a certain statement at a certain indent. Dynamic programming state
will be an array dp[i][j] which stores the number of such programs ending with statement i at indent j.

The starting state is a one-dimensional array for i = 0: there is exactly one program which consists of the first
statement only, and its last statement has indent 0.

The recurrent formula can be figured out from the description of the statements. When we add command i + 1, its possible
indents depend on the possible indents of command i and on the type of command i. If command i is a for loop, command
i + 1 must have indent one greater than the indent of command i, so dp[i + 1][0]= 0 and dp[i + 1][j]= dp[i][j - 1] for
j>0. If command i is a simple statement, command i + 1 can belong to the body of any loop before it, and have any indent
from 0 to the indent of command i. If we denote the indent of command i (simple statement) as k, the indent of command
i + 1 as j, we need to sum over all cases where k ≥ j: .

The answer to the task is the total number of programs which end with the last command at any indent: .

The complexity of this solution is O(N2).