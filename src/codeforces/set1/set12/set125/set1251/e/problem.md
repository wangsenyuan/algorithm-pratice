Now elections are held in Berland and you want to win them. More precisely, you want everyone to vote for you.

There are 𝑛
 voters, and two ways to convince each of them to vote for you. The first way to convince the 𝑖
-th voter is to pay him 𝑝𝑖
 coins. The second way is to make 𝑚𝑖
 other voters vote for you, and the 𝑖
-th voter will vote for free.

Moreover, the process of such voting takes place in several steps. For example, if there are five voters with 𝑚1=1
, 𝑚2=2
, 𝑚3=2
, 𝑚4=4
, 𝑚5=5
, then you can buy the vote of the fifth voter, and eventually everyone will vote for you. Set of people voting for you will change as follows: 5→1,5→1,2,3,5→1,2,3,4,5
.

Calculate the minimum number of coins you have to spend so that everyone votes for you.


### ideas
1. 得仔细思考一下。
2. 这里必须分两个阶段，第一个阶段，进行贿选（通过付钱来获得m个voters)
3. 第二个阶段，看是否能够在贿选m个人后，能够顺利拿到所有的选票
4. 但这里还是有两个问题
5. a. 贿选时，当然希望花的钱越少越好，但是存在一种情况，同样是贿选m个人，如果把a替换成b（花更多的钱）就能成功
6. b. 当然上面的情况也不一定必然出现
7. 或者反过来考虑，在贿选m个人的情况下，仍然能够完成时，可以节省的最多的钱
8. 那么第m+1个人，他必须满足want[i] <= m
9. 然后从中选择最贵的那个人，然后在 want[i] <= m+1 个人里面， 选择最贵的那个。
10. 依次类推