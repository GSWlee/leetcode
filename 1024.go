func sort(arr [][]int)[][]int{
    for i:=len(arr)-1;i>=0;i--{
        for j:=0;j<i;j++{
            if arr[j][0]>arr[j+1][0]{
                temp1,temp2:=arr[j][0],arr[j][1]
                arr[j][0],arr[j][1]=arr[j+1][0],arr[j+1][1]
                arr[j+1][0],arr[j+1][1]=temp1,temp2
            }
        }
    }
    return arr
}

func min(a,b int)int{
    if a<b{
        return a
    }else{
        return b
    }
}

func videoStitching(clips [][]int, T int) int {
    dp:=make([]int,101)
    for i:=0;i<101;i++{
        dp[i]=101
    }
    dp[0]=0
    clips=sort(clips)
    for i:=0;i<len(clips);i++{
        start,end:=clips[i][0],clips[i][1]
        if dp[start]==101{
            break
        }else{
            for j:=start+1;j<=end;j++{
                dp[j]=min(dp[j],dp[start]+1)
            }
        }
    }
    if dp[T]==101{
        return -1
    }
    return dp[T]
}