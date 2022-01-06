func orangesRotting(grid [][]int) int {
    time,flag:=2,true
    for flag{
        flag=false
        for i:=0;i<len(grid);i++{
            for j:=0;j<len(grid[0]);j++{
                if grid[i][j]<time{
                    continue
                }else if grid[i][j]==time{
                    if i-1>=0&&grid[i-1][j]==1{
                        grid[i-1][j]=time+1
                        flag=true
                    }
                    if j-1>=0&&grid[i][j-1]==1{
                        grid[i][j-1]=time+1
                        flag=true
                    }
                    if i+1<len(grid)&&grid[i+1][j]==1{
                        grid[i+1][j]=time+1
                        flag=true
                    }
                    if j+1<len(grid[0])&&grid[i][j+1]==1{
                        grid[i][j+1]=time+1
                        flag=true
                    }
                }else{
                    continue
                }
            }
        }
        time++
    }
    for i:=0;i<len(grid);i++{
        for j:=0;j<len(grid[0]);j++{
            if grid[i][j]==1{
                return -1
            }
        }
    }
    return time-3
}