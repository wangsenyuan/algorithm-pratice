Little Petya likes points a lot. Recently his mom has presented him n points lying on the line OX. Now Petya is wondering in how many ways he can choose three distinct points so that the distance between the two farthest of them doesn't exceed d.

Note that the order of the points inside the group of three chosen points doesn't matter.


### ideas
1. i < j < k, and a[i] >= a[j] - d, and a[j] >= a[k] - d
2. let dp[j] = count of k, that a[k] <= a[j] + d
3. fp[i] = sum of dp[j], that a[j] <= a[i] + d 
4. 理解错了，是 a[k] <= a[i] + d