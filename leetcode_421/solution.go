package leetcode_421

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

func findMaximumXOR(nums []int) int {

	trie := &Trie{}

	for _, num := range nums {
		trie.insert(num)
	}

	maxXor := 0
	for _, num := range nums {
		maxXor = maxInt(maxXor, xor(num, trie.find(num)))
	}

	return maxXor
}

func maxInt(a, b int) int {

	if a > b {
		return a
	}
	return b
}

func xor(a, b int) int {

	return a ^ b
}
