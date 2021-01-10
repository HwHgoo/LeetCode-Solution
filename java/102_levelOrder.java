
class Solution {
    public List<List<Integer>> levelOrder(TreeNode root) {
			List<List<Integer>> result = new ArrayList<ArrayList<>>();
			traverse(root, result, 0);
			return result;
    }

		private void traverse(TreeNode root, List result, int level){
			if(root == null) return;
			if(result.get(level)== null){
				result.add(level, new ArrayList<Integer>();)
			}
			result.get(level).add(root.val)	;
			traverse(root.left, result, level+1);
			traverse(root.right, result, level+1);
		}
}
