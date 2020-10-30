func islandPerimeter(grid [][]int) int {
    sizex:=len(grid)
    var sizey int
    if sizex!=0{
        sizey=len(grid[0])
    }else{
        return 0
    }
    ans:=0
    var edges func(i,j int)
    edges=func(i,j int){
        if grid[i][j]==0{
            return
        }
        if i==0{
            ans++
        }else{
            if grid[i-1][j]==0{
                ans++
            }
        }
        if i==sizex-1{
            ans++
        }else{
            if grid[i+1][j]==0{
                ans++
            }
        }
        if j==sizey-1{
            ans++
        }else{
            if grid[i][j+1]==0{
                ans++
            }
        }
        if j==0{
            ans++
        }else{
            if grid[i][j-1]==0{
                ans++
            }
        }
    }
    for m:=0;m<sizex;m++{
        for n:=0;n<sizey;n++{
            edges(m,n)
        }
    }
    return ans
}