func kClosest(points [][]int, K int) [][]int {
    var dis func(node []int) float64
    dis=func(node []int) float64{
        x:=float64(node[0])
        y:=float64(node[1])
        return math.Sqrt(x*x+y*y)
    }
    var ans=[][]int{}
    var dist []float64
    var adjust func()float64
    adjust=func()float64{
        K:=len(ans)
        for i:=K-1;i>0;i--{
            if dist[i]<dist[i-1]{
                temp:=ans[i-1]
                ans[i-1]=ans[i]
                ans[i]=temp
                tempd:=dist[i-1]
                dist[i-1]=dist[i]
                dist[i]=tempd
            }else{
                break
            }
        }
        return dist[K-1]
    }
    var min float64
    for _,i:=range(points){
        if len(ans)<K{
            ans=append(ans,i)
            dist=append(dist,dis(i))
            min=adjust()
        }else{
            if dis(i)<min{
                ans[K-1]=i
                dist[K-1]=dis(i)
                min=adjust()
            }
        }
    }
    return ans
}