Takahashi is hosting an sports meet. There are 
N people who will participate. These people are conveniently numbered 
1 through 
N. Also, there are 
M options of sports for this event. These sports are numbered 
1 through 
M. Among these options, Takahashi will select one or more sports (possibly all) to be played in the event.

Takahashi knows that Person 
i's 
j-th favorite sport is Sport 
A 
ij
​
 . Each person will only participate in his/her most favorite sport among the ones that are actually played in the event, and will not participate in the other sports.

Takahashi is worried that one of the sports will attract too many people. Therefore, he would like to carefully select sports to be played so that the number of the participants in the sport with the largest number of participants is minimized. Find the minimum possible number of the participants in the sport with the largest number of participants.



### ideas
1. 假设选择了一组活动（sports), 那么每个选手，会选择其中，自己最喜欢的那项参加
2. 目标是，使得那些被最多人选择的活动的参与人数最少
3. 这个可以二分吗？限制最多可以被选择x次，
4. 似乎不大行。主要是选择了sport i后，可能会影响到之前的对sport j的选择（如果没有选择i，那么有一些人会选择j）
5. 但是选择了j后，部分选择i的人会重新选择j.
6. 比如在检查i的时候，发现没法满足条件（超过了x)，但是在选了j后，就满足了条件(<= x)
7. 用dp？
8. dp[i][j]表示前i个人，在前j个sports中，能够达到的最优解
9. 现在考虑第j+1个sports
10. 如果这个j+1个项目被忽略掉, dp[i][j+1] = dp[i][j] 
11. 否则的话，那么j+1个项目，在前i个人中，排在第一位的情况下，假设有w个人
12. 但是没法直接用 max(dp[i][j], w)来处理（因为，有可能前面有些任务的优先级比j+1要高）
13. 在不清楚它们的选择（与否）的情况下，没法继续
14. 感觉有个可以贪心的地方
15. 假设所有的都选择了，这个时候，可以计算出，被选择最大的那个sports， 假设它是i，和它的值f(i)
16. 那么在不选择它的时候，应该会有一个新的sport j,和 f(j) 如果f(j) < f(i), 那么不选择i是合理的
17. 而f(j)又可以更一步的变成，选择它的结果，和不选择它的结果 