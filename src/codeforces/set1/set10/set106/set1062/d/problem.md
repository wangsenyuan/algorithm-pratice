You are given a positive integer 𝑛
greater or equal to 2
. For every pair of integers 𝑎
and 𝑏
(2≤|𝑎|,|𝑏|≤𝑛
), you can transform 𝑎
into 𝑏
if and only if there exists an integer 𝑥
such that 1<|𝑥|
and (𝑎⋅𝑥=𝑏
or 𝑏⋅𝑥=𝑎
), where |𝑥|
denotes the absolute value of 𝑥
.

After such a transformation, your score increases by |𝑥|
points and you are not allowed to transform 𝑎
into 𝑏
nor 𝑏
into 𝑎
anymore.

Initially, you have a score of 0
. You can start at any integer and transform it as many times as you like. What is the maximum score you can achieve?

### ideas

1. abs(x) > 1