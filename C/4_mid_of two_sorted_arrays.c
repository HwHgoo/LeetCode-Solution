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

double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    if(!nums1Size || !nums2Size)
        return !nums1Size?(nums2Size%2 == 0?(nums2[nums2Size/2] + nums2[nums2Size/2 -1])/2.0 : nums2[nums2Size/2]) : (nums1Size%2 == 0?(nums1[nums1Size/2] + nums1[nums1Size/2 -1])/2.0 : nums1[nums1Size/2]);
    
    int* newArray = (int*)malloc(sizeof(int) * (nums1Size + nums2Size));
    for(int i = 0;i < nums1Size; i++)
        newArray[i] = nums1[i];
    for(int i = 0; i < nums2Size; i++)
        newArray[i + nums1Size] = nums2[i];
        
    quicksort(newArray, 0, nums1Size + nums2Size -1);

    int size = nums2Size + nums1Size;

    if(size%2 == 0)
        return (newArray[size/2 - 1] + newArray[size / 2])/2.0;
    return newArray[size/2]/1.0;
}