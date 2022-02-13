The string B will appear in string X as a substring if there is a subsequence B in the string X. So corresponding to a particular subsequence B, we will have the delete the unwanted characters in between the first and last character to obtain B as a substring. We have to pick such a subsequence in which the difference between the position of the first and last character is minimum, thereby leading to the minimum deletions from X.

If we fix the position of the first character of B in X(in this case ‘a’), the next character should be chosen as near as possible to minimize the deleted characters. If we proceed in the above manner the whole subsequence gets fixed and hence the position of characters of B.
So we want to fix the first character at every position possible and check the difference b/w the first and last character of the resultant subsequence and pick the one with the least difference, hence the minimum deletions. If we carry out these operations individually and directly we would require O(n) time at each step giving time limit errors.

To carry out these operations optimally we have to make some calculations and store some info prior to this step. In the naive approach, we calculate the position of the next element every time and repeat this step. By invoking our knowledge of binary lifting, we try to calculate the position of the next to next element, the 4th next element, and so on. This can be easily calculated using the dynamic programming technique.

After storing these values we can calculate the (m-1)th next element from each starting point using the binary lifting technique and pick the subsequence that requires the minimum deletions.

