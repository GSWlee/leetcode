func kidsWithCandies(candies []int, extraCandies int) []bool {
    if candies==nil{
        return nil
    }else{
        max:=candies[0]
        for _,i:=range(candies){
            if i>max{
                max=i
            }
        }
        ans:=[]bool{}
        for j,_:=range(candies){
            if candies[j]+extraCandies>=max{
                ans=append(ans,true)
            }else{
                ans=append(ans,false)
            }
        }
        return ans
    }
}