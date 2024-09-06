You are given a string 𝑠
 consisting of lowercase Latin letters. Let the length of 𝑠
 be |𝑠|
. You may perform several operations on this string.

In one operation, you can choose some index 𝑖
 and remove the 𝑖
-th character of 𝑠
 (𝑠𝑖
) if at least one of its adjacent characters is the previous letter in the Latin alphabet for 𝑠𝑖
. For example, the previous letter for b is a, the previous letter for s is r, the letter a has no previous letters. Note that after each removal the length of the string decreases by one. So, the index 𝑖
 should satisfy the condition 1≤𝑖≤|𝑠|
 during each operation.

For the character 𝑠𝑖
 adjacent characters are 𝑠𝑖−1
 and 𝑠𝑖+1
. The first and the last characters of 𝑠
 both have only one adjacent character (unless |𝑠|=1
).

Consider the following example. Let 𝑠=
 bacabcab.

During the first move, you can remove the first character 𝑠1=
 b because 𝑠2=
 a. Then the string becomes 𝑠=
 acabcab.
During the second move, you can remove the fifth character 𝑠5=
 c because 𝑠4=
 b. Then the string becomes 𝑠=
 acabab.
During the third move, you can remove the sixth character 𝑠6=
'b' because 𝑠5=
 a. Then the string becomes 𝑠=
 acaba.
During the fourth move, the only character you can remove is 𝑠4=
 b, because 𝑠3=
 a (or 𝑠5=
 a). The string becomes 𝑠=
 acaa and you cannot do anything with it.
Your task is to find the maximum possible number of characters you can remove if you choose the sequence of operations optimally.

### ideas
1. 显然要从z开始删除，这是因为，z不是任何字符的前一个字符，所以如果删除掉它，不会造成更坏的结果
2. 看样子这个策略不大对