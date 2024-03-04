Vlad found a string 𝑠
consisting of 𝑛
lowercase Latin letters, and he wants to make it as short as possible.

To do this, he can remove any pair of adjacent characters from 𝑠
any number of times, provided they are different. For example, if 𝑠
=racoon, then by removing one pair of characters he can obtain the strings coon, roon, raon, and raco, but he cannot
obtain racn (because the removed letters were the same) or rcon (because the removed letters were not adjacent).

What is the minimum length Vlad can achieve by applying any number of deletions?

### thoughts

1. 如果最多的字符不超过1半，那么最后肯定剩下0（或者1）