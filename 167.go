// 利用哈希
func twoSum(numbers []int, target int) []int {
    dict:=map[int]int{}
    for k,v:=range numbers{
        dict[v]=k+1
    }
    for k,v:=range numbers{
        tar:=target-v
        if v,ok:=dict[tar];ok{
            return []int{k+1,v}
        }
    }
    return []int{}
}

//双指针
func twoSum(numbers []int, target int) []int {
    s,e:=0,len(numbers)-1
    for s<e{
        if numbers[s]+numbers[e]==target{
            return []int{s+1,e+1}
        }else if numbers[s]+numbers[e]>target{
            e--
        }else{
            s++
        }
    }
    return []int{}
}