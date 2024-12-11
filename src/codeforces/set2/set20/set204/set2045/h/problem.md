You have a dictionary, which is a list of distinct words sorted in alphabetical order. Each word consists of uppercase English letters.

You want to print this dictionary. However, there is a bug with the printing system, and all words in the list are printed next to each other without any separators between words. Now, you ended up with a string 𝑆
 that is a concatenation of all the words in the dictionary in the listed order.

Your task is to reconstruct the dictionary by splitting 𝑆
 into one or more words. Note that the reconstructed dictionary must consist of distinct words sorted in alphabetical order. Furthermore, you want to maximize the number of words in the dictionary. If there are several possible dictionaries with the maximum number of words, you can choose any of them.

### ideas
1. 假设当前分配出串s, （如果是从右到左）那么必须保证s是唯一的，且s < 下一个串
2. 一个想法是， dp[i][j] 表示是否可以找出s[i...j]作为一个word，开始的子串
3. dp[i...j] = dp[j+1.?] 且 s[i...j] < s[i+1...]
4. 这里有3层迭代（i, j, 还有?), n * n * n
5. 不大行
6. 先对这样的字符子串s[i...j]进行排序
7. 这样有个好处是，不用对别字符，只需要对比他们的位置即可（这个可以用suffix array去处理）
8. 然后呢？s[i...j]的位置，假设是w，那么就是在j+1后面，找到w+1的位置？
9. 好像是可以的