The coronation of King Berl XXII is soon! The whole royal family, including 𝑛
 daughters of Berl XXII, will be present.

The King has ordered his jeweler to assemble 𝑛
 beautiful necklaces, so each of the princesses could wear exactly one necklace during the ceremony — and now these necklaces are finished. Each necklace consists of 𝑚
 gems attached to a gold chain. There are two types of gems used in the necklaces — emeralds and sapphires. So, each necklace can be represented by a sequence of 𝑚
 gems (listed from left to right), and each gem is either an emerald or a sapphire. Formally, the 𝑖
-th necklace can be represented by a binary string 𝑠𝑖
 of length 𝑚
; if the 𝑗
-th character of 𝑠𝑖
 is 0, then the 𝑗
-th gem in the 𝑖
-th necklace is an emerald; otherwise, this gem is a sapphire.

Now, looking at the necklaces, the King is afraid that some of his daughters may envy the other daughters' necklaces. He wants all necklaces to look similar. Two necklaces are considered similar if there are at least 𝑘
 positions where these necklaces contain the same type of gems.

For example, if there is a necklace represented by a sequence 01010111 and a necklace represented by a sequence 01100000, then there are 3
 positions where these necklaces contain the same type of gems (both first gems are emeralds, both second gems are sapphires, and both fifth gems are emeralds). So if 𝑘=3
, these necklaces are similar, and if 𝑘=4
, they are not similar.

The King thinks that if two of his daughters notice that their necklaces are not similar, then they may have a conflict — and, obviously, he doesn't want any conflicts during the coronation! So Berl XXII wants to tell some of his daughters to wear their necklaces backward. If a necklace is worn backward, then the sequence of gems in this necklace is reversed. For example, if a necklace is represented by a sequence 01100, then, if worn backward, it would be represented by a sequence 00110. The King wants to find the minimum number of necklaces to be worn backward during the coronation so that there are no conflicts.

Berl XXII is too busy with preparation for the coronation, so he ordered you to resolve this issue for him. Help him — and he will give you a truly royal reward!


### ideas
1. n <= 50， m <= 50
2. 所有的项链，两两之间都要相似;
3. 如果要翻转，翻转的个数不会超过n/2， 否则翻转少的那部分即可
4. 所以，可以根据a, b之间是否需要翻转后，才能相似可以分成2组
5. 这里有3种情况， (a, b) 不需要翻转，就已经相似了；(a, b)翻转后，是相似的; (a, b)无论如何都不相似 
6. 如果是第三种情况，可以直接返回 -1
7. 如果是满足第一种情况，且不满足第二种情况， (a, b)必须在同一边
8. 如果是满足第二种情况，且不满足第一种情况， (a, b)必须在两边
9. 如果同时满足两种情况， 则没有关系
10. 如此一来，就变成了一个sat-2？但是问题是怎么最小化改变的情况呢？
11. 那么就是找到一个size = n的ssc吗？
12. 试试？