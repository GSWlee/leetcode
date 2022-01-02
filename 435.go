type IntSlice [][]int
 
func (s IntSlice) Len() int { return len(s) }
func (s IntSlice) Swap(i, j int){ s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { 
    if s[i][1]==s[j][1]{
        return s[i][0]<s[j][0]
    }else{
        return s[i][1] < s[j][1] 
    }
}


func eraseOverlapIntervals(intervals [][]int) int {
    sort.Sort(IntSlice(intervals))
    end:=intervals[0][1]
    ans:=0
    for k,v:=range intervals{
        if k==0{
            ans++
        }else{
            if v[0]<end{
                continue
            }else{
                ans++
                end=v[1]
            }
        }
    }
    return len(intervals)-ans
}