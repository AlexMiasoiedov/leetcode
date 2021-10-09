package prefixtrie

type Trie struct {
    IsRoot bool
    IsWord bool
    Nodes map[byte]*Trie
}


func Constructor() Trie {
    var trie *Trie = new(Trie)
    trie.IsRoot = true
    return *trie
}


func (this *Trie) Insert(word string)  {
    if len(word) < 1 { 
        this.IsWord = true
        return
    }

    if this.Nodes == nil {
        this.Nodes = make(map[byte]*Trie)
    }

    if _, ok := this.Nodes[word[0]]; !ok {
        this.Nodes[word[0]] = new(Trie)
    }

    this.Nodes[word[0]].Insert(word[1:])
}


func (this *Trie) Search(word string) bool {
    if len(word) < 1 { return this.IsWord }
    
    if _, ok := this.Nodes[word[0]]; !ok { return false }
    
    return this.Nodes[word[0]].Search(word[1:])
}


func (this *Trie) StartsWith(prefix string) bool {
    if len(prefix) < 1 { return true }
       
    if _, ok := this.Nodes[prefix[0]]; ok {
        return this.Nodes[prefix[0]].StartsWith(prefix[1:])
    } else {
        return false
    }    
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

