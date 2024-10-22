You are planning your trip to Saint Petersburg. After doing some calculations, you estimated that you will have to spend k
 rubles each day you stay in Saint Petersburg — you have to rent a flat, to eat at some local cafe, et cetera. So, if the day of your arrival is L
, and the day of your departure is R
, you will have to spend k(R−L+1)
 rubles in Saint Petersburg.

You don't want to spend a lot of money on your trip, so you decided to work in Saint Petersburg during your trip. There are n
 available projects numbered from 1
 to n
, the i
-th of them lasts from the day li
 to the day ri
 inclusive. If you choose to participate in the i
-th project, then you have to stay and work in Saint Petersburg for the entire time this project lasts, but you get paid pi
 rubles for completing it.

Now you want to come up with an optimal trip plan: you have to choose the day of arrival L
, the day of departure R
 and the set of projects S
 to participate in so that all the following conditions are met:

your trip lasts at least one day (formally, R≥L
);
you stay in Saint Petersburg for the duration of every project you have chosen (formally, for each s∈S
 L≤ls
 and R≥rs
);
your total profit is strictly positive and maximum possible (formally, you have to maximize the value of ∑s∈Sps−k(R−L+1)
, and this value should be positive).
You may assume that no matter how many projects you choose, you will still have time and ability to participate in all of them, even if they overlap.

### ideas
1. 把收益算入 ri的日期，dp[i]表示到第i天时的总收益
2. 那么答案 = dp[r] - dp[l] - g(从l及之前开始的，在l后面结束的收益) - (r - l) * k
3. let R[l]表示在l前开始，但是最远的结束日期
4. -dp[l] - g(l...R[l]) + l * k + (dp[r] - r * k) 所以要找R[l]后面的最大的dp[r] - r * k 的值