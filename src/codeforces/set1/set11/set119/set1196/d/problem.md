You are given a string 𝑠
 consisting of 𝑛
 characters, each character is 'R', 'G' or 'B'.

You are also given an integer 𝑘
. Your task is to change the minimum number of characters in the initial string 𝑠
 so that after the changes there will be a string of length 𝑘
 that is a substring of 𝑠
, and is also a substring of the infinite string "RGBRGBRGB ...".

A string 𝑎
 is a substring of string 𝑏
 if there exists a positive integer 𝑖
 such that 𝑎1=𝑏𝑖
, 𝑎2=𝑏𝑖+1
, 𝑎3=𝑏𝑖+2
, ..., 𝑎|𝑎|=𝑏𝑖+|𝑎|−1
. For example, strings "GBRG", "B", "BR" are substrings of the infinite string "RGBRGBRGB ..." while "GR", "RGR" and "GGG" are not.

You have to answer 𝑞
 independent queries.

 ### ideas
 1. RGBRGB....  这个模式是固定的
 2. 也就是说要在S中找出一段k长度的子串，它里面改变最少的情况下，满足 RGBRGB... 或者 RGBGBR... 或者 BRGBRG ... 的模式
 3. dp[i][x] 表示s[i]和r,g,b匹配时的最少改变量
 4. 然后找到 i - k + 1 的位置， 计算出对应的y,然后就可以计算出来了