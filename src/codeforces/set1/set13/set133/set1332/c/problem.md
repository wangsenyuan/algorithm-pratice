Word 𝑠
 of length 𝑛
 is called 𝑘
-complete if

𝑠
 is a palindrome, i.e. 𝑠𝑖=𝑠𝑛+1−𝑖
 for all 1≤𝑖≤𝑛
;
𝑠
 has a period of 𝑘
, i.e. 𝑠𝑖=𝑠𝑘+𝑖
 for all 1≤𝑖≤𝑛−𝑘
.
For example, "abaaba" is a 3
-complete word, while "abccba" is not.

Bob is given a word 𝑠
 of length 𝑛
 consisting of only lowercase Latin letters and an integer 𝑘
, such that 𝑛
 is divisible by 𝑘
. He wants to convert 𝑠
 to any 𝑘
-complete word.

To do this Bob can choose some 𝑖
 (1≤𝑖≤𝑛
) and replace the letter at position 𝑖
 with some other lowercase Latin letter.

So now Bob wants to know the minimum number of letters he has to replace to convert 𝑠
 to any 𝑘
-complete word.

Note that Bob can do zero changes if the word 𝑠
 is already 𝑘
-complete.

You are required to answer 𝑡
 test cases independently.

 ### ideas
 1. a[i] = a[i+k]
 2. a[l] = a[r] => 推出那些必须是相同字符的部分