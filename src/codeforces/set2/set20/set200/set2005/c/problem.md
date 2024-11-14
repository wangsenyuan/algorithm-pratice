Narek is too lazy to create the third problem of this contest. His friend Artur suggests that he should use ChatGPT. ChatGPT creates 𝑛
 problems, each consisting of 𝑚
 letters, so Narek has 𝑛
 strings. To make the problem harder, he combines the problems by selecting some of the 𝑛
 strings possibly none and concatenating them without altering their order. His chance of solving the problem is defined as 𝑠𝑐𝑜𝑟𝑒𝑛−𝑠𝑐𝑜𝑟𝑒𝑐
, where 𝑠𝑐𝑜𝑟𝑒𝑛
 is Narek's score and 𝑠𝑐𝑜𝑟𝑒𝑐
 is ChatGPT's score.

Narek calculates 𝑠𝑐𝑜𝑟𝑒𝑛
 by examining the selected string (he moves from left to right). He initially searches for the letter "𝚗"
, followed by "𝚊"
, "𝚛"
, "𝚎"
, and "𝚔"
. Upon finding all occurrences of these letters, he increments 𝑠𝑐𝑜𝑟𝑒𝑛
 by 5
 and resumes searching for "𝚗"
 again (he doesn't go back, and he just continues from where he left off).

After Narek finishes, ChatGPT scans through the array and increments 𝑠𝑐𝑜𝑟𝑒𝑐
 by 1
 for each letter "𝚗"
, "𝚊"
, "𝚛"
, "𝚎"
, or "𝚔"
 that Narek fails to utilize (note that if Narek fails to complete the last occurrence by finding all of the 5
 letters, then all of the letters he used are counted in ChatGPT's score 𝑠𝑐𝑜𝑟𝑒𝑐
, and Narek doesn't get any points if he doesn't finish finding all the 5 letters).

Narek aims to maximize the value of 𝑠𝑐𝑜𝑟𝑒𝑛−𝑠𝑐𝑜𝑟𝑒𝑐
 by selecting the most optimal subset of the initial strings.

 ### ideas
 1. 干烧脑子的一个题目
 2. n <= 1000, m <= 1000, n * m <= 1e6
 3. dp[i][j] 表示到i为止，且最后narek选择到的字符是0~5 (0表示，可以开始一个新的选择, 1选择了n,2选择了a...)
 4. 如果不选择s[i] => 保持不变
 5. 如果选择s[i], 那么就可以计算得分
 6. 这样子可以搞