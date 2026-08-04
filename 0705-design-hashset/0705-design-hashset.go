type MyHashSet struct {
	Hash []bool
}

func Constructor() MyHashSet {
	return MyHashSet{
		Hash: make([]bool, 1_000_001),
	}
}

func (this *MyHashSet) Add(key int) {
	if this.Hash[key] == false {
		this.Hash[key] = true
	}
}

func (this *MyHashSet) Remove(key int) {
	if this.Hash[key] == true {
		this.Hash[key] = false
	}
}

func (this *MyHashSet) Contains(key int) bool {
	if this.Hash[key] == false {
		return false
	}

	return true
}