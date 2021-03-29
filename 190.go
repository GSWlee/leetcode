func reverseBits(num uint32) uint32 {
    var ans uint32
    ans=0
    for i:=0;i<32;i++{
        if num%2==1{
            ans=ans<<1
            ans++
        }else{
            ans=ans<<1
        }
        num=num>>1
    }
    return ans
}