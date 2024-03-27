Molly Hooper has n different kinds of chemicals arranged in a line. Each of the chemicals has an affection value, The
i-th of them has affection value ai.

Molly wants Sherlock to fall in love with her. She intends to do this by mixing a contiguous segment of chemicals
together to make a love potion with total affection value as a non-negative integer power of k. Total affection value of
a continuous segment of chemicals is the sum of affection values of each chemical in that segment.

Help her to do so in finding the total number of such segments.

### ideas

1. a[i] 可以小于0
2. k也可以小于0，但是 abs(k) <= 10
3. 可不可以就对每位检查，有多少个位置满足sum[i..j] = k**p
4. 但这个上下届咋处理呢？
5. 50～

### solution

Hint: The number of possible powers will be less than 50 for any k.

Editorial: We are going to loop over all possible non-negative powers of k. Since the maximum possible value of subarray
sum can be 105 × 109 = 1014, there can be at most possible powers that can be the sum of subarrays. Let p[i] be the sum
of elements from index 1 to index i. We can precalculate p[i] in O(n) time complexity. (Prefix sum)

We will try to find the count of subarrays starting from index l. The sum of any such subarray ending at index r can be
written as p[r]- p[l - 1]. Now, where w is a power of k. We have to count the values of r ≥ l such that p[r]= w +
p[l - 1]. For this part, we can store the count of p[r] in a dictionary as we move from right of the array and use the
dictionary to find count of p[r] for corresponding p[l] and w.

PS: Do take care of a corner case for k = ± 1 while calculating powers of k.