type Node struct {
    Key int
    Value int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

func (ll *LinkedList) Get(key int) int {
    current := ll.Head
    for current != nil {
        if current.Key == key {
            return current.Value
        }
        current = current.Next
    }
    return -1
}

func(ll *LinkedList) Put(key int, value int) {
    current := ll.Head
    for current != nil {
        if current.Key == key {
            current.Value = value
            return
        }
        current = current.Next
    }
    newNode := &Node{
        Key: key,
        Value: value,
        Next: ll.Head,
    }
    ll.Head = newNode
}

func (ll *LinkedList) Remove(key int) {
    if ll.Head == nil {
        return
    }
    if ll.Head.Key == key {
        ll.Head = ll.Head.Next
        return
    }
    current := ll.Head
    for current.Next != nil {
        if current.Next.Key == key {
            current.Next = current.Next.Next
            return
        }
        current = current.Next
    }
} 

type MyHashMap struct {
    n int
    buckets []*LinkedList
}


func Constructor() MyHashMap {
    n := 991
    buckets := make([]*LinkedList, n)
    for i := range buckets {
        buckets[i] = &LinkedList{}
    }
    return MyHashMap {
        n: n,
        buckets: buckets, 
    }
}

func (this *MyHashMap) hash(x int) int {
    return x % this.n
}

func (this *MyHashMap) Put(key int, value int)  {
    n := this.hash(key)
    this.buckets[n].Put(key, value)
}


func (this *MyHashMap) Get(key int) int {
    n := this.hash(key)
    return this.buckets[n].Get(key)   
}


func (this *MyHashMap) Remove(key int)  {
    n := this.hash(key)
    this.buckets[n].Remove(key)      
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */