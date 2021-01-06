int find(int left, int right, int * map){
  if(left > right) return 1;
  if(map[right-left]) return map[right-left];
  int res = 0;
  for(int i = left; i <= right; i++){
    int leftSum = find(left, i-1, map);
    int rightSum = find(i+1, right, map);
    res += leftSum*rightSum;
  }
  map[right - left] = res;
  return res;
  
}

int numTrees(int n){
  int map[19] = {0};
  
  return find(1, n, map);
}
