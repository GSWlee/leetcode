func evalRPN(tokens []string) int {
    ans := []int{}
    for _,item :=range tokens{
        switch item{
            case "+":
                op1,op2:=ans[len(ans)-2],ans[len(ans)-1]
                temp:=op1+op2
                ans=ans[:len(ans)-1]
                ans[len(ans)-1]=temp
            case "-":
                op1,op2:=ans[len(ans)-2],ans[len(ans)-1]
                temp:=op1-op2
                ans=ans[:len(ans)-1]
                ans[len(ans)-1]=temp
            case "/":
                op1,op2:=ans[len(ans)-2],ans[len(ans)-1]
                temp:=op1/op2
                ans=ans[:len(ans)-1]
                ans[len(ans)-1]=temp
            case "*":
                op1,op2:=ans[len(ans)-2],ans[len(ans)-1]
                temp:=op1*op2
                ans=ans[:len(ans)-1]
                ans[len(ans)-1]=temp
            default:
                temp,_:=strconv.Atoi(item)
                ans=append(ans,temp)
        }
    }
    return ans[0]
}