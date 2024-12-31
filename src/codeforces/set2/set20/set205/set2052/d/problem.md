Consider a simple single-bit boolean register that supports two operations:

set — sets the register to true if it was false, and returns true; otherwise, it returns false;
unset — sets the register to false if it was true, and returns true; otherwise, it returns false.
The initial state of the register is false. Suppose there were 𝑛
 operations 𝑜𝑝𝑖
 (for 1≤𝑖≤𝑛
) where at most two operations returned true. Also, we are given the partial order of operations as a directed acyclic graph (DAG): an edge 𝑖→𝑗
 means that 𝑜𝑝𝑖
 happened before 𝑜𝑝𝑗
. You are asked whether it is possible to put these operations in some linear sequential order that satisfies the given partial order and such that if operations are applied to the register in that order, their results are the same as given.

Input
In the first line, you are given an integer 𝑛
 — the number of operations (1≤𝑛≤105
). In the following 𝑛
 lines, you are given operations in the format "type result", where type is either "set" or "unset" and result is either "true" or "false". It is guaranteed that at most two operations have "true" results.

In the next line, you are given an integer 𝑚
 — the number of arcs of the DAG (0≤𝑚≤105
). In the following 𝑚
 lines, you are given arcs — pairs of integers 𝑎
 and 𝑏
 (1≤𝑎,𝑏≤𝑛
; 𝑎≠𝑏
). Each arc indicates that operation 𝑜𝑝𝑎
 happened before operation 𝑜𝑝𝑏
.

Output
Print any linear order of operations that satisfies the DAG constraints and ensures the results of the operations match the ones given in the input. If a correct operation order does not exist, print −1
.

### ideas
1. 很有意思的一个题目，只有两次成功，肯定是一次 set true, 一次 unset true
2. 所有的 set-false, 必须在 set-true 的后面
3. 且在 set-true 和 set-false 中间不能有 un-set (否则，肯定会出现第二次set-true)
4. unset-true 必须在set-true（所有的set）后面，unset-false 可以分布在 set-true/unset-true 的前后
5. unset-false... set-true, set-false..... unset-true, unset-false....
6. 应该是这样的一个顺序
7. 必须保证set-true不能在set-false 的后面(要么 set-true -> set-false, 要么没有顺序),
8. 且任何set-之间不能有unset 
9. 可以搞