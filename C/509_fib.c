#include <stdio.h>
#include <stdlib.h>

int helper(int* flag, int n){
    if(n == 1 || n == 2) return 1;
    if(flag[n]) return flag[n];
    flag[n] = helper(flag, n - 1) + helper(flag, n -2);
    return flag[n];
}

int fib(int n){
    if (n < 1) return 0;
    int * flag = (int*)calloc(30, sizeof(int));
    return helper(flag, n);
}


int main(){
	int n = 10;
	printf("fib(%d) = %d\n", n, fib(n));
	return 0;
}
