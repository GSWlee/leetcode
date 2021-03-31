//低端版本，先排序，后去重
type sortInt [][]int

func (s sortInt) Len() int {
	return len(s)
}

func (s sortInt)Swap(i,j int)  {
	s[i],s[j]=s[j],s[i]
}

func (s sortInt) Less(i,j int) bool {
	if len(s[i])<len(s[j]){
		return true
	}else if len(s[i])>len(s[j]){
		return false
	}else{
		length:=len(s[i])
		for m:=0;m<length;m++{
			if s[i][m]<s[j][m]{
				return true
			}else if s[i][m]>s[j][m]{
				return false
			}else {
				continue
			}
		}
	}
	return true
}

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

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	temp:=subsets(nums)
	sort.Sort(sortInt(temp))
	length:=len(temp)
	ans:=make([][]int,0)
	for i:=0;i<length;i++{
		if i==length-1||!reflect.DeepEqual(temp[i],temp[i+1]){
			ans=append(ans,temp[i])
		}
	}
	return ans
}

//利用深度优先加回溯法
func subsetsWithDup(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t:=[]int{}
	var dfs func(prechose bool, cur int)
	dfs= func(prechose bool, cur int) {
		if cur==len(nums){
			ans=append(ans,append([]int{},t...))
			return
		}
		dfs(false,cur+1)
		if !prechose&&cur>0&&nums[cur]==nums[cur-1]{
			return
		}
		t=append(t,nums[cur])
		dfs(true,cur+1)
		t=t[:len(t)-1]
	}
	dfs(false,0)
	return ans
}