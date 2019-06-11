#include<stdio.h>
#include <stdlib.h>
#include <memory.h>
#include <limits.h>

typedef unsigned int uint32_t;

int firstMissingPositive(int*, int);

int main(){
    int array[] = {1,3,3};
    printf("%d\n",firstMissingPositive(array,3));
    getchar();
    return 0;
}

int firstMissingPositive(int* nums, int numsSize){
    if(numsSize == 1)
        return *nums == 1? 2 : 1;

    int min = INT_MAX;
    for(int i = 0; i < numsSize; i++)
        if(nums[i] > 0 && nums[i] < min)
            min = nums[i];

    if(min > 1)
        return 1;

    uint32_t* list = (uint32_t*)malloc((numsSize + 33L)/32 * sizeof(uint32_t));
    memset(list, 0, (numsSize + 33L)/32 * sizeof(int));
    for(int i = 0; i < numsSize; i++){
        if(nums[i] > 0){
        //if a - min >= numsSize then one number, at least, between min to a is not included in nums; 
        if(nums[i] - min >= numsSize)
            continue;
        int diff = nums[i] - min;
        int listIndex = diff / 32;
        int index = diff % 32;
        list[listIndex] |= 0x1u << index;
        }
    }

    for(int i = 0; i < sizeof(uint32_t) * (numsSize + 32L)/32; i++){
        if(list[i] == 0xFFFFFFFFu)
            continue;
        
        for(int j = 0; j < 32; j++){
            if(!(list[i] & (0x1u<<j)))
                return min + 32*i + j;
        }
    }
    
    return 0;
}