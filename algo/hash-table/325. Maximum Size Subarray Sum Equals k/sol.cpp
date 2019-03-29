/*

325. Maximum Size Subarray Sum Equals k

Medium

Given an array nums and a target value k, find the maximum length of a subarray that sums to k. If there isn't one, return 0 instead.

Note:
The sum of the entire nums array is guaranteed to fit within the 32-bit signed integer range.

Example 1:

Input: nums = [1, -1, 5, -2, 3], k = 3
Output: 4 
Explanation: The subarray [1, -1, 5, -2] sums to 3 and is the longest.
Example 2:

Input: nums = [-2, -1, 2, 1], k = 1
Output: 2 
Explanation: The subarray [-1, 2] sums to 1 and is the longest.
Follow Up:
Can you do it in O(n) time?

*/

class Solution {
public:
    int maxSubArrayLen(vector<int>& nums, int k) {
        unordered_map<int, int> hs;
        hs[0] = -1;
        int sum = 0, res = 0, prev_sum = 0;
        
        for (int i = 0; i < nums.size(); ++i) {
            sum += nums[i];
            if (hs.find(sum) == hs.end())
                hs[sum] = i;
            prev_sum = sum - k;
            if (hs.find(prev_sum) != hs.end())
                res = max(res, i - hs[prev_sum]);
        }
        return res;
    }
};
