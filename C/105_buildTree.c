struct TreeNode *buildTree(int *preorder, int preorderSize, int *inorder,
                           int inorderSize) {
  if (!preorderSize)
    return NULL;
  struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
  root->val = preorder[0];
  int rootIndex = 0;
  while (inorder[rootIndex] != *preorder)
    ++rootIndex;
  root->left = buildTree(preorder + 1, rootIndex, inorder, rootIndex);
  root->right =
      buildTree(preorder + rootIndex + 1, preorderSize - rootIndex - 1,
                inorder + rootIndex + 1, preorderSize - rootIndex - 1);
  return root;
}
