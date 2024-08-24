You are given a string 𝑠[1…𝑛]
 consisting of lowercase Latin letters. It is guaranteed that 𝑛=2𝑘
 for some integer 𝑘≥0
.

The string 𝑠[1…𝑛]
 is called 𝑐
-good if at least one of the following three conditions is satisfied:

The length of 𝑠
 is 1
, and it consists of the character 𝑐
 (i.e. 𝑠1=𝑐
);
The length of 𝑠
 is greater than 1
, the first half of the string consists of only the character 𝑐
 (i.e. 𝑠1=𝑠2=⋯=𝑠𝑛2=𝑐
) and the second half of the string (i.e. the string 𝑠𝑛2+1𝑠𝑛2+2…𝑠𝑛
) is a (𝑐+1)
-good string;
The length of 𝑠
 is greater than 1
, the second half of the string consists of only the character 𝑐
 (i.e. 𝑠𝑛2+1=𝑠𝑛2+2=⋯=𝑠𝑛=𝑐
) and the first half of the string (i.e. the string 𝑠1𝑠2…𝑠𝑛2
) is a (𝑐+1)
-good string.
For example: "aabc" is 'a'-good, "ffgheeee" is 'e'-good.

In one move, you can choose one index 𝑖
 from 1
 to 𝑛
 and replace 𝑠𝑖
 with any lowercase Latin letter (any character from 'a' to 'z').

Your task is to find the minimum number of moves required to obtain an 'a'-good string from 𝑠
 (i.e. 𝑐
-good string for 𝑐=
 'a'). It is guaranteed that the answer always exists.

You have to answer 𝑡
 independent test cases.

Another example of an 'a'-good string is as follows. Consider the string 𝑠=
"cdbbaaaa". It is an 'a'-good string, because:

the second half of the string ("aaaa") consists of only the character 'a';
the first half of the string ("cdbb") is 'b'-good string, because:
the second half of the string ("bb") consists of only the character 'b';
the first half of the string ("cd") is 'c'-good string, because:
the first half of the string ("c") consists of only the character 'c';
the second half of the string ("d") is 'd'-good string.


### ideas
1. dp[i][k][x] 表示从i开始，长度为1<<k 的串需要修改多少次，变成 x-good的
2. dp[i][k][x] = min(dp[i][k-1][x] + dp[i+(1 << (k-1))][x+1], or dp[...])