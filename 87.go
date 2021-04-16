func isScramble(s1 string, s2 string) bool {
    lens1:=len(s1)
    lens2:=len(s2)
    if lens1!=lens2{
        return false
    }
    dp:=make([][][]int,lens1)
    for i:=range dp{
        dp[i]=make([][]int,lens2)
        for j:=range dp[i]{
            dp[i][j]=make([]int,lens1+1)
        }
    }
    var find func(i int,j int,k int) int
    find=func(i int,j int,k int) int{
        if dp[i][j][k]!=0{
            return dp[i][j][k]
        }
        if lens1-i<k||lens1-j<k{
            dp[i][j][k]=-1
            return -1
        }else if k==1{
            if s1[i]==s2[j]{
                dp[i][j][k]=1
                return 1
            }
        }else{
            dp[i][j][k]=-1
            for m:=1;m<k;m++{
                x1,y1:=find(i,j,m),find(i+m,j+m,k-m)
                x2,y2:=find(i,j+k-m,m),find(i+m,j,k-m)
                if x1==1&&y1==1{
                    dp[i][j][k]=1
                    return 1
                }
                if x2==1&&y2==1{
                    dp[i][j][k]=1
                    return 1
                }
            }
            return -1
        }
        return -1
    }
    if find(0,0,lens1)==1{
        return true
    }
    return false
}