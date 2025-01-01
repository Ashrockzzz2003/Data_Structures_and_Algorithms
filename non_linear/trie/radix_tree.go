package radixtree
type Node struct {
	Key rune
	Children map[rune]*Node
	IsComplete bool
}


func NewNode(key rune) *Node {
	return &Node{
		Key: key,
		Children: make(map[rune]*Node),
	}
}


func (n *Node) AddChild(key rune, isComplete bool) *Node {
	child, exists := n.Children[key]
	if !exists {
		child = NewNode(key)
		n.Children[key] = child
	}
	if isComplete {
		child.IsComplete = true
	}
	return child
}


func (n *Node) GetChild(key rune) *Node {
	return n.Children[key]
}


func (n *Node) RemoveChild(key rune) {
	delete(n.Children, key)
}
type RadixTree struct {
	Head *Node
}


func NewRadixTree() *RadixTree {
	return &RadixTree{
		Head: NewNode('*'),
	}
}


func (t *RadixTree) Insert(word string) {
	currentNode := t.Head
	for i, char := range word {
		isComplete := i == len(word)-1
		currentNode = currentNode.AddChild(char, isComplete)
	}
}


func (t *RadixTree) Delete(word string) {
	var depthFirstDelete func(currentNode *Node, index int)
	depthFirstDelete = func(currentNode *Node, index int) {
		if index >= len(word) {
			return
		}
		char := rune(word[index])
		nextNode := currentNode.GetChild(char)
		if nextNode == nil {
			return
		}
		depthFirstDelete(nextNode, index+1)
		if index == len(word)-1 {
			nextNode.IsComplete = false
		}
		if len(nextNode.Children) == 0 && !nextNode.IsComplete {
			currentNode.RemoveChild(char)
		}
	}
	depthFirstDelete(t.Head, 0)
}


func (t *RadixTree) Find(word string) bool {
	node := t.getLastCharacterNode(word)
	return node != nil && node.IsComplete
}


func (t *RadixTree) getLastCharacterNode(word string) *Node {
	currentNode := t.Head
	for _, char := range word {
		currentNode = currentNode.GetChild(char)
		if currentNode == nil {
			return nil
		}
	}
	return currentNode
}
