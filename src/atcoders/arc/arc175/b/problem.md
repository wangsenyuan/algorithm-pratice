You are given a string 
S of length 
2N consisting of the characters ( and ). Let 
S 
i
​
  denote the 
i-th character from the left of 
S.

You can perform the following two types of operations zero or more times in any order:

Choose a pair of integers 
(i,j) such that 
1≤i<j≤2N. Swap 
S 
i
​
  and 
S 
j
​
 . The cost of this operation is 
A.

Choose an integer 
i such that 
1≤i≤2N. Replace 
S 
i
​
  with ( or ). The cost of this operation is 
B.

Your goal is to make 
S a correct parenthesis sequence. Find the minimum total cost required to achieve this goal. It can be proved that the goal can always be achieved with a finite number of operations.

### ideas
1. 如果 2 * B <= A， 那么操作2显然更优（没有限制，且不产生更差的结果）
2. 如果 2 * B > A, 那么显然，应该尽可能使用操作1，然后使用操作2
3. 现在考虑如何进行操作1。如果现在出现了一个多余的右括号，那么应该把最远处的左扩号给替换过来？
4. )))(()
5. -1, -2, -3, -2, -1, -2
6. 1, 0, -1, 0, -1, -2
7. 1, 0, 1, 2, 1, 0
8. 