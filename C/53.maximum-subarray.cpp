/*
 * @lc app=leetcode id=53 lang=cpp
 *
 * [53] Maximum Subarray
 */

#include <vector>
using std::vector;

// @lc code=start
class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        int max = nums[nums.size() - 1];
        int cur = nums[nums.size() - 1];

        for (int i = nums.size() - 2; i >= 0; --i) {
            cur = nums[i] + (cur > 0 ? cur : 0);
            max = cur > max ? cur : max;
        }

        return max;
    }
};
// @lc code=end

