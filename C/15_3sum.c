#include<stdio.h>
#include<stdlib.h>

int partition(int*, int, int);
void quicksort(int*, int, int);
int** threeSum(int*, int, int*);


int main(){
    int a[7] = {-5,4,1,3,2,-6,11};    
    int* returnSize = (int*)malloc(sizeof(int));
    threeSum(a ,7 ,returnSize);
    getchar();
    return 0;
}


//找到target所在的index
int find(int* nums, int start, int end, int target){
    for(int i = start; i < end; i++){
        if (nums[i] == target) {
            return i;
        }
    }
    return -1;
}

//判断两个数组是否相同，仅适用于本题
int equal(int* nums1, int* nums2){
    return nums1[0] == nums2[0] && nums1[1] == nums2[1] && nums1[2] == nums2[2];
}

/**
 * 添加新的结果到结果数组中
 * arraies 结果数组
 * nums 要添加的结果
 * arraySize 添加前结果数组的长度
 */
int** add(int** result, int* nums, int* returnSize){
    for(int i = 0; i < *returnSize; i++){
        if (equal(result[i], nums)) {
            printf("本次没有添加\n");
            return result;
        }
    }
    result = (int**)realloc(result,sizeof(int*)*(*returnSize + 1));
    result[*returnSize] = nums;
    *returnSize += 1;
    printf("添加了一次\n");
    return result;
}

int** threeSum(int* nums, int numsSize, int* returnSize){
    int** result = (int**)malloc(sizeof(int*));
    *returnSize = 0;
    quicksort(nums, 0, numsSize -1); //排序
    for(int low = 0; low < numsSize -2; low ++){
        for(int high = numsSize - 1; high > low + 1; high--){

            int target = 0-nums[low]-nums[high];
            int mid = find(nums, low + 1, high, target);
            if (mid != -1){
                //int res[3] = {nums[low], nums[mid], nums[high]}
                //是错误的是因为在这次循环中，并不会每次使用都重新创建导致第n次答案会覆盖掉上一次的答案
                //所以需要每次都重新申请内存以存储不同的答案
                int *res = (int*)malloc(sizeof(int)*3);
                res[0] = nums[low];
                res[1] = nums[mid];
                res[2] = nums[high];
                result = add(result, res, returnSize);
            }
        }
    }
    printf("%d\n", *returnSize);
    return result;
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