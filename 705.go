type MyHashSet struct {
    hashmap [20][]int
}


/** Initialize your data structure here. */
func Constructor() MyHashSet {
    return MyHashSet{hashmap: [20][]int{}}
}

func (this *MyHashSet) hashnum(key int)int{
    return key%20
}

func (this *MyHashSet) Add(key int)  {
    flag:=this.hashnum(key)
    for _,item:=range this.hashmap[flag]{
        if item==key{
            return
        }
    }
    this.hashmap[flag]=append(this.hashmap[flag],key)
}


func (this *MyHashSet) Remove(key int)  {
    flag:=this.hashnum(key)
    target:=-1
    for i,item:=range this.hashmap[flag]{
        if item==key{
            target=i
            break
        }
    }
    if target!=-1{
        this.hashmap[flag]=append(this.hashmap[flag][:target],this.hashmap[flag][target+1:]...)
    }
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    flag:=this.hashnum(key)
    for _,item:=range this.hashmap[flag]{
        if item==key{
            return true
        }
    }
    return false
}