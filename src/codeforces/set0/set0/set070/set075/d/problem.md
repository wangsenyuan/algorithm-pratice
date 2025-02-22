Ahmed and Mostafa used to compete together in many programming contests for several years. Their coach Fegla asked them to solve one challenging problem, of course Ahmed was able to solve it but Mostafa couldn't.

This problem is similar to a standard problem but it has a different format and constraints.

In the standard problem you are given an array of integers, and you have to find one or more consecutive elements in this array where their sum is the maximum possible sum.

But in this problem you are given n small arrays, and you will create one big array from the concatenation of one or more instances of the small arrays (each small array could occur more than once). The big array will be given as an array of indexes (1-based) of the small arrays, and the concatenation should be done in the same order as in this array. Then you should apply the standard problem mentioned above on the resulting big array.

For example let's suppose that the small arrays are {1, 6, -2}, {3, 3} and {-5, 1}. And the indexes in the big array are {2, 3, 1, 3}. So the actual values in the big array after formatting it as concatenation of the small arrays will be {3, 3, -5, 1, 1, 6, -2, -5, 1}. In this example the maximum sum is 9.

Can you help Mostafa solve this problem?