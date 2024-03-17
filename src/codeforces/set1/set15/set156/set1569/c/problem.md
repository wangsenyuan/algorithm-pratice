𝑛
people gathered to hold a jury meeting of the upcoming competition, the 𝑖
-th member of the jury came up with 𝑎𝑖
tasks, which they want to share with each other.

First, the jury decides on the order which they will follow while describing the tasks. Let that be a permutation 𝑝
of numbers from 1
to 𝑛
(an array of size 𝑛
where each integer from 1
to 𝑛
occurs exactly once).

Then the discussion goes as follows:

If a jury member 𝑝1
has some tasks left to tell, then they tell one task to others. Otherwise, they are skipped.
If a jury member 𝑝2
has some tasks left to tell, then they tell one task to others. Otherwise, they are skipped.
...
If a jury member 𝑝𝑛
has some tasks left to tell, then they tell one task to others. Otherwise, they are skipped.
If there are still members with tasks left, then the process repeats from the start. Otherwise, the discussion ends.
A permutation 𝑝
is nice if none of the jury members tell two or more of their own tasks in a row.

Count the number of nice permutations. The answer may be really large, so print it modulo 998244353
.

### ideas

1. 没有一个人连续的讲
2. 假设有x个故事讲，他讲完后，大家都轮了一次，下次轮到i是n-1次后，什么时候会出现连续的情况呢？
3. 就是要按照从大往小排，相同的可以随便排