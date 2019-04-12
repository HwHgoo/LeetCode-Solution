#include<stdio.h>

int search(int*, int, int);

int main(){
    int nums[] = {3,1};
    printf("%d",search(nums,2,3));
    getchar();
    return 0;
}


int binarySearch(int* nums, int high, int low, int target){
    int cur = (low + high)/2;
    if(nums[cur] == target)
        return cur;
    if(high == low)
        return -1;
    if(high - low == 1)
        return nums[high] == target? high : nums[low] == target?low : -1;
    return nums[cur] < target?binarySearch(nums, high, cur, target) : binarySearch(nums, cur, low, target);
}

int search(int* nums, int numsSize, int target){
    if(!numsSize)
        return -1;
    if(nums[0] < nums[numsSize -1])
        return binarySearch(nums, numsSize -1, 0, target);
    int pivot = 0;
    if(numsSize == 1)
        return nums[0] == target?0 : -1;
    for(int i = 1; i < numsSize; i++){
        if(nums[i] < nums[i -1]){
            pivot = i;
            break;
        }
    }
    return target > nums[0]?binarySearch(nums,pivot-1,0,target) : binarySearch(nums,numsSize-1,pivot,target);
}