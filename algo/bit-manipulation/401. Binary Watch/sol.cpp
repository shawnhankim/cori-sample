/*

401. Binary Watch

Easy

A binary watch has 4 LEDs on the top which represent the hours (0-11), and the 6 LEDs on the bottom represent the minutes (0-59).

Each LED represents a zero or one, with the least significant bit on the right.

 +--------------------+
 |  o    o    *    o  |
 |  8    4    2    1  |
 |                    |
 |  o  *  *  o  o  *  |
 |  32 16 8  4  2  1  | 
 +--------------------+

For example, the above binary watch reads "3:25".

Given a non-negative integer n which represents the number of LEDs that are currently on, return all possible times the watch could represent.

Example:

Input: n = 1
Return: ["1:00", "2:00", "4:00", "8:00", "0:01", "0:02", "0:04", "0:08", "0:16", "0:32"]
Note:
The order of output does not matter.
The hour must not contain a leading zero, for example "01:00" is not valid, it should be "1:00".
The minute must be consist of two digits and may contain a leading zero, for example "10:2" is not valid, it should be "10:02".

*/

class Solution {
public:
    vector<string> readBinaryWatch(int num) {
        vector<string> res;
        vector<int>    hh{8, 4, 2, 1}, mm{32, 16, 8, 4, 2, 1};
        for (int i = 0; i <= num; ++i) {
            vector<int> hhs = generate(hh, i);
            vector<int> mms = generate(mm, num-i);
            for (int h : hhs) {
                if (h > 11) continue;
                for (int m : mms) {
                    if (m > 59) continue;
                    res.push_back(to_string(h) +
                                  (m < 10 ? ":0" : ":") + to_string(m));
                }
            }
        }
        return res;
    }
    
    vector<int> generate(vector<int>& nums, int cnt) {
        vector<int> res;
        helper(nums, cnt, 0, 0, res);
        return res;
    }
    
    void helper(vector<int>& nums, int cnt, int pos, int out, vector<int>& res) {
        if (cnt == 0) {
            res.push_back(out);
            return;
        }
        for (int i = pos; i < nums.size(); ++i) {
            helper(nums, cnt-1, i+1, out+nums[i], res);
        }
    }
};
