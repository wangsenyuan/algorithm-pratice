You are given an unweighted tree with n vertices. Then n - 1 following operations are applied to the tree. A single
operation consists of the following steps:

choose two leaves;
add the length of the simple path between them to the answer;
remove one of the chosen leaves from the tree.
Initial answer (before applying operations) is 0. Obviously after n - 1 such operations the tree will consist of a
single vertex.

Calculate the maximal possible answer you can achieve, and construct a sequence of operations that allows you to achieve
this answer!

Input
The first line contains one integer number n (2 ≤ n ≤ 2·105) — the number of vertices in the tree.

Next n - 1 lines describe the edges of the tree in form ai, bi (1 ≤ ai, bi ≤ n, ai ≠ bi). It is guaranteed that given
graph is a tree.

Output
In the first line print one integer number — maximal possible answer.

In the next n - 1 lines print the operations in order of their applying in format ai, bi, ci, where ai, bi — pair of the
leaves that are chosen in the current operation (1 ≤ ai, bi ≤ n), ci (1 ≤ ci ≤ n, ci = ai or ci = bi) — choosen leaf
that is removed from the tree in the current operation.

See the examples for better understanding.

### ideas

1. find the diameter
2. remove others before diameter