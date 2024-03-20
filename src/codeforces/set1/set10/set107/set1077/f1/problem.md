Vova likes pictures with kittens. The news feed in the social network he uses can be represented as an array of 𝑛
consecutive pictures (with kittens, of course). Vova likes all these pictures, but some are more beautiful than the
others: the 𝑖
-th picture has beauty 𝑎𝑖
.

Vova wants to repost exactly 𝑥
pictures in such a way that:

each segment of the news feed of at least 𝑘
consecutive pictures has at least one picture reposted by Vova;
the sum of beauty values of reposted pictures is maximum possible.
For example, if 𝑘=1
then Vova has to repost all the pictures in the news feed. If 𝑘=2
then Vova can skip some pictures, but between every pair of consecutive pictures Vova has to repost at least one of
them.

Your task is to calculate the maximum possible sum of values of reposted pictures if Vova follows conditions described
above, or say that there is no way to satisfy all conditions.

### ideas

1. dp[i][j][k] 表示到i为止，前面已经选择了j个，且最近的在i的前面k个位置（k = 0时，表示在i处进行了选择)