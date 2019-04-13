#include<stdio.h>
#include<stdlib.h>
#include<limits.h>

void nextPermutation(int*, int);
void getIndexes(int*, int, int*);
void quicksort(int*, int, int);

int main(){
    int a[5] = {1,1,1,1,1};
    nextPermutation(a,5);
    getchar();
    return 0;
}

//1.找到数字pivot，要求pivot后面至少有一个数字比它大
//2.将pivot与它后面的差绝对值最小的数交换
//3.将原pivot所在位置后面进行升序排序

//找到pivot和与之交换的数的下标,存放到result中
//第一个位置存放pivot下标
void getIndexes(int* nums, int numsSize, int* result){
    result[0] = result[1] = -1;
    int pivotIndex = -1;
    int max = INT_MIN;
    for(int i = numsSize -1; i >= 0; i--){
        if(nums[i] >= max)
            max = nums[i];
        else{
            pivotIndex = i;
            break;
        }
    }
    result[0] = pivotIndex;

    if(pivotIndex == -1)
        return;
    
    int index = -1;
    int min = INT_MAX;
    for(int i = numsSize -1; i > pivotIndex; i--){
        if(nums[i] > nums[pivotIndex] && nums[i] - nums[pivotIndex] < min){
            index = i;
            min = nums[i] - nums[pivotIndex];
        }
    }
    result[1] = index;
}

void nextPermutation(int* nums, int numsSize) {
    if(numsSize <= 1)
        return;
    
    int* result = (int*)malloc(sizeof(int)*2);
    getIndexes(nums,numsSize,result);
    if(result[0] == -1){
        quicksort(nums, 0, numsSize - 1);
        return;
    }
    
    //交换
    int tmp = nums[result[0]];
    nums[result[0]] = nums[result[1]];
    nums[result[1]] = tmp;

    //排序
    quicksort(nums, result[0]+1, numsSize -1);
}


int partition(int* nums, int low, int high){
    int pivotkey = nums[low];
    while(low < high){
        while(low < high && nums[high] >= pivotkey) 
            --high;
        nums[low] = nums[high];
        while(low < high && nums[low] <= pivotkey) 
            ++low;
        nums[high] = nums[low]; 
    }
    nums[low] = pivotkey;
    return low;
}

void quicksort(int* nums, int low, int high){
    if (low < high) {
        int pivotkey = partition(nums, low, high);
        quicksort(nums, low, pivotkey -1);
        quicksort(nums, pivotkey +1, high);
    }
}