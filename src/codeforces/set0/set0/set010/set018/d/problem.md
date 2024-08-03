Last year Bob earned by selling memory sticks. During each of n days of his work one of the two following events took place:

A customer came to Bob and asked to sell him a 2x MB memory stick. If Bob had such a stick, he sold it and got 2x berllars.
Bob won some programming competition and got a 2x MB memory stick as a prize. Bob could choose whether to present this memory stick to one of his friends, or keep it.
Bob never kept more than one memory stick, as he feared to mix up their capacities, and deceive a customer unintentionally. It is also known that for each memory stick capacity there was at most one customer, who wanted to buy that memory stick. Now, knowing all the customers' demands and all the prizes won at programming competitions during the last n days, Bob wants to know, how much money he could have earned, if he had acted optimally.

Input
The first input line contains number n (1 ≤ n ≤ 5000) — amount of Bob's working days. The following n lines contain the description of the days. Line sell x stands for a day when a customer came to Bob to buy a 2x MB memory stick (0 ≤ x ≤ 2000). It's guaranteed that for each x there is not more than one line sell x. Line win x stands for a day when Bob won a 2x MB memory stick (0 ≤ x ≤ 2000).

Output
Output the maximum possible earnings for Bob in berllars, that he would have had if he had known all the events beforehand. Don't forget, please, that Bob can't keep more than one memory stick at a time.

### ideas
1. dp[i] = 在后缀上能够获得的最大值
2. dp[i] = dp[j] + x , if it is win x, find the last sell x pos j
3.      or dp[i+1]