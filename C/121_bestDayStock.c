int maxProfit(int *prices, int pricesSize) {
  int profit = 0;
  int minBuyPrice = 100000;
  for (int i = 0; i < pricesSize; i++) {
    if (prices[i] < minBuyPrice) {
      minBuyPrice = prices[i];
      for (int j = i + 1; j < pricesSize; j++) {
        int cur = prices[j] - prices[i];
        profit = cur > profit ? cur : profit;
      }
    }
  }
  return profit;
}

//优化后的
int maxProfit(int *prices, int pricesSize) {
  int profit = 0;
  int lowPrice = 100000;
  for (int i = 0; i < pricesSize; i++) {
    if (prices[i] < lowPrice)
      lowPrice = prices[i];
    else
      profit = profit < prices[i] - lowPrice ? prices[i] - lowPrice : profit;
  }
  return profit;
}

#define MAX(x, y) (x) < (y) ? (y) : (x)
//使用动态规划
int maxProfit(int *prices, int pricesSize) {
  if (pricesSize < 2)
    return 0;
  int dp[2] = {0};
  // dp[0] 存储当天最大的利润，dp[1]存储最小成本
  dp[1] = -prices[0];
  for (int i = 1; i < pricesSize; i++) {
    //更新最大利润和最小成本
    dp[0] = MAX(dp[0], dp[1] + prices[i]);
    dp[1] = MAX(dp[1] - prices[i]);
  }
  return dp[0];
}
