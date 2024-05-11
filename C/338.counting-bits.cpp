/*
 * @lc app=leetcode id=338 lang=cpp
 *
 * [338] Counting Bits
 */

// @lc code=start
#include <vector>
using std::vector;

class Solution {
public:
    vector<int> countBits(int n) {
        vector<int> ans(n+1, 0);
        for (int i = 0; i <= n; i++) {
            ans[i] = ans[i >> 1] + (i&1);
        }
        return ans;
    }
};
// @lc code=end

