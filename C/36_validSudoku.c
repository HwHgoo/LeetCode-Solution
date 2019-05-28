
// 2019/5/28 8ms
bool isValidSudoku(char** board, int size, int *useless){
    int rowMonitor[9][9] = {0};
    int colMonitor[9][9] = {0};
    int blockMonitor[9][9] = {0};

    for(int row = 0; row < 9; row++){
        for(int col = 0; col < 9; col++){
            if(board[row][col]== '.')
                continue;
            //检验行
            if(rowMonitor[row][board[row][col] - 49])
                return false;
            rowMonitor[row][board[row][col] - 49] = 1;

            //检验列
            if(colMonitor[col][board[row][col] - 48])
                return false;
            colMonitor[col][board[row][col] - 48] = 1;

            //检验所在的九宫格
            if(blockMonitor[(row/3)*3 + col/3][board[row][col] - 48])
                return false;
            blockMonitor[(row/3)*3 + col/3][board[row][col] - 48] = 1;
        }
    }

    return true;
}