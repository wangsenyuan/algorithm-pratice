Kevin used to get into Rio's Memories, and in Rio's Memories, a series of contests was once held. Kevin remembers all the participants and all the contest problems from that time, but he has forgotten the specific rounds, the distribution of problems, and the exact rankings.

There are 𝑚
 problems in total, with the 𝑖
-th problem having a difficulty of 𝑏𝑖
. Let each contest consist of 𝑘
 problems, resulting in a total of ⌊𝑚𝑘⌋
 contests. This means that you select exactly ⌊𝑚𝑘⌋⋅𝑘
 problems for the contests in any combination you want, with each problem being selected at most once, and the remaining 𝑚mod𝑘
 problems are left unused. For example, if 𝑚=17
 and 𝑘=3
, you should create exactly 5
 contests consisting of 3
 problems each, and exactly 2
 problems will be left unused.

There are 𝑛
 participants in the contests, with Kevin being the 1
-st participant. The 𝑖
-th participant has a rating of 𝑎𝑖
. During the contests, each participant solves all problems with a difficulty not exceeding their rating, meaning the 𝑖
-th participant solves the 𝑗
-th problem if and only if 𝑎𝑖≥𝑏𝑗
. In each contest, Kevin's rank is one plus the number of participants who solve more problems than he does.

For each 𝑘=1,2,…,𝑚
, Kevin wants to know the minimum sum of his ranks across all ⌊𝑚𝑘⌋
 contests. In other words, for some value of 𝑘
, after selecting the problems for each contest, you calculate the rank of Kevin in each contest and sum up these ranks over all ⌊𝑚𝑘⌋
 contests. Your goal is to minimize this value.

Note that contests for different values of 𝑘
 are independent. It means that for different values of 𝑘
, you can select the distribution of problems into the contests independently.