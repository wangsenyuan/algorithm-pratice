In an English class Nick had nothing to do at all, and remembered about wonderful strings called palindromes. We should remind you that a string is called a palindrome if it can be read the same way both from left to right and from right to left. Here are examples of such strings: «eye», «pop», «level», «aba», «deed», «racecar», «rotor», «madam».

Nick started to look carefully for all palindromes in the text that they were reading in the class. For each occurrence of each palindrome in the text he wrote a pair — the position of the beginning and the position of the ending of this occurrence in the text. Nick called each occurrence of each palindrome he found in the text subpalindrome. When he found all the subpalindromes, he decided to find out how many different pairs among these subpalindromes cross. Two subpalindromes cross if they cover common positions in the text. No palindrome can cross itself.

Let's look at the actions, performed by Nick, by the example of text «babb». At first he wrote out all subpalindromes:

• «b» — 1..1
• «bab» — 1..3
• «a» — 2..2
• «b» — 3..3
• «bb» — 3..4
• «b» — 4..4
Then Nick counted the amount of different pairs among these subpalindromes that cross. These pairs were six:

1. 1..1 cross with 1..3
2. 1..3 cross with 2..2
3. 1..3 cross with 3..3
4. 1..3 cross with 3..4
5. 3..3 cross with 3..4
6. 3..4 cross with 4..4
Since it's very exhausting to perform all the described actions manually, Nick asked you to help him and write a program that can find out the amount of different subpalindrome pairs that cross. Two subpalindrome pairs are regarded as different if one of the pairs contains a subpalindrome that the other does not.

### ideas
1. 一个长度为n的，它至少可以贡献 n + (n-2)/2个 （单个字符 + 内部的回文）
2. 假设一共有w个回文，任选两个的个数 = w * (w - 1) / 2
3. 然后 - 不重叠的，是不是就可以了？