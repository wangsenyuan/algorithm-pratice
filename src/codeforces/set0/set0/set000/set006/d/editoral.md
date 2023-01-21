Let's assume that the archers are numbered starting from 0. I will use the words archer and soldier interchangeably.

Let me first introduce a simplification. The problem says to make the health of the soldier < 0. Let us assume that we have to make the health of the soldier <= 0. To achieve this transition and obtain the same no. of min. moves as the case with < 0, we increment each h[i] by 1.

The only way to kill the 0th and the (n-1)th soldier is to shoot directly at the 1st and the (n-2)th soldiers respectively. Note that 1st and the (n-2)th soldier may be the same. So, at first we kill the 0th and the (n-1)th soldier via this process.

Now, we are left with soldiers from 1 .. (n-2). We'll use DP to solve the problem. Let's process the soldiers in sequence. When we are at the 1st soldier, we may either kill it directly by shooting appropriate no. of balls/bullets on it, or we may leave some of its health which will ultimately be made 0 by the direct shots that we make at the soldier 2.

When we are at the 2nd soldier, the above said things for the 1st soldier hold true for it as well. However, note here that the health of the 2nd soldier must have been reduced indirectly already due to the direct attacks that we made to the 1st soldier, and again we may or may not completely kill it and the remaining damage will be incurred to it by the next soldier (3rd soldier).

These observations tell us while we are at a particular soldier, our state is dependent on:

The index of the soldier that we are at.

The health of the previous soldier remaining. (This will ensure that we shoot atleast as many bullets on the current soldier so that the previous soldier gets killed / his health becomes <= 0).

The no. of direct attacks made on the previous soldier. (this can reduce the health of the current soldier).

Another observation: Since we only have to make the health <= 0, we can take any health < 0 to be 0. This will help us in avoiding negative indexing for the 2nd state variable mentioned above.

So, that's it. We can use a 3 state DP: dp[i][j][k]: where dp[i][j][k]= Min. no. of shots to kill the archers starting from the ith archer when the previous archer still has j health points left and the no. of direct shots at the arches soldiers were k.

