#include<stdio.h>
#include<stdlib.h>

int lengthOfLongestSubstring(char*);
int lengthOfSubString(char*);


int main(){
    char* s = "abcabcabc";
    printf("%d", lengthOfSubString(s));
    getchar();
    return 0;
}

int inString(char* string, char ch, int end){
    for(int i = 0; i < end ; i++)
        if(ch == string[i])
            return 1;
    return 0;
}

//O(N^2)时间复杂度
int lengthOfLongestSubstring(char* s) {
    if (strlen(s) <= 1) 
        return strlen(s);
    char* ptr = s + 1;
    int count = 1;
    while(count < strlen(s) && !inString(s, *(ptr++), count))
        count++;
    if(count == strlen(s))
        return count;
    int len = lengthOfLongestSubstring(s+1);
    return count > len? count : len;
}

//O(N)时间复杂度
int lengthOfSubString(char *s){
    if(strlen(s) <= 0)
        return strlen(s);
    
    int a[95];
    for(int i = 0; i < 95; i++)
        a[i] = 0;

    int result = 1;
    a[*s - 33] = 1;
    char* ptr = s + 1;
    while (result < strlen(s) && !a[*ptr - 33]){
        result ++;
        a[(*ptr++) - 33] = 1;
    }
    int len = lengthOfSubString(s + 1);

    return result > len? result : len;
}