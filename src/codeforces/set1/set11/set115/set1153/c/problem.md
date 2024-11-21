Serval soon said goodbye to Japari kindergarten, and began his life in Japari Primary School.

In his favorite math class, the teacher taught him the following interesting definitions.

A parenthesis sequence is a string, containing only characters "(" and ")".

A correct parenthesis sequence is a parenthesis sequence that can be transformed into a correct arithmetic expression by inserting characters "1" and "+" between the original characters of the sequence. For example, parenthesis sequences "()()", "(())" are correct (the resulting expressions are: "(1+1)+(1+1)", "((1+1)+1)"), while ")(" and ")" are not. Note that the empty string is a correct parenthesis sequence by definition.

We define that |𝑠|
 as the length of string 𝑠
. A strict prefix 𝑠[1…𝑙]
 (1≤𝑙<|𝑠|)
 of a string 𝑠=𝑠1𝑠2…𝑠|𝑠|
 is string 𝑠1𝑠2…𝑠𝑙
. Note that the empty string and the whole string are not strict prefixes of any string by the definition.

Having learned these definitions, he comes up with a new problem. He writes down a string 𝑠
 containing only characters "(", ")" and "?". And what he is going to do, is to replace each of the "?" in 𝑠
 independently by one of "(" and ")" to make all strict prefixes of the new sequence not a correct parenthesis sequence, while the new sequence should be a correct parenthesis sequence.

After all, he is just a primary school student so this problem is too hard for him to solve. As his best friend, can you help him to replace the question marks? If there are many solutions, any of them is acceptable.


### ideas
1. 最终的s是平衡的，所有之前的都是不平衡的，但是不能到达 <= 0
2. 考虑一个 ？，如果要把它变成(, 那么应该是越早越好
3. 如果把它变成), 那么就是越晚越好
4. 因为它在左边时，可以增加level，在右边时减少level