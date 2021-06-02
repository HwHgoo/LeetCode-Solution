#define Max(x, y) (x) < (y) ? (y) : (x)

int longestValidParentheses(char * s){
	int len = strlen(s);
	int* dp = (int*)calloc(len, sizeof(int));
	int cur = 1;
	int max = 0;
	int flag = 0;
	while(cur < len){

		if(s[cur] == '('){
			if(s[cur - 1] == '(')
				dp[cur] = 0;
			else
				dp[cur] = dp[cur - 1];
		}
		else{
			if(s[cur - 1] == '('){
				dp[cur] = dp[cur - 1] + 1;
				max = Max(max, dp[cur]);
			}
			else {
				flag = cur - 1 - dp[cur -1]*2;
				while(flag >= 0){
					if(s[flag] == '('){
						dp[cur] = dp[cur - 1] + 1 + dp[flag];
						max = Max(max, dp[cur]);
						break;
					}
					if(dp[flag] == 0){
						dp[cur] = 0;
						break;
					}
					flag = flag - dp[flag]*2;
				}
			}
		}

		cur++;
	}
	return max*2;
}
