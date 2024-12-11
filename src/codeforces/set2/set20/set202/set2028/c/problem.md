Alice is at the Mad Hatter's tea party! There is a long sheet cake made up of 𝑛
 sections with tastiness values 𝑎1,𝑎2,…,𝑎𝑛
. There are 𝑚
 creatures at the tea party, excluding Alice.

Alice will cut the cake into 𝑚+1
 pieces. Formally, she will partition the cake into 𝑚+1
 subarrays, where each subarray consists of some number of adjacent sections. The tastiness of a piece is the sum of tastiness of its sections. Afterwards, she will divvy these 𝑚+1
 pieces up among the 𝑚
 creatures and herself (her piece can be empty). However, each of the 𝑚
 creatures will only be happy when the tastiness of its piece is 𝑣
 or more.

Alice wants to make sure every creature is happy. Limited by this condition, she also wants to maximize the tastiness of her own piece. Can you help Alice find the maximum tastiness her piece can have? If there is no way to make sure every creature is happy, output −1
.