Gaius Julius Caesar, a famous general, loved to line up his soldiers. Overall the army had n1 footmen and n2 horsemen.
Caesar thought that an arrangement is not beautiful if somewhere in the line there are strictly more that k1 footmen
standing successively one after another, or there are strictly more than k2 horsemen standing successively one after
another. Find the number of beautiful arrangements of the soldiers.

Note that all n1 + n2 warriors should be present at each arrangement. All footmen are considered indistinguishable among
themselves. Similarly, all horsemen are considered indistinguishable among themselves.

### ideas

1. dp[i][j][0/1][k] i个footman，和j个horseman，最后一个是footman（0）时
2. 且和最后一个人相同的长度为k时的计数