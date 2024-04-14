Polycarp has prepared 𝑛
competitive programming problems. The topic of the 𝑖
-th problem is 𝑎𝑖
, and some problems' topics may coincide.

Polycarp has to host several thematic contests. All problems in each contest should have the same topic, and all
contests should have pairwise distinct topics. He may not use all the problems. It is possible that there are no
contests for some topics.

Polycarp wants to host competitions on consecutive days, one contest per day. Polycarp wants to host a set of contests
in such a way that:

number of problems in each contest is exactly twice as much as in the previous contest (one day ago), the first contest
can contain arbitrary number of problems;
the total number of problems in all the contests should be maximized.
Your task is to calculate the maximum number of problems in the set of thematic contests. Note, that you should not
maximize the number of contests.

### ideas

1. sort by freq
2. last one is determined