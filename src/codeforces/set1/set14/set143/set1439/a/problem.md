This is the easy version of the problem. The difference between the versions is in the number of possible operations that can be made. You can make hacks if and only if you solved both versions of the problem.

You are given a binary table of size 𝑛×𝑚
. This table consists of symbols 0
 and 1
.

You can make such operation: select 3
 different cells that belong to one 2×2
 square and change the symbols in these cells (change 0
 to 1
 and 1
 to 0
).

Your task is to make all symbols in the table equal to 0
. You are allowed to make at most 3𝑛𝑚
 operations. You don't need to minimize the number of operations.

It can be proved that it is always possible.

### ideas
1. 1111 =》0001
2. 始终保证上面一行都是0，在可能的情况下，保证前面一列是0