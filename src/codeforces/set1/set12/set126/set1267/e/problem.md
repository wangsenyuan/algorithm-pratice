Byteburg Senate elections are coming. Usually "United Byteland", the ruling Byteland party, takes all the seats in the Senate to ensure stability and sustainable development. But this year there is one opposition candidate in one of the constituencies. Even one opposition member can disturb the stability in the Senate, so the head of the Party asks you to ensure that the opposition candidate will not be elected.

There are 𝑛
 candidates, numbered from 1 to 𝑛
. Candidate 𝑛
 is the opposition candidate. There are 𝑚
 polling stations in the constituency, numbered from 1 to 𝑚
. You know the number of votes cast for each candidate at each polling station. The only thing you can do to prevent the election of the opposition candidate is to cancel the election results at some polling stations. The opposition candidate will be elected if the sum of the votes cast in their favor at all non-canceled stations will be strictly greater than the analogous sum for every other candidate.

Your task is to prevent the election of the opposition candidate by canceling the election results at the minimal possible number of polling stations. Notice that solution always exists, because if you cancel the elections at all polling stations, the number of votes for each candidate will be 0, and the opposition candidate will not be elected.

### ideas
1. 要取消的那些行，肯定是a[i][n] 最大的那些（相比同一行内，最大）
2. 不大对。应该需要一个i，尽量的让它的总分数 >= n
3. 