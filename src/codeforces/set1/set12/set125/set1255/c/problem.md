The tournament consisted of 𝑛
 (𝑛≥5
) teams around the world. Before the tournament starts, Bob has made a prediction of the rankings of each team, from 1
-st to 𝑛
-th. After the final, he compared the prediction with the actual result and found out that the 𝑖
-th team according to his prediction ended up at the 𝑝𝑖
-th position (1≤𝑝𝑖≤𝑛
, all 𝑝𝑖
 are unique). In other words, 𝑝
 is a permutation of 1,2,…,𝑛
.

As Bob's favorite League player is the famous "3ga", he decided to write down every 3
 consecutive elements of the permutation 𝑝
. Formally, Bob created an array 𝑞
 of 𝑛−2
 triples, where 𝑞𝑖=(𝑝𝑖,𝑝𝑖+1,𝑝𝑖+2)
 for each 1≤𝑖≤𝑛−2
. Bob was very proud of his array, so he showed it to his friend Alice.

After learning of Bob's array, Alice declared that she could retrieve the permutation 𝑝
 even if Bob rearranges the elements of 𝑞
 and the elements within each triple. Of course, Bob did not believe in such magic, so he did just the same as above to see Alice's respond.

For example, if 𝑛=5
 and 𝑝=[1,4,2,3,5]
, then the original array 𝑞
 will be [(1,4,2),(4,2,3),(2,3,5)]
. Bob can then rearrange the numbers within each triple and the positions of the triples to get [(4,3,2),(2,3,5),(4,1,2)]
. Note that [(1,4,2),(4,2,2),(3,3,5)]
 is not a valid rearrangement of 𝑞
, as Bob is not allowed to swap numbers belong to different triples.

As Alice's friend, you know for sure that Alice was just trying to show off, so you decided to save her some face by giving her any permutation 𝑝
 that is consistent with the array 𝑞
 she was given.

 ### ideas
 1. 如果 (a, b, c), (b, c, d)
 2. 那么它们两个肯定是相邻的，
 3. 如果 (a, b, c) c只在一个里面出现，那么c要么在最前面，要么在最后面
 4. 可以选定在最前面