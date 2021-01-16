
#define MIN(x, y) ((x) < (y) ? (x) : (y))
#define MAX(x, y) ((x) < (y) ? (y) : (x))

// need to malloc a new dp array, memory usage is huge.
int maximalSquare(char **matrix, int matrixSize, int *matrixColSize) {
  int maxSide = 0;
  if (matrix == NULL || matrixSize == 0 || *matrixColSize == 0)
    return maxSide;
  int rows = matrixSize;
  int columns = *matrixColSize;
  int **dp = (int **)malloc(sizeof(int *) * rows);
  for (int i = 0; i < rows; i++)
    dp[i] = (int *)calloc(columns, sizeof(int));
  for (int i = 0; i < rows; i++) {
    for (int j = 0; j < columns; j++) {
      if (matrix[i][j] == '1') {
        if (!i || !j)
          dp[i][j] = 1;
        else {
          //这句不写就错？
          int min = MIN(MIN(dp[i][j - 1], dp[i - 1][j]), dp[i - 1][j - 1]);
          dp[i][j] = 1 + min;
        }
        maxSide = MAX(maxSide, dp[i][j]);
      }
    }
  }
  return maxSide * maxSide;
}

// do modification directly to matrix
int maximalSquare(char **matrix, int matrixSize, int *matrixColSize) {
  int maxSide = 0;
  for (int i = 0; i < matrixSize; i++) {
    for (int j = 0; j < *matrixColSize; j++) {
      if (matrix[i][j] && i && j)
        matrix[i][j] =
            MIN(matrix[i - 1][j - 1], MIN(matrix[i - 1][j], matrix[i][j - 1])) +
            1;
      else
        matrix[i][j] = matrix[i][j] - '0';
      maxSide = MAX(maxSide, matrix[i][j]);
    }
  }
  return maxSide * maxSide;
}
