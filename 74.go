func searchMatrix(matrix [][]int, target int) bool {
    m,n:=len(matrix[0]),len(matrix)
    var temp int
    for start,end:=0,n-1;start>=0&&end<n&&start<=end;{
        temp=(start+end)/2
        if temp==end{
            if matrix[temp][0]<target{
                break
            }else if matrix[temp][0]==target{
                return true
            }else{
                end=temp-1
                continue
            }
        }else{
            if matrix[temp][0]<target&&matrix[temp+1][0]>target{
                break
            }else if matrix[temp][0]==target{
                return true
            }else{
                if matrix[temp][0]<target{
                    start=temp+1
                    continue
                }else{
                    end=temp-1
                    continue
                }
            }
        }
    }
    for start,end:=1,m-1;start>0&&end<m&&start<=end;{
        t:=(start+end)/2
        if matrix[temp][t]==target{
            return true
        }else if matrix[temp][t]>target{
            end=t-1
        }else{
            start=t+1
        }
    }
    return false
}