/*
 * @lc app=leetcode id=91 lang=cpp
 *
 * [91] Decode Ways
 */

#include <string>
using std::string;

// @lc code=start
class Solution {
  public:
    int numDecodings(string s) {
        if (s[0] == '0') return 0;
        int  fn2 = 1;    // f(n - 2)
        int  fn1 = 1;    // f(n - 1)
        char pre = s[0]; // a(n -1)
        int tmp;

        for (auto c = s.begin() + 1; c != s.end(); ++c) {
            if (*c == '0') {
                if (pre >= '3' || pre == '0') return 0;
                fn1 = fn2;
            } else if (pre == '1' || (pre == '2' && *c <= '6')) {
                fn1 = fn1 + fn2;
            }
            fn2 = tmp;
            pre = *c;
        }

        return fn1;
    }
};
// @lc code=end
