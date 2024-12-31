A string is called bracket sequence if it does not contain any characters other than "(" and ")". A bracket sequence is called regular (shortly, RBS) if it is possible to obtain correct arithmetic expression by inserting characters "+" and "1" into this sequence. For example, "", "(())" and "()()" are RBS and ")(" and "(()" are not.

We can see that each opening bracket in RBS is paired with some closing bracket, and, using this fact, we can define nesting depth of the RBS as maximum number of bracket pairs, such that the 2
-nd pair lies inside the 1
-st one, the 3
-rd one — inside the 2
-nd one and so on. For example, nesting depth of "" is 0
, "()()()" is 1
 and "()((())())" is 3
.

Now, you are given RBS 𝑠
 of even length 𝑛
. You should color each bracket of 𝑠
 into one of two colors: red or blue. Bracket sequence 𝑟
, consisting only of red brackets, should be RBS, and bracket sequence, consisting only of blue brackets 𝑏
, should be RBS. Any of them can be empty. You are not allowed to reorder characters in 𝑠
, 𝑟
 or 𝑏
. No brackets can be left uncolored.

Among all possible variants you should choose one that minimizes maximum of 𝑟
's and 𝑏
's nesting depth. If there are multiple solutions you can print any of them.

### ideas
1. 是不是正好一半呢？
2. ((())) => (()) ()