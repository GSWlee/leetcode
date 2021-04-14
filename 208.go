type Trie struct {
    sons map[string]*Trie
    end map[string]bool
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{sons:make(map[string]*Trie),end: make(map[string]bool)}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    if v,ok:=this.sons[word[0:1]];ok{
        if len(word)>1{
            v.Insert(word[1:])
        }else{
            this.end[word[0:1]]=true
        }
    }else{
        temp:=Constructor()
        this.sons[word[0:1]]=&temp
        if len(word)>1{
            this.sons[word[0:1]].Insert(word[1:])
        }else{
            this.end[word[0:1]]=true
        }
    }
    return
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    if word==""{
        return false
    }else{
        if v,ok:=this.sons[word[0:1]];ok{
            if len(word)>1{
                return v.Search(word[1:])
            }else{
                return this.end[word[0:1]]
            }
            
        }else{
            return false
        }
    }
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    if prefix==""{
        return true
    }else{
        if v,ok:=this.sons[prefix[0:1]];ok{
            return v.StartsWith(prefix[1:])
        }else{
            return false
        }
    }
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */