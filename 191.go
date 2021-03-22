func hammingWeight(num uint32) int {
    var ans int
    for num!=0{
        if num%2==1{
            ans++
            num--
        }
        num=num/2
    }
    return ans
}