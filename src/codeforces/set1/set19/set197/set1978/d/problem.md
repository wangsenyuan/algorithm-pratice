Elections are taking place in Berland. There are 𝑛
 candidates participating in the elections, numbered from 1
 to 𝑛
. The 𝑖
-th candidate has 𝑎𝑖
 fans who will vote for him. Additionally, there are 𝑐
 people who are undecided about their favorite candidate, let's call them undecided. Undecided people will vote for the candidate with the lowest number.

The candidate who receives the maximum number of votes wins the elections, and if multiple candidates receive the same maximum number of votes, the candidate with the lowest number among them wins.

You found these elections too boring and predictable, so you decided to exclude some candidates from them. If you do not allow candidate number 𝑖
 to participate in the elections, all 𝑎𝑖
 of his fans will become undecided, and will vote for the candidate with the lowest number.

You are curious to find, for each 𝑖
 from 1
 to 𝑛
, the minimum number of candidates that need to be excluded from the elections for candidate number 𝑖
 to win the elections.


### ideas 
1. 假设希望a[i]胜选，且部分人退选了
2. 在没有人退选的情况下，a[i]应该是最大的，且 a[i] > a[0] + c (如果是0，特别处理)
3. 如果不满足这个条件，那必须有人退选， 这里分两种情况
4.  退选那些人的支持者到别人那里去了。
5.  退选那些人的支持者到i这里了
6. 到别人那里肯定是不行的，因为在无人退选的情况下, i都不能赢，更不要说，其他人还增加了额外的支持者
7. 所以，只能是第二种情况 
8. 貌似也可以用heap