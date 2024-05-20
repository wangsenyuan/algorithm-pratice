Isaac begins his training. There are 𝑛
 running tracks available, and the 𝑖
-th track (1≤𝑖≤𝑛
) consists of 𝑎𝑖
 equal-length sections.

Given an integer 𝑢
 (1≤𝑢≤109
), finishing each section can increase Isaac's ability by a certain value, described as follows:

Finishing the 1
-st section increases Isaac's performance by 𝑢
.
Finishing the 2
-nd section increases Isaac's performance by 𝑢−1
.
Finishing the 3
-rd section increases Isaac's performance by 𝑢−2
.
…
Finishing the 𝑘
-th section (𝑘≥1
) increases Isaac's performance by 𝑢+1−𝑘
. (The value 𝑢+1−𝑘
 can be negative, which means finishing an extra section decreases Isaac's performance.)
You are also given an integer 𝑙
. You must choose an integer 𝑟
 such that 𝑙≤𝑟≤𝑛
 and Isaac will finish each section of each track 𝑙,𝑙+1,…,𝑟
 (that is, a total of ∑𝑟𝑖=𝑙𝑎𝑖=𝑎𝑙+𝑎𝑙+1+…+𝑎𝑟
 sections).

Answer the following question: what is the optimal 𝑟
 you can choose that the increase in Isaac's performance is maximum possible?

If there are multiple 𝑟
 that maximize the increase in Isaac's performance, output the smallest 𝑟
.

To increase the difficulty, you need to answer the question for 𝑞
 different values of 𝑙
 and 𝑢
.

### ideas

1. 显然，如果整个track的净收益是负值时，就没有必要继续下去
2. 对于给定的l，u， 当sum[r] > u时，它就开始进入负值了