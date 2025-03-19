Tyndex is again well ahead of the rivals! The reaction to the release of Zoozle Chrome browser was the release of a new browser Tyndex.Brome!

The popularity of the new browser is growing daily. And the secret is not even the Tyndex.Bar installed (the Tyndex.Bar automatically fills the glass with the finest 1664 cognac after you buy Tyndex.Bottles and insert in into a USB port). It is highly popular due to the well-thought interaction with the user.

Let us take, for example, the system of automatic address correction. Have you entered codehorses instead of codeforces? The gloomy Zoozle Chrome will sadly say that the address does not exist. Tyndex.Brome at the same time will automatically find the closest address and sent you there. That's brilliant!

How does this splendid function work? That's simple! For each potential address a function of the F error is calculated by the following rules:

for every letter ci from the potential address c the closest position j of the letter ci in the address (s) entered by the user is found. The absolute difference |i - j| of these positions is added to F. So for every i (1 ≤ i ≤ |c|) the position j is chosen such, that ci = sj, and |i - j| is minimal possible.
if no such letter ci exists in the address entered by the user, then the length of the potential address |c| is added to F.
After the values of the error function have been calculated for all the potential addresses the most suitable one is found.

To understand the special features of the above described method better, it is recommended to realize the algorithm of calculating the F function for an address given by the user and some set of potential addresses. Good luck!