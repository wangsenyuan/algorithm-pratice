The next "Data Structures and Algorithms" lesson will be about Longest Increasing Subsequence (LIS for short) of a sequence. For better understanding, Nam decided to learn it a few days before the lesson.

Nam created a sequence a consisting of n (1 ≤ n ≤ 105) elements a1, a2, ..., an (1 ≤ ai ≤ 105). A subsequence ai1, ai2, ..., aik where 1 ≤ i1 < i2 < ... < ik ≤ n is called increasing if ai1 < ai2 < ai3 < ... < aik. An increasing subsequence is called longest if it has maximum length among all increasing subsequences.

Nam realizes that a sequence may have several longest increasing subsequences. Hence, he divides all indexes i (1 ≤ i ≤ n), into three groups:

group of all i such that ai belongs to no longest increasing subsequences.
group of all i such that ai belongs to at least one but not every longest increasing subsequence.
group of all i such that ai belongs to every longest increasing subsequence.
Since the number of longest increasing subsequences of a may be very large, categorizing process is very difficult. Your task is to help him finish this job.

