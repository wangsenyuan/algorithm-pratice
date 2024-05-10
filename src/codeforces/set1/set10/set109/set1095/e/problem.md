You are given a bracket sequence 𝑠
 consisting of 𝑛
 opening '(' and closing ')' brackets.

A regular bracket sequence is a bracket sequence that can be transformed into a correct arithmetic expression by inserting characters '1' and '+' between the original characters of the sequence. For example, bracket sequences "()()", "(())" are regular (the resulting expressions are: "(1)+(1)", "((1+1)+1)"), and ")(" and "(" are not.

You can change the type of some bracket 𝑠𝑖
. It means that if 𝑠𝑖=
 ')' then you can change it to '(' and vice versa.

Your task is to calculate the number of positions 𝑖
 such that if you change the type of the 𝑖
-th bracket, then the resulting bracket sequence becomes regular.


### ideas
1. 先理解一下，就是改变某个i，从( 变成 ), 或者反过来，结果仍然是regular的
2. 如果长度是奇数 => 0 
3. ()() => 0
4. 如果本来就是regular的，那么也是0，因为只能改变一个位置，改变后，必然就不平衡了
5. 那么剩下的就是，检查balance，如果最后等于2,那么就改 ( 且blance = 2的位置
6. 如果balance = -2, 就改 )