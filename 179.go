type numbers []int

func (n numbers) Len() int {
	return len(n)
}

func (n numbers) Less(i int,j int) bool {
	a,b:=strconv.Itoa(n[i]),strconv.Itoa(n[j])
	lena,lenb:=len(a),len(b)
	if lena==lenb{
		return n[i]<n[j]
	}else{
        sa,sb:=10,10
        for i:=1;i<lena;i++{
            sa*=10
        }
        for i:=1;i<lenb;i++{
            sb*=10
        }
        return sb*n[i]+n[j]<sa*n[j]+n[i]
    }
}

func (n numbers) Swap(i int,j int)  {
	n[i],n[j]=n[j],n[i]
}

func largestNumber(nums []int) string {
	temp:=numbers(nums)
	sort.Sort(sort.Reverse(temp))
	ans:=""
	for _,v:=range temp{
		ans=ans+strconv.Itoa(v)
	}
    for len(ans)>1&&ans[0]=='0'{
        ans=ans[1:]
    }
	return ans
}