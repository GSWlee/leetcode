func subsets(nums []int) [][]int {
	length:=len(nums)
	max:=math.Pow(2,float64(length))
	var m int=int(max)
	ans:=make([][]int,0)
	for i:=0;i<m;i++{
		temp:=make([]int,0)
		j:=i
		q:=0
		for j!=0{
			if j%2==1{
				temp=append(temp,nums[q])
			}
			q++
            j=j>>1
		}
		ans=append(ans,temp)
	}
	return ans
}