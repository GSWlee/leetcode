func partitionLabels(S string) []int {
	tag :=make(map[byte]int)
	length:=len(S)
	var ans []int
	flag:=make([]int,length)
	for j:=length-1;j>=0;j--{
		f,v:=tag[S[j]]
		if v{
			flag[j]=f
		}else{
			tag[S[j]]=j
			flag[j]=j
		}
	}
	i:=0
	for i<length{
		max:=flag[i]
		for ;i<max;i++{
			if flag[i]>max{
				max=flag[i]
			}
		}
		i++
		ans=append(ans,i)
	}
	for i:=len(ans)-1;i>=1;i--{
		ans[i]=ans[i]-ans[i-1]
	}
	return ans
}