//利用栈解决
func removeDuplicates(S string) string {
	stack:= []byte{}
    for i:=range S{
        if len(stack)>0&&stack[len(stack)-1]==S[i]{
            stack=stack[:len(stack)-1]
        }else{
            stack=append(stack,S[i])
        }
    }
	return string(stack)
}

//用循环解决，垃圾中的垃圾
func removeDuplicates(S string) string {
	ans := S
	flag := true
	for flag{
		flag=false
		temp:=""
		target:=0
		for target<len(ans){
			if target+1<len(ans){
				if ans[target+1]!=ans[target]{
					temp+=ans[target:target+1]
					target++
				}else{
					flag=true
					target=target+2
				}
			}else{
				temp += ans[target : target+1]
				target++
			}
		}
		if flag{
			ans=temp
		}
	}
	return ans
}