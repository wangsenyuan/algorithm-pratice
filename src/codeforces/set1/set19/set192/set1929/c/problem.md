Sasha decided to give his girlfriend the best handbag, but unfortunately for Sasha, it is very expensive. Therefore, Sasha wants to earn it. After looking at earning tips on the internet, he decided to go to the casino.

Sasha knows that the casino operates under the following rules. If Sasha places a bet of 𝑦
 coins (where 𝑦
 is a positive integer), then in case of winning, he will receive 𝑦⋅𝑘
 coins (i.e., his number of coins will increase by 𝑦⋅(𝑘−1)
). And in case of losing, he will lose the entire bet amount (i.e., his number of coins will decrease by 𝑦
).

Note that the bet amount must always be a positive (>0
) integer and cannot exceed Sasha's current number of coins.

Sasha also knows that there is a promotion at the casino: he cannot lose more than 𝑥
 times in a row.

Initially, Sasha has 𝑎
 coins. He wonders whether he can place bets such that he is guaranteed to win any number of coins. In other words, is it true that for any integer 𝑛
, Sasha can make bets so that for any outcome that does not contradict the rules described above, at some moment of time he will have at least 𝑛
 coins.

 ### ideas
 1. 可以假设sasha在x的循环结中处理，在前x-1次，它投入1元，然后在最后一次，投入所有的(cur - (x - 1)) * k
 2. 然后重复
 3. 如果 y < x => false (永远都不到)
 4. y = x, 如果 k <= y, 同样永远也达不到
 5. 否则有限次后，肯定能达到
 6. 这个想法是不对的
 7. 因为在这x次中的，任何一次，他都可能会赢
 8. 如果这个时候，bet的是1，那么收益就是k - 1 元
 9. 如果 k - 1 > x, 那么就可以得到无穷多的收益 (a > x)
 10. 但是
  
### solution

Let's notice that the condition that we can achieve arbitrarily large values means that we need to guarantee at least a +1
 to our coins. At the very first win. In this case, we can repeat this strategy indefinitely.

Also, let's notice that if we have lost a total of 𝑧
 before, then in the next round we need to bet 𝑦
 such that 𝑦⋅(𝑘−1)>𝑧
, because otherwise the casino can give us a win. In this case, the condition of not losing more than 𝑥
 times in a row will disappear, and we will end up in the negative. Therefore, the tactic is not optimal.

Therefore, the solution is as follows: we bet 1
 at first, then we bet the minimum number such that the win covers our loss. And if we have enough to make such a bet for 𝑥+1
, then the casino must end up in the negative, otherwise we cannot win.

So the solution is in 𝑂(𝑥)
 time complexity, where we simply calculate these values in a loop.