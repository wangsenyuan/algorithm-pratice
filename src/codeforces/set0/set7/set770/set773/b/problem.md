### ideas

1. 其他人的时间是固定的，只需要考虑是否解决了某道题
2. 增加x个人，分母 + x
3. 考虑三种情况
4. 第一个中A错，B对的情况，新账号只能提交错误答案（这种情况下，有可能会增加B的分数）
5. 第二种情况A错， B错的情况，也只能提交错误答案
6. 第三种A对，B错的情况，这种情况提交错误答案，对A肯定是有利的，但提交正确答案，却不一定不利（需要仔细计算）
7. 第四种情况A对B对的情况，
8. 如果A更快完成，那么提交错误答案更有利，通过错误答案，增加分数，因为A更快完成，减少的比例更少，进而分数更好
9. 如果B更快完成，那么提交正确答案更有利，通过正确答案，减低分数，因为B更快完成，减少的比例更大，

### solution

Dynamic problem scoring used to be used more often in Codeforces rounds, including some tournament rounds like VK Cup
2015 Finals.

Once you read the problem statement carefully, the problem itself isn't overly difficult.

Consider new accounts Vasya puts into play. Correct solutions to which problems does he submit from these new accounts?

If Vasya hasn't solved a problem, he can't submit correct solutions to it.
If Vasya has solved a problem which Petya hasn't solved, then clearly Vasya wants the maximum point value of this
problem to be as high as possible, thus it doesn't make sense to submit its solution from the new accounts.
Suppose Vasya solved the problem at minute v, Petya solved it at minute p and the problem's maximum point value is m,
then Vasya's and Petya's scores for this problem are m·(1 - v / 250) and m·(1 - p / 250), respectively. Let's denote the
difference between these values by d = m·(p - v)/ 250. Vasya wants to maximize this value.
If p - v is positive (that is, Vasya solved the problem faster than Petya), then d is maximized when m is maximized. To
maximize m, Vasya shouldn't submit correct solutions to this problem from the new accounts.
On the other hand, if p - v is negative (that is, Petya solved the problem faster than Vasya), then d is maximized when
m is minimized. To minimize m, Vasya should submit correct solutions to this problem from the new accounts.
Finally, if p - v is zero (that is, Petya and Vasya solved the problem at the same moment), then d = 0 for any value of
m, so it doesn't matter if Vasya submits correct solutions to this problem or not.
It follows from the above that Vasya should always do the same for all new accounts he puts into play.

Let's iterate over x — the number of new accounts Vasya puts into play, starting from 0. Then we can determine what
solutions Vasya should submit from these accounts using the reasoning above. Then we can calculate the maximum point
values of the problems, and then the number of points Vasya and Petya will score. If Vasya's score is higher than Petya'
s score, then the answer is x, otherwise we increase x by one and continue.

When do we stop? If Vasya submits solutions to a problem from the new accounts, then after putting at least n accounts
into play the maximum point value of this problem will reach 500 and won't change anymore. If Vasya doesn't, then after
putting at least 31n accounts into play the maximum point value of this problem will reach 3000 and won't change
anymore. Therefore, if x exceeds 31n, we can stop and output -1.

Note that we can't find the value of x using binary search due to the fact that Vasya can't submit solutions to the
problems he hasn't solved. That is, more accounts do not mean more profit. For example, consider the following test
case:

3
0 -1 -1 -1 -1
-1 0 -1 -1 -1
-1 0 -1 -1 -1
If Vasya doesn't use any new accounts, his score will be 1000, while Petya's score will be 500. If Vasya uses at least
61 accounts, both his and Petya's score will be 3000.