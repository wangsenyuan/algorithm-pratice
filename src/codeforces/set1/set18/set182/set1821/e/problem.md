A regular bracket sequence is a bracket sequence that can be transformed into a correct arithmetic expression by
inserting characters "1" and "+" between the original characters of the sequence. For example:

bracket sequences "()()" and "(())" are regular (the resulting expressions are: "(1)+(1)" and "((1+1)+1)");
bracket sequences ")(", "(" and ")" are not.
You are given a regular bracket sequence. In one move, you can remove a pair of adjacent brackets such that the left one
is an opening bracket and the right one is a closing bracket. Then concatenate the resulting parts without changing the
order. The cost of this move is the number of brackets to the right of the right bracket of this pair.

The cost of the regular bracket sequence is the smallest total cost of the moves required to make the sequence empty.

Actually, you are not removing any brackets. Instead, you are given a regular bracket sequence and an integer 𝑘
. You can perform the following operation at most 𝑘
times:

extract some bracket from the sequence and insert it back at any position (between any two brackets, at the start or at
the end; possibly, at the same place it was before).
After all operations are performed, the bracket sequence has to be regular. What is the smallest possible cost of the
resulting regular bracket sequence?