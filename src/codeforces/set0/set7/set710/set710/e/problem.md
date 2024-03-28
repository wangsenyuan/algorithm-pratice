zscoder wants to generate an input file for some programming competition problem.

His input is a string consisting of n letters 'a'. He is too lazy to write a generator so he will manually generate the
input in a text editor.

Initially, the text editor is empty. It takes him x seconds to insert or delete a letter 'a' from the text file and y
seconds to copy the contents of the entire text file, and duplicate it.

zscoder wants to find the minimum amount of time needed for him to create the input file of exactly n letters 'a'. Help
him to determine the amount of time needed to generate the input.

### ideas

1. dp[i] = dp[i/2] + y (if i % 2 = 0) min dp[i-1] + x
2. 但这里还有一些需要考虑的内容，就是超出部分
3. 如果 j * 2 > i => dp[i] = dp[j] + y + (j* 2 - i) * x
4. 所以还需知道 [i/2...i-1]中间的dp[j] - j * 2 的最小值