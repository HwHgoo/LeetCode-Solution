#include <memory.h>
#include <stdio.h>
#include <stdlib.h>

int maxDepth(struct TreeNode *root, int level, int *colSizes) {
  if (root == NULL)
    return 0;
  colSizes[level] += 1;
  int leftDepth = maxDepth(root->left, level + 1, colSizes);
  int rightDepth = maxDepth(root->right, level + 1, colSizes);
  return 1 + (leftDepth > rightDepth ? leftDepth : rightDepth);
}

void traverse(struct TreeNode *root, int level, int *colSizes, int *cols,
              int **result) {
  // if root is null, return
  if (!root)
    return;
  // result array of this level not malloced
  if (!result[level])
    result[level] = (int *)malloc(sizeof(int) * cols[level]);
  // write node value to result array
  result[level][*colSizes++] = root->val;
  // traver left and right
  traverse(root->left, level + 1, colSizes, cols, result);
  traverse(root->right, level + 1, colSizes, cols, result);
}

int **levelOrder(struct TreeNode *root, int *returnSize,
                 int **returnColumnSizes) {

  int cols[15] = {0};
  // collect tree info. depth equals to returnSize and size of each column will
  // be written to array 'cols'
  *returnSize = maxDepth(root, 0, cols);
  // the result to be returned
  int **result = (int **)calloc(*returnSize, sizeof(int *));
  // initiate result
  for (int i = 0; i < *returnSize; i++) {
    result[i] = NULL;
  }
  // column sizes to be returned
  int *colSizes = (int *)calloc(*returnSize, sizeof(int));
  returnColumnSizes = &colSizes;
  // traverse tree to write node values to result
  // params: treeroot, current level, column sizes to be returned, collected
  // info, result to be returned
  traverse(root, 0, colSizes, cols, result);
  return result;
}

int main() { return 0; }
