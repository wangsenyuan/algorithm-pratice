A regular bracket sequence is a bracket sequence that can be transformed into a correct arithmetic expression by inserting characters '1' and '+' between the original characters of the sequence. For example:

bracket sequences "()()" and "(())" are regular (the resulting expressions are: "(1)+(1)" and "((1+1)+1)");
bracket sequences ")(", "(" and ")" are not.
Let's define the inverse of the bracket sequence as follows: replace all brackets '(' with ')', and vice versa (all brackets ')' with '('). For example, strings "()((" and ")())" are inverses of each other.

You are given a regular bracket sequence 𝑠
. Calculate the number of pairs of integers (𝑙,𝑟)
 (1≤𝑙≤𝑟≤|𝑠|
) such that if you replace the substring of 𝑠
 from the 𝑙
-th character to the 𝑟
-th character (inclusive) with its inverse, 𝑠
 will still be a regular bracket sequence.

 ### ideas
 1. 需要找到regular的特征
 2. 假设dp[i]表示从左到i为止，level的值，fp[i]表示从右到左，到i为止的值
 3. dp[i] >= 0 && fp[j] >= 0, 
 4. replace后，这个条件仍然要满足
 5. 假设选中了l...r， 那么对于其中的i (l <= i <= r)
 6. 这个中间左右括号的值要相同，各一半
 7. 假设这个中间的最小的dp[i], 
 8. 对于给定的r
 9. 这样的l是符合条件的，在[l..r]中间左右括号的数量相同
 10. 且假设[l..i]的左括号-右的数量最大，则这个数字<= dp[l-1]
 11. 且假设[i...r]的右括号-左的数量最大，则这个数字 <= fp[r+1]
 12. 如果[l...r]满足条件，那它里面的是否都满足条件呢？
 13. 如果选定i，那么另外一个是不是一定是i+1？
 14. 如果在l-1处有3个左括号没有匹配
 15. 因为[l..r]中间的左右括号相同，那么意味着在[r+1]后面要有3个右括号
 16. 且假设[l...i]的左括号的剩余为x，那么在[j...r]中右括号的剩余，肯定也为x，且j = i + 1
 17. 这是因为，在[l...r]中间的左右括号是相同的，如果在i处剩余了x个左括号，那么必须在后面消耗掉
 18. 假设在[l...i]的左括号剩余为x，且dp[l-1] >= x, 那么所有那些[i+1...r], fp[r+1] >= x 的位置
 19. 进步了一点点。
 20. 考虑上面的i+1， x = fp[i+1] - fp[r+1]
 21. 是不是只要满足 dp[l-1] >= x 就可以？x = fp[i+1] - fp[r+1]
 22. 就是在l后后面找到位置i,且fp[i+1] - fp[r+1] = x部分？
 23. 在给定l的时候， fp[r+1]是确定的，那么fp[i+1] - x 是确定的
 24. dp[l-1] >= fp[i+1] - fp[r+1]
 25. 2 * dp[l-1] >= fp[i+1] 就可以了？