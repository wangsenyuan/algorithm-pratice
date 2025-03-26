Every student enrolled in the algorithms course is required to submit an assignment this week. The task is to implement an 𝑂(𝑛2)
-time algorithm to sort 𝑛
 given integers in non-decreasing order. Alice has already completed her assignment, and her implementation is shown below.

int alice_sort(int *s, int n){
  for(int i = 0; i < n; ++i){
    for(int j = i + 1; j < n; ++j){
      if(s[i] > s[j]){
        int swap = s[i];
        s[i] = s[j];
        s[j] = swap;
      }
    }
  }
  return 0;
}
While you have access to Alice's code, you prefer not to simply copy it. Instead, you want to use Alice's sorting function as a building block for your own solution. There are two ways as listed below you can utilize her function, but each of them can be applied at most once. The order in which these two operations are invoked can be arbitrary.

Prefix sort: choose a length 𝑖∈{1,2,…,𝑛}
 and call 𝚊𝚕𝚒𝚌𝚎𝚜𝚘𝚛𝚝(𝑠,𝑖)
. This sorts the first 𝑖
 elements in the array 𝑠
.
Suffix sort: choose a length 𝑖∈{1,2,…,𝑛}
 and call 𝚊𝚕𝚒𝚌𝚎𝚜𝚘𝚛𝚝(𝑠+𝑛−𝑖,𝑖)
. This sorts the last 𝑖
 elements in the array 𝑠
.
Due to the time complexity of the sorting algorithm, the cost of performing either a prefix or suffix sort is 𝑖2
, where 𝑖
 is the length of the chosen subarray. Your goal is to determine the minimum cost to sort the input array 𝑠
 of 𝑛
 integers in non-decreasing order using Alice's function, following the rules mentioned above.

For example, Let 𝑠=[3,2,5,5,4,1]
. We can first perform a suffix sort of length 4
, and the array becomes [3,2,1,4,5,5]
. Then, we perform a prefix sort of length 3
, and the array becomes [1,2,3,4,5,5]
, which is a sorted array. The cost is 42+32=25
. Here is another example, let 𝑠=[4,3,2,1]
. We can complete the sorting by performing only a prefix sort of length 4
, and the cost is 42=16
.

### solution

Let the range where the first (resp. second) operation is applied be [1, x] (resp. [y, n]).
• Guess y.
• If [1, x] and [y, n] are disjoint, find the x that is furthest from y such that the elements
within the range [x, y] are in the correct positions.
• Otherwse [1, x] and [y, n] intersect. Find the smallest x such that if you sort the first x
elements then the first y − 1 elements are in the correct positions.
• The above can be implemented in O(n) time. An O(n log n)-time algorithm also can pass
all the test cases within the time limit.