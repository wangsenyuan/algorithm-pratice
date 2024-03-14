A string 𝑎=𝑎1𝑎2…𝑎𝑛
is called even if it consists of a concatenation (joining) of strings of length 2
consisting of the same characters. In other words, a string 𝑎
is even if two conditions are satisfied at the same time:

its length 𝑛
is even;
for all odd 𝑖
(1≤𝑖≤𝑛−1
), 𝑎𝑖=𝑎𝑖+1
is satisfied.
For example, the following strings are even: "" (empty string), "tt", "aabb", "oooo", and "ttrrrroouuuuuuuukk". The
following strings are not even: "aaa", "abab" and "abba".

Given a string 𝑠
consisting of lowercase Latin letters. Find the minimum number of characters to remove from the string 𝑠
to make it even. The deleted characters do not have to be consecutive.

