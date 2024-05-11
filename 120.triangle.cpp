/*
 * @lc app=leetcode id=120 lang=cpp
 *
 * [120] Triangle
 */

#include <vector>
using std::vector;


// @lc code=start
class Solution {
  public:
    int minimumTotal(vector<vector<int>> &triangle) {
        auto cache = triangle.back();
        for (auto row = triangle.rbegin() + 1; row != triangle.rend(); ++row) {
            for (int i = 0; i < row->size(); ++i) {
                cache[i] = (*row)[i] + (cache[i] < cache[i+1]? cache[i] : cache[i+1]);
            }
        }

        return cache[0];
    }
};
// @lc code=end
