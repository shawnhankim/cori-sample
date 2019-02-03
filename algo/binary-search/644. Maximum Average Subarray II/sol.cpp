/*

644. Maximum Average Subarray II

Given an array consisting of n integers, find the contiguous subarray whose length is greater than or equal to k that has the maximum average value. And you need to output the maximum average value.

Example 1:

Input: [1,12,-5,-6,50,3], k = 4
Output: 12.75
Explanation:
when length is 5, maximum average value is 10.8,
when length is 6, maximum average value is 9.16667.
Thus return 12.75.
Note:

1 <= k <= n <= 10,000.
Elements of the given array will be in range [-10,000, 10,000].
The answer with the calculation error less than 10-5 will be accepted.

*/

// 21/74 passed
class Solution {
public:
    double findMaxAverage(vector<int>& nums, int k) {
        // Deciare variable
        double res = 0.0f, tmp = 0.0f;
        double val = 0.0f;
        int n = nums.size();
        
        // Check boundary
        if (n == 0) return res;
        if (n == 1) return (double)nums[0];

        // Initialize DP memory
        vector<int> dp(n, 0);
        dp[0] = nums[0];
        for (int i = 1; i < n; ++i) {
            dp[i] = dp[i-1] + nums[i];
        }
        
        // Algorithm
        for (int i = k - 1; i < n; ++i) {
            int len = i + 1;
            int sum = 0;
            int o = len - i - 1;
            
            for (int j = i; j < n; ++j) {
                if (j == i) tmp = (double) dp[j] / len;
                else        tmp = (double)(dp[j] - sum) /len;
                sum += dp[o];
                ++o;
                res = max(res, tmp);
            }
        }
        return res;
    }
};

