You are given three strings: 𝑎
, 𝑏
, and 𝑐
, consisting of lowercase Latin letters. The string 𝑐
 was obtained in the following way:

At each step, either string 𝑎
 or string 𝑏
 was randomly chosen, and the first character of the chosen string was removed from it and appended to the end of string 𝑐
, until one of the strings ran out. After that, the remaining characters of the non-empty string were added to the end of 𝑐
.
Then, a certain number of characters in string 𝑐
 were randomly changed.
For example, from the strings 𝑎=abra
 and 𝑏=cada
, without character replacements, the strings caabdraa
, abracada
, acadabra
 could be obtained.


### ideas
1. dp[i][j] 表示以a[:i] + b[:j]得到的c[:i+j]的最小replacements
2. dp[i][j] = min(dp[i-1][j] + 1 - c[i+j] == a[i], dp[i][j-1] + 1 - c[i+j] == b[j])