Polycarp has a string 𝑠
. Polycarp performs the following actions until the string 𝑠
 is empty (𝑡
 is initially an empty string):

he adds to the right to the string 𝑡
 the string 𝑠
, i.e. he does 𝑡=𝑡+𝑠
, where 𝑡+𝑠
 is a concatenation of the strings 𝑡
 and 𝑠
;
he selects an arbitrary letter of 𝑠
 and removes from 𝑠
 all its occurrences (the selected letter must occur in the string 𝑠
 at the moment of performing this action).
Polycarp performs this sequence of actions strictly in this order.

Note that after Polycarp finishes the actions, the string 𝑠
 will be empty and the string 𝑡
 will be equal to some value (that is undefined and depends on the order of removing).

E.g. consider 𝑠
="abacaba" so the actions may be performed as follows:

𝑡
="abacaba", the letter 'b' is selected, then 𝑠
="aacaa";
𝑡
="abacabaaacaa", the letter 'a' is selected, then 𝑠
="c";
𝑡
="abacabaaacaac", the letter 'c' is selected, then 𝑠
="" (the empty string).
You need to restore the initial value of the string 𝑠
 using only the final value of 𝑡
 and find the order of removing letters from 𝑠
.

### ideas
1. 在t中出现的字符，肯定是在s中出现的字符。因为着至少要操作这么多次
2. 在t的prefix，肯定包含s，s包含所有的字符，某个后缀包含除去m-1个字符，m-2。。。m-3，一直到1
3. 这个是必要条件
4. 通过这个可以判断出来，每个字符的删除顺序
5. 怎么判断一定是可以的？充分条件是什么？
6. 在删除最后一个字符前，比如c，它的出现次数是多少能判断出来吗？可以的
7. 它就是sum / m, 下一个是 sum / (m - 1) 如果不能整除，就是有问题的