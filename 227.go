func calculate(s string) int {
	stack :=[]byte{}
	news:=[]string{}
	nums:=""
    if s[0]=='-'{
        s="0"+s
    }
	for i:=range s{
		if s[i]==' '{
			continue
		}else if strings.Contains("+-",s[i:i+1]){
			if nums!=""{
				news=append(news,nums)
				nums=""
			}
			for len(stack)>0{
				news=append(news,string(stack[len(stack)-1]))
				stack=stack[:len(stack)-1]
			}
			stack=append(stack,s[i])
		}else if strings.Contains("*/",s[i:i+1]){
            if nums!=""{
				news=append(news,nums)
				nums=""
			}
            for len(stack)>0&&stack[len(stack)-1]!='+'&&stack[len(stack)-1]!='-'{
				news=append(news,string(stack[len(stack)-1]))
				stack=stack[:len(stack)-1]
			}
			stack=append(stack,s[i])
        }else{
			nums+=s[i:i+1]
		}
	}
	if nums!=""{
		news=append(news,nums)
		nums=""
	}
	for len(stack)>0{
		news=append(news,string(stack[len(stack)-1]))
		stack=stack[:len(stack)-1]
	}
	ans:=[]int{}
	for _,item:=range news{
		if strings.Contains("+-*/",item){
			if item=="+"{
				temp:=ans[len(ans)-1]+ans[len(ans)-2]
				ans=ans[:len(ans)-1]
				ans[len(ans)-1]=temp
			}else if item=="-"{
				temp:=ans[len(ans)-2]-ans[len(ans)-1]
				ans=ans[:len(ans)-1]
				ans[len(ans)-1]=temp
			}else if item=="*"{
                temp:=ans[len(ans)-2]*ans[len(ans)-1]
				ans=ans[:len(ans)-1]
				ans[len(ans)-1]=temp
            }else{
                temp:=ans[len(ans)-2]/ans[len(ans)-1]
				ans=ans[:len(ans)-1]
				ans[len(ans)-1]=temp
            }
		}else {
			temp,_:=strconv.Atoi(item)
			ans=append(ans,temp)
		}
	}
	return ans[0]
}