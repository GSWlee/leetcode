func insert(intervals [][]int, newInterval []int) [][]int {
    for i:=0;i<len(intervals);i++{
        if intervals[i][1]<newInterval[0]{
            continue
        }else if intervals[i][0]>newInterval[1]{
            intervals=append(intervals,newInterval)
            for j:=len(intervals)-1;j>i;j--{
                temp:=intervals[j]
                intervals[j]=intervals[j-1]
                intervals[j-1]=temp
            }
            return intervals
        }else{
            ans:=intervals[:i]
            j:=i
            if newInterval[0]>intervals[j][0]{
                newInterval[0]=intervals[j][0]
            }
            for ;j<len(intervals);j++{
                if intervals[j][0]<=newInterval[1]&&intervals[j][1]>=newInterval[1]{
                    newInterval[1]=intervals[j][1]
                    break
                }else if intervals[j][0]>newInterval[1]{
                    j--
                    break
                }
            }
            ans=append(ans,newInterval)
            for j=j+1;j<len(intervals);j++{
                ans=append(ans,intervals[j])
            }
            return ans
        }
    }
    return append(intervals,newInterval)
}