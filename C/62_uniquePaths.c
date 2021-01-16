int util(int m, int n, int **memory) {
  if (m == 1 || n == 1)
    return 1;
  if (memory[m - 1][n - 1])
    return memory[m - 1][n - 1];
  memory[m - 1][n - 1] = util(m - 1, n, memory) + util(m, n - 1, memory);
  return memory[m - 1][n - 1];
}
int uniquePaths(int m, int n) {
  int **memory = (int **)malloc(sizeof(int *) * m);
  for (int i = 0; i < m; i++)
    memory[i] = (int *)calloc(n, sizeof(int));
  int res = util(m, n, memory);
  free(memory);
  return res;
}
