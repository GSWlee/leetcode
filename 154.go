func findMin(nums []int) int {
    length:=len(nums)
    start,end:=0,length-1
    var find func(start int,end int) (int,bool)
    find=func(start int,end int) (int,bool){
        temp:=(start+end)/2
        rf,lf:=true,true
        if start>end{
            return 0,false
        }
        if temp==(length-1){
            return nums[0],true
        }
        if nums[temp]>nums[temp+1]{
            return nums[temp+1],true
        }
        right,left:=0,0
        if nums[start]<nums[temp]{
            right=nums[start]
        }else{
            if v,ok:=find(start,temp-1);ok{
                right=v
            }else{
                rf=false
            }
        }
        if nums[temp]<nums[end]{
            left=nums[temp]
        }else{
            if v,ok:=find(temp+1,end);ok{
                left=v
            }else{
                lf=false
            }
        }
        if rf&&lf{
            if right<left{
                return right,true
            }else{
                return left,true
            }
        }else if rf&&!lf{
            return right,true
        }else if !rf&&lf{
            return left,true
        }else{
            return 0,false
        }
    }
    v,_:=find(start,end)
    return v
}