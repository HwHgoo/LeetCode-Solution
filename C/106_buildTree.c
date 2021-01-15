struct TreeNode *buildTree(int *inorder, int inorderSize, int *postorder,
                           int postorderSize) {
  if (inorderSize <= 0)
    return NULL;
  struct TreeNode *root = (struct TreeNode *)malloc(sizeof(struct TreeNode));
  root->val = postorder[postorderSize - 1];
  int inorderRootIndex = -1;
  for (int i = 0; i < inorderSize; i++) {
    if (inorder[i] == postorder[postorderSize - 1]) {
      inorderRootIndex = i;
      break;
    }
  }
  root->left =
      buildTree(inorder, inorderRootIndex, postorder, inorderRootIndex);
  root->right = buildTree(
      inorder + inorderRootIndex + 1, inorderSize - inorderRootIndex - 1,
      postorder + inorderRootIndex, postorderSize - inorderRootIndex - 1);
  return root;
}
