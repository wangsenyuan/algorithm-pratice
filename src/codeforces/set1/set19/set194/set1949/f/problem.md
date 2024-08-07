You are the developer of a dating app which ignores gender completely. The app has 𝑛
 users, indexed from 1
 to 𝑛
. Each user's profile features a list of the activities they enjoy doing. There are 𝑚
 possible activities, indexed from 1
 to 𝑚
.

A match between two users is good if they share at least one activity and, at the same time, both of them like at least one activity that the other user does not like.

Find a good match if it exists.

Input
The first line contains two integers 𝑛
 and 𝑚
 (2≤𝑛≤200000
, 1≤𝑚≤106
) — the number of users and the number of activities.

Each of the following 𝑛
 lines contains a number 𝑘𝑖
 (0≤𝑘𝑖≤𝑚
) — the number of activities that user 𝑖
 likes — followed by 𝑘𝑖
 distinct integers from 1
 to 𝑚
 — the activities user 𝑖
 likes.

It is guaranteed that 𝑘1+𝑘2+⋯+𝑘𝑛
 does not exceed 106
.

Output
Print 𝚈𝙴𝚂
 if a good match exists. Otherwise, print 𝙽𝙾
.

If a good match exists, on the next line print two integers — the indexes of two users that make a match.

### ideas
1. use bitset? 
2. hash？
3. 假设把他们排序，相邻两行，如果一行完全包含另外一行的记录，那么是肯定不行的
4. 如果存在答案，是不是肯定是相邻的两行？
5. [1, 2], [1, 2, 3], [2, 3], [2, 3, 4]
6. 这个例子可以看出，上面的结论不成立
7. 如果freq[1] = m， 直接排除掉1
8. 否则的话，可以把1当成一个不同的爱好
9. 然后可以将他们分成两组，一组爱好1，一组不爱好1，
10. 然后记录这两组，是否有某个兴趣i是共同的，如果没有，舍弃它，然后继续
11. 如果有，那么答案就找到了
12. m * n / 64
13. 如果i包含j, 那么让 j -> i, 如果j有两个children，那么他的两个children就是答案
14. tree的节点应该是兴趣，然后某个人的组合，相当于把这些节点连起来
15. 。。。
16. 不会～
17. 