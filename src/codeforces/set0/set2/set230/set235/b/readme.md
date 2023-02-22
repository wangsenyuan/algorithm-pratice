Let us take a deep look at how this score is calculated. For an n long 'O' block, it contributes n2 to the answer.

Let us reformat this problem a bit and consider the following alternative definition of the score: (1) For each two 'O' pair which there is no 'X' between them, they add 2 to the score. (2) For each 'O', it adds 1 to the score.

We claim that this new definition of the score is equivalent to the definition in the problem statement.

Proof of the claim: For an n long 'O' block, there are Cn2 pairs of 'O' in it and n 'O' in it. Note that 2Cn2 + n = n2.

So now we work with the new definition of the score. For each event(i,j) (which means s[i] and s[j] are 'O', and there is no 'X' between them). If event(i,j) happens, it adds 2 to the score.

So we only need to sum up the probabilities of all events and multiply them by 2, and our task becomes how to calculate the sum of probabilities of all the event(i,j). Let P(i,j) be the probability of event(i,j).

We can see that P(i,j) can be computed by . Then we denote P(j) as the sum of all event(i,j) for i<j. We have dp(0)=0 and dp(j)=(dp(j-1)+pj - 1)*pj