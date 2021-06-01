#define Min(x, y) (x) < (y)? (x) : (y)

int coinChange(int* coins, int coinsSize, int amount){
    int * dp = (int*)malloc((amount + 1) * sizeof(int));
    dp[0] = 0;
    for(int i = 1; i <= amount; ++i){
        dp[i] = amount + 1;
        for(int j = 0; j < coinsSize; ++j){
            if(coins[j] > i) continue;;
            dp[i] = Min(dp[i], 1 + dp[i - coins[j]]);
        }
    }
    return dp[amount] == amount + 1? -1 : dp[amount];
}

