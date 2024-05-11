/*
 * @lc app=leetcode id=45 lang=cpp
 *
 * [45] Jump Game II
 */

#include <codecvt>
#include <iostream>
#include <vector>
using std::vector;

// @lc code=start
int minimum(vector<int> &nums, int start, int len) {

    int res = 10000;
    for (auto i = start; i < start + len; ++i) {
        if (nums[i] < res) res = nums[i];
    }

    return res;
}

class Solution {
  public:
    int jump(vector<int> &nums) {
        int         distance = 0;
        int         lastidx  = nums.size() - 1;
        int         min      = 10000;
        vector<int> dp       = vector<int>(nums.size(), 0);
        for (int i = nums.size() - 2; i >= 0; --i) {
            distance = lastidx - i;
            if (nums[i] == 0) {
                dp[i] = nums.size();
            }
            else if (nums[i] >= distance) {
                dp[i] = 1;
            } else {
                dp[i] = 1 + minimum(dp, i + 1, nums[i]);
            }
        }

        return dp[0] >= nums.size()? 0 : dp[0];
    }
};
// @lc code=end

class PSolution {
public:

    int jump(vector<int>& nums) {
    int n = nums.size();
    if (n == 1) return 0;

    int next_reach = 0 , cnt = 0 , max_reach = 0;
    for(int i=0;i<n-1;++i){
        next_reach = std::max(next_reach,i+nums[i]);
        
        if (i == max_reach) {
            max_reach = next_reach;
            cnt++;
            
            if (max_reach >= n-1) {
                break;
            }
        }
    }

    
    return cnt;
}
};
