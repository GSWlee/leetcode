//慢
func trap(height []int) int {
    max:=0
    for _,v := range height{
        if v>max{
            max=v
        }
    }
    ans:=0
    for i:=0;i<=max;i++{
        l,r,sum:=0,0,0
        flag:=true
        for m,n:=range height{
            if n>i{
                if flag{
                    flag=false
                    l=m
                }
                r=m
                sum++
            }
        }
        if !flag&&r>l{
            ans=ans+r-l-sum+1
        }else{
            break;
        }
    }
    return ans
} 

//单调栈(冗余版)
func trap(height []int) int {
	ans:=0
	stack:=make([]int,0)
	length:=len(height)
	for i:=0;i<length;i++{
		if len(stack)==0&&height[i]==0{
			continue
		}else{
			if len(stack)==0{
				stack=append(stack,i)
			}else{
				if height[stack[len(stack)-1]]>height[i]{
					stack=append(stack,i)
				}else{
					for len(stack)>0&&height[stack[len(stack)-1]]<=height[i]{
						if len(stack)==1{
							stack=stack[:len(stack)-1]
						}else{
							l1:=stack[len(stack)-2]
							if height[l1]>=height[i]{
								ans=ans+(height[i]-height[stack[len(stack)-1]])*(i-l1-1)
								stack=stack[:len(stack)-1]
							}else{
								ans=ans+(height[l1]-height[stack[len(stack)-1]])*(i-l1-1)
								stack=stack[:len(stack)-1]
							}
						}
					}
					stack=append(stack,i)
				}
			}
		}
	}
	return ans
}
//单调栈（精美版）
func trap(height []int) (ans int) {
    stack := []int{}
    for i, h := range height {
        for len(stack) > 0 && h > height[stack[len(stack)-1]] {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if len(stack) == 0 {
                break
            }
            left := stack[len(stack)-1]
            curWidth := i - left - 1
            curHeight := min(height[left], h) - height[top]
            ans += curWidth * curHeight
        }
        stack = append(stack, i)
    }
    return
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
