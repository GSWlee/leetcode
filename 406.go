func reconstructQueue(people [][]int) [][]int {
    for i:=len(people)-1;i>0;i--{
        for j:=0;j<i;j++{
            if people[j][0]>people[j+1][0]{
                temp1,temp2:=people[j][0],people[j][1]
                people[j][0],people[j][1]=people[j+1][0],people[j+1][1]
                people[j+1][0],people[j+1][1]=temp1,temp2
            }else if people[j][0]==people[j+1][0]&&people[j][1]<people[j+1][1]{
                temp1,temp2:=people[j][0],people[j][1]
                people[j][0],people[j][1]=people[j+1][0],people[j+1][1]
                people[j+1][0],people[j+1][1]=temp1,temp2
            }
        }
    }
    ans:=make([][]int,len(people))
    for i:=0;i<len(people);i++{
        ans[i]=make([]int,2)
        ans[i][1]=-1
    }
    for i:=0;i<len(people);i++{
        sum:=people[i][1]
        for j:=0;j<len(ans);j++{
            if ans[j][1]!=-1{
                continue
            }else{
                if sum==0{
                    ans[j][0],ans[j][1]=people[i][0],people[i][1]
                    break
                }else {
                    sum--
                }
            }
        }
    }
    return ans
}