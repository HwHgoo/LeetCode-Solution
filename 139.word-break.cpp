/*
 * @lc app=leetcode id=139 lang=cpp
 *
 * [139] Word Break
 */

#include <algorithm>
#include <iterator>
#include <random>
#include <set>
#include <string>
#include <string_view>
#include <vector>
using std::string;
using std::string_view;
using std::vector;

// @lc code=start
class Solution {
  public:
    bool solve(const string_view s, vector<string> &dict) {
        for (const string &word : dict) {
            if (s == word) return true;
        }

        for (const string &word : dict) {
            if (s.starts_with(word)) {
                if (solve(s.substr(word.length()), dict)) { return true; }
            }
        }

        return false;
    }

    bool wordBreak(string s, vector<string> &wordDict) {
        std::set<char> ss(s.begin(), s.end());
        std::set<char> sw;
        for (const string& w : wordDict) {
            sw.insert(w.begin(), w.end());
        }
        if (!std::includes(sw.begin(), sw.end(), ss.begin(), ss.end())) return false;

        std::sort(wordDict.begin(), wordDict.end(), [](const string &str1, const string &str2) { return str1.length() < str2.length(); });
        string_view sv(s);
        return solve(sv, wordDict);
    }
};
// @lc code=end
