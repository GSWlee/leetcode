func countBits(num int) []int {
	ans:=make([]int,num+1)
	for i:=1;i<=num;i++{
        if i&1==1{
            ans[i]=ans[i-1]+1
        }else{
            ans[i]=ans[i/2]
        }
    }
	return ans
}