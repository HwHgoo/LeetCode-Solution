************************************************************************Original Solution**************************************************************************************

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


*******************************************************************************************************************************************************************************
******************************************************************* Improved Solution *****************************************************************************************


int cmp(const void* arg1, const void* arg2){
    return *(int*)arg1 - *(int*)arg2;
}

int** addSolution(int a, int b, int c, int** result, int* returnSize){
    if(*returnSize % 100 == 0)
        result = (int**)realloc(result, sizeof(int*) * (*returnSize + 100));
    int* solution = (int*)malloc(sizeof(int)*3);
    solution[0] = a;
    solution[1] = b;
    solution[2] = c;
    result[(*returnSize)++] = solution;
    return result;
}


//这个解法参考于 https://leetcode.com/problems/3sum/discuss/7523/26ms-C-solution
int** threeSum(int* nums, int numsSize, int* returnSize){
    *returnSize = 0;
    int i , j, k;
    int nc = 0, pc = 0, zc = 0;//negativecount;positivecount;zerocount

    //申请空间，每次申请100个
    int** result = (int**)malloc(sizeof(int*) * 100);


    if(numsSize < 3)
        return result;

    //排序后求出nc pc 和 zc
    qsort(nums, numsSize, sizeof(int), cmp);
    for(i = 0; i < numsSize; i++){
        if(nums[i] < 0)
            nc++;
        else if(nums[i] == 0)
            zc++;
        else
            break;
    }
    pc = numsSize - nc - zc;

    //解中有3个0的
    if(zc >= 3)
        result = addSolution(0, 0, 0, result, returnSize);

    //解中有1个零
    //那么另外两个数就必须异号，从而只需要从nums的两端进行遍历即可
    if(zc >= 1){
        for(i = 0,j = numsSize - 1 ;i < nc && j > numsSize -1 - pc;){
            if(nums[i] + nums[j] < 0) i++;
            else if(nums[i] + nums[j] > 0) j--;
            else if(i > 0 && nums[i] == nums[i-1]) i++;             //跳过重复元素
            else if(j < numsSize -1 && nums[j] == nums[j+1]) j--;   //跳过重复元素
            else result = addSolution(nums[i++], 0, nums[j--],result,returnSize);
        }
    }

    //解中有两个负数
    if(nc >= 2 && pc >=1){
        for(i = 0; i < nc - 1; i++){
            if(nums[i] + nums[numsSize -1] < 0);
            else if(i > 0 && nums[i] == nums[i -1]);
            else
                for(k = i + 1, j = numsSize -1; k < nc && j > numsSize - 1 - pc;){
                    if(nums[i] + nums[j] + nums[k] < 0) k++; 
                    else if(nums[i] + nums[j] + nums[k] > 0) j--;
                    else if(k > i +1 && nums[k] == nums[k -1]) k++;     //跳过重复元素
                    else if(j < numsSize -1 && nums[j] == nums[j+1]) j--;//跳过重复元素
                    else result = addSolution(nums[i],nums[k++],nums[j--],result,returnSize);
                }
        }
    }

    //解中有两个正数
    if(pc >= 2 && nc >= 1){
        for(i = numsSize -1; i > numsSize -2 - pc; i--){
            if(nums[i] + nums[0] > 0);
            else if(i < numsSize -1 && nums[i] == nums[i+1]);
            else
                for(j = 0, k = i -1; j < nc && k > numsSize - 1 - pc;){
                    if(nums[i] + nums[j] + nums[k] < 0) j++;
                    else if(nums[i] + nums[j] + nums[k] > 0) k--;
                    else if(j > 0 && nums[j] == nums[j-1]) j++;
                    else if(k < i -1 && nums[k] == nums[k+1]) k--;
                    else result = addSolution(nums[i], nums[j++], nums[k--],result,returnSize);

                }
        }
    }

    return result;
}



