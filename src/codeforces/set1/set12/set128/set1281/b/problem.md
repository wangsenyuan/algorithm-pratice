Your friend Jeff Zebos has been trying to run his new online company, but it's not going very well. He's not getting a lot of sales on his website which he decided to call Azamon. His big problem, you think, is that he's not ranking high enough on the search engines. If only he could rename his products to have better names than his competitors, then he'll be at the top of the search results and will be a millionaire.

After doing some research, you find out that search engines only sort their results lexicographically. If your friend could rename his products to lexicographically smaller strings than his competitor's, then he'll be at the top of the rankings!

To make your strategy less obvious to his competitors, you decide to swap no more than two letters of the product names.

Please help Jeff to find improved names for his products that are lexicographically smaller than his competitor's!

Given the string 𝑠
 representing Jeff's product name and the string 𝑐
 representing his competitor's product name, find a way to swap at most one pair of characters in 𝑠
 (that is, find two distinct indices 𝑖
 and 𝑗
 and swap 𝑠𝑖
 and 𝑠𝑗
) such that the resulting new name becomes strictly lexicographically smaller than 𝑐
, or determine that it is impossible.

Note: String 𝑎
 is strictly lexicographically smaller than string 𝑏
 if and only if one of the following holds:

𝑎
 is a proper prefix of 𝑏
, that is, 𝑎
 is a prefix of 𝑏
 such that 𝑎≠𝑏
;
There exists an integer 1≤𝑖≤min(|𝑎|,|𝑏|)
 such that 𝑎𝑖<𝑏𝑖
 and 𝑎𝑗=𝑏𝑗
 for 1≤𝑗<𝑖
.
