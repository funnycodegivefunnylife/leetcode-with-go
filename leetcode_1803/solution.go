package leetcode_1803

type Trie struct {
	children [2]*Trie
	value    int
}

func (t *Trie) insert(num int) {

	node := t

	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[bit] == nil {
			node.children[bit] = &Trie{}
		}
		node = node.children[bit]
	}

	node.value = num
}

func (t *Trie) find(num int) int {

	node := t

	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1
		if node.children[1-bit] != nil {
			node = node.children[1-bit]
		} else {
			node = node.children[bit]
		}
	}

	return node.value
}

func countPairs(nums []int, low int, high int) int {
	trie := &Trie{}

	count := 0

	for _, num := range nums {
		trie.insert(num)
	}

	return count
}
