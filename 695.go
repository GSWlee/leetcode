func maxAreaOfIsland(grid [][]int) int {
    width,high:=len(grid[0]),len(grid)
	flag:=[][]int{}
	for i:=0;i<high;i++{
		flag=append(flag,make([]int,width))
	}
    ans:=0
    for i:=0;i<high;i++{
        for j:=0;j<width;j++{
            if grid[i][j]==0 ||flag[i][j]==1{
                continue
            }else{
                stack:=[][]int{{i,j}}
                temp:=0
                for len(stack)>0{
                    x,y:=stack[0][0],stack[0][1]
                    if grid[x][y]==0{
                        stack=stack[1:]
                    }else{
                        temp++
                        flag[x][y]=1
                        if x-1>=0 &&grid[x-1][y]!=0&&flag[x-1][y]==0{
                            stack=append(stack,[]int{x-1,y})
                            flag[x-1][y]=1
                        }
                        if x+1<high &&grid[x+1][y]!=0&&flag[x+1][y]==0{
                            stack=append(stack,[]int{x+1,y})
                            flag[x+1][y]=1
                        }
                        if y+1<width &&grid[x][y+1]!=0&&flag[x][y+1]==0{
                            stack=append(stack,[]int{x,y+1})
                            flag[x][y+1]=1
                        }
                        if y-1>=0 &&grid[x][y-1]!=0&&flag[x][y-1]==0{
                            stack=append(stack,[]int{x,y-1})
                            flag[x][y-1]=1
                        }
                        stack=stack[1:]
                    }
                }
                if temp>ans{
                    ans=temp
                }
            }
        }
    }

    return ans
}