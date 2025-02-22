Once Petya read a problem about a bracket sequence. He gave it much thought but didn't find a solution. Today you will face it.

You are given string s. It represents a correct bracket sequence. A correct bracket sequence is the sequence of opening ("(") and closing (")") brackets, such that it is possible to obtain a correct mathematical expression from it, inserting numbers and operators between the brackets. For example, such sequences as "(())()" and "()" are correct bracket sequences and such sequences as ")()" and "(()" are not.

In a correct bracket sequence each bracket corresponds to the matching bracket (an opening bracket corresponds to the matching closing bracket and vice versa). For example, in a bracket sequence shown of the figure below, the third bracket corresponds to the matching sixth one and the fifth bracket corresponds to the fourth one.


You are allowed to color some brackets in the bracket sequence so as all three conditions are fulfilled:

Each bracket is either not colored any color, or is colored red, or is colored blue.
For any pair of matching brackets exactly one of them is colored. In other words, for any bracket the following is true: either it or the matching bracket that corresponds to it is colored.
No two neighboring colored brackets have the same color.
Find the number of different ways to color the bracket sequence. The ways should meet the above-given conditions. Two ways of coloring are considered different if they differ in the color of at least one bracket. As the result can be quite large, print it modulo 1000000007 (109 + 7).

### ideas
1. 任何一对可以提前找出来
2. 然后没对有 2 * 2 = 4 （状态， 黑/红， 红/黑， 黑/蓝， 蓝/黑）
3. 如果是 (...) 的结构， 在确定它们状态的情况下, 里面的是会被影响到的
4. 还有就是前面的也会被影响到，后面的
5. dp[s][x][y] 表示字符串s，它的头/尾的状态是x/y时的计数
6. 如果 s = (...) 的结构， 那么x,y有限制，否则没有特殊的限制，只要满足第三个条件即可
7. 是不是可以搞呢？