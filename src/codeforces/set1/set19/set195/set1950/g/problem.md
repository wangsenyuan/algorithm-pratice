Vladislav has a playlist consisting of 𝑛
 songs, numbered from 1
 to 𝑛
. Song 𝑖
 has genre 𝑔𝑖
 and writer 𝑤𝑖
. He wants to make a playlist in such a way that every pair of adjacent songs either have the same writer or are from the same genre (or both). He calls such a playlist exciting. Both 𝑔𝑖
 and 𝑤𝑖
 are strings of length no more than 104
.

It might not always be possible to make an exciting playlist using all the songs, so the shuffling process occurs in two steps. First, some amount (possibly zero) of the songs are removed, and then the remaining songs in the playlist are rearranged to make it exciting.

Since Vladislav doesn't like when songs get removed from his playlist, he wants the making playlist to perform as few removals as possible. Help him find the minimum number of removals that need to be performed in order to be able to rearrange the rest of the songs to make the playlist exciting.

Input
The first line of the input contains a single integer 𝑡
 (1≤𝑡≤1000
) — the number of test cases. The description of test cases follows.

The first line of each test case contains a single integer 𝑛
 (1≤𝑛≤16
) — the number of songs in the original playlist.

Then 𝑛
 lines follow, the 𝑖
-th of which contains two strings of lowercase letters 𝑔𝑖
 and 𝑤𝑖
 (1≤|𝑔𝑖|,|𝑤𝑖|≤104
) — the genre and the writer of the 𝑖
-th song. Where |𝑔𝑖|
 and |𝑤𝑖|
 are lengths of the strings.

The sum of 2𝑛
 over all test cases does not exceed 216
.

The sum of |𝑔𝑖|+|𝑤𝑖|
 over all test cases does not exceed 4⋅105
.


### ideas
1. 需要将每个名称进行hash， double hash以避免冲突
2. dp[state][i]表示以state为集合，且以i结尾的playlist
3. 
