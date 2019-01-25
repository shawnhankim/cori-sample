/*
135. Candy

There are N children standing in a line. Each child is assigned a rating value.

You are giving candies to these children subjected to the following requirements:

Each child must have at least one candy.
Children with a higher rating get more candies than their neighbors.
What is the minimum candies you must give?

Example 1:

Input: [1,0,2]
Output: 5
Explanation: You can allocate to the first, second and third child with 2, 1, 2 candies respectively.
Example 2:

Input: [1,2,2]
Output: 4
Explanation: You can allocate to the first, second and third child with 1, 2, 1 candies respectively.
             The third child gets 1 candy because it satisfies the above two conditions.


*/

class Solution {
public:
    int candy(vector<int>& ratings) {
        int N = (int)ratings.size();
        if (N <= 1) return N;
        
        int sum = 0, i = 0;
        bool isChanged = true;
        vector<int> vals(N, 1);
        
        while(isChanged) {
            isChanged = false;
            for (i = 0; i < N; ++i) {
                if (i != N-1 && ratings[i] > ratings[i+1] && vals[i] <= vals[i+1]) {
                    vals[i] = vals[i+1] + 1; isChanged = true;
                }
                if (i != 0 && ratings[i] > ratings[i-1] && vals[i] <= vals[i-1]) {
                    vals[i] = vals[i-1] + 1; isChanged = true;
                }
            }
        }
        
        for (int v : vals)
            sum += v;
        return sum;
    }
};
