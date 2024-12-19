Thanos wants to destroy the avengers base, but he needs to destroy the avengers along with their base.

Let we represent their base with an array, where each position can be occupied by many avengers, but one avenger can occupy only one position. Length of their base is a perfect power of 2
. Thanos wants to destroy the base using minimum power. He starts with the whole base and in one step he can do either of following:

if the current length is at least 2
, divide the base into 2
 equal halves and destroy them separately, or
burn the current base. If it contains no avenger in it, it takes 𝐴
 amount of power, otherwise it takes his 𝐵⋅𝑛𝑎⋅𝑙
 amount of power, where 𝑛𝑎
 is the number of avengers and 𝑙
 is the length of the current base.
Output the minimum power needed by Thanos to destroy the avengers' base.

### ideas
1. just brute force and add some cache