void swap(int *nums, int left, int right) {
  int tmp = nums[left];
  nums[left] = nums[right];
  nums[right] = tmp;
}

void reverse(int *nums, int start, int end) {
  while (start < end) {
    swap(nums, start, end);
    start++;
    end--;
  }
}

void rotate(int *nums, int numsSize, int k) {
  k %= numsSize;
  reverse(nums, 0, numsSize - k - 1);
  reverse(nums, numsSize - k, numsSize - 1);
  reverse(nums, 0, numsSize - 1);
}
