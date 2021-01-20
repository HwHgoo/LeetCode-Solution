int countVowelStrings(int n) {
  int **dp = (int **)calloc(n, sizeof(int *));
  for (int i = 0; i < n; i++)
    dp[i] = (int *)calloc(4, sizeof(int));
  for (int i = 0; i < 4; i++)
    dp[0][i] = 1;
  for (int i = 1; i < n; i++) {
    dp[i][0] = dp[i - 1][0] + dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + 1;
    dp[i][1] = dp[i - 1][1] + dp[i - 1][2] + dp[i - 1][3] + 1;
    dp[i][2] = dp[i - 1][2] + dp[i - 1][3] + 1;
    dp[i][3] = dp[i - 1][3] + 1;
  }
  return dp[n - 1][0] + dp[n - 1][1] + dp[n - 1][2] + dp[n - 1][3] + 1;
}
