#include<stdio.h>
#include<string.h>

int strStr(char*, char*);

int main(){
    char* s = "bbbaa";
    char* needle = "";
    printf("%d\n", strStr(s,needle));
    getchar();
    return 0;
}

int equal(char* s, char* needle, int len){
    for(int i = 0; i < len ; i++)
        if(s[i] != needle[i])
            return 0;

    return 1;
}

int strStr(char* haystack, char* needle) {
    int haylen = strlen(haystack);
    int len = strlen(needle);
    for(int i = 0; i < haylen - len + 1; i++){
        if(haystack[i] == *needle)
            if(equal(haystack + i, needle, len))
                return i;
    }
    return strlen(needle) == 0? 0 : -1;
}