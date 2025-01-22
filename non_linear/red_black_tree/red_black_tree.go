package main

import (
	"fmt"
)

// Red Black Trees
// 1. A node is either red or black
// 2. The root and leaves(NIL) are black
// 3. If a node is red, then it's children are black
// 4. All paths frm node to it's NIL descendants contain the same no.of black nodes

var RED = false
var BLACK = true

type RBNode struct {
	parent *RBNode
	left   *RBNode
	right  *RBNode
	val    int
	colour bool //True for black False fr Red
}

type RBTree struct {
	root     *RBNode
	sentinel *RBNode //All nil nodes and the root's parent are RBTree.sentinel
}

/*
functions:
	insert
	delete
	search
	left-rotate
	right-rotate
*/

func (tree *RBTree) leftRotate(x *RBNode) {
	y := x.right
	x.right = y.left

	// If y wasn't a leaf node, y's left child's parent is now x
	if y.left != tree.sentinel {
		y.left.parent = x
	}

	y.parent = x.parent

	// If x was the root of the tree, make y the new root
	// If x was a left child, make y a left child
	// If x was a right child, make y a right child
	if x.parent == tree.sentinel {
		tree.root = y
	} else if x.parent.right == x {
		x.parent.right = y
	} else {
		x.parent.left = y
	}

	y.left = x
	x.parent = y

}

// Mirror of leftRotate
func (tree *RBTree) rightRotate(x *RBNode) {
	y := x.left
	x.left = y.right

	if y.right != tree.sentinel {
		y.right.parent = x
	}

	y.parent = x.parent
	if x.parent == tree.sentinel {
		tree.root = y
	} else if x.parent.left == x {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.right = x
	x.parent = y

}

func (tree *RBTree) insertNode(val int) {
	// Creating the node to be inserted
	// It's a leaf node, and while insertion the node will alwasy be coloured RED
	node := &RBNode{
		left:   tree.sentinel,
		right:  tree.sentinel,
		val:    val,
		colour: RED,
	}

	itr := tree.root
	itrParent := tree.sentinel
	for itr != tree.sentinel {
		itrParent = itr
		if node.val == itr.val {
			return
		} else if node.val < itr.val {
			itr = itr.left
		} else {
			itr = itr.right
		}
	}

	node.parent = itrParent
	// Making appropriate changes for the cases where
	// 1. Tree was empty before insertion
	// 2. node is a right child
	// 3. node is a left child
	if itrParent == tree.sentinel {
		tree.root = node
	} else if itrParent.val < node.val {
		itrParent.right = node
	} else {
		itrParent.left = node
	}

	// Balance/Correct the RB Tree after insertion
	tree.balanceTreeInsertion(node)

}

// After insertion,
// 1. If parent of the node is BLACK. The tree retains it's properties and you don't need to do anything
// 2. If parent of the node is RED and the parent is a left child:
// 		2.1 If z's UNCLE is also red
// 		2.2 If z's UNCLE is black
// 2. If parent of the node is RED and the parent is a right child - Mirror of 2

func (tree *RBTree) balanceTreeInsertion(node *RBNode) {
	for node.parent.colour == RED {

		// If the node's parent is a left chid
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			// If uncle is red, recolour
			if uncle.colour == RED {
				node.parent.colour = BLACK
				uncle.colour = BLACK
				node.parent.parent.colour = RED
				node = node.parent.parent
			} else {
				// If node is a right child, perform an extra left rotation on its parent
				// Then, change colours and perform a right rotation on its grand parent
				if node == node.parent.right {
					// Changing node to node.parent so after the rotation, the node will remain a leaf node
					node = node.parent
					tree.leftRotate(node)
				}
				node.parent.colour = BLACK
				node.parent.parent.colour = RED
				tree.rightRotate(node.parent.parent)
			}
		} else {
			// Same as the above, with right and left exchanged
			uncle := node.parent.parent.left
			if uncle.colour == RED {
				node.parent.colour = BLACK
				uncle.colour = BLACK
				node.parent.parent.colour = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.rightRotate(node)
				}
				node.parent.colour = BLACK
				node.parent.parent.colour = RED
				tree.leftRotate(node.parent.parent)
			}
		}
	}

	// To ensure the root is always black
	tree.root.colour = BLACK
}

// Basic BST search
func (tree *RBTree) searchVal(val int) *RBNode {
	itr := tree.root
	for itr != tree.sentinel {
		if val == itr.val {
			return itr
		} else if val < itr.val {
			itr = itr.left
		} else {
			itr = itr.right
		}
	}
	return nil
}

// To find the successor for deleteNode
func (tree *RBTree) minVal(node *RBNode) *RBNode {
	for node.left != tree.sentinel {
		node = node.left
	}
	return node
}

// Changing the subtree rooted at u with the subtree rooted at v
// A helper function for deleteNode
func (tree *RBTree) transplant(u *RBNode, v *RBNode) {
	if u.parent == tree.sentinel {
		tree.root = v
	} else if u.parent.left == u {
		u.parent.left = v
	} else {
		u.parent.right = v
	}
	v.parent = u.parent
}

func (tree *RBTree) deleteNode(val int) {
	node := tree.searchVal(val)
	var x *RBNode //This is the new node that replaces the deleted node
	// If the node to be deleted doesn't exist in the tree, do nothing and return
	if node == nil {
		return
	}

	originalColour := node.colour
	// If the node is a leaf node OR has only right child, replace the node by it's right child
	if node.left == tree.sentinel {
		x = node.right
		tree.transplant(node, node.right)
	} else if node.right == tree.sentinel {
		// If the node only has a left child, replace the node with it's left child
		x = node.left
		tree.transplant(node, node.left)
	} else {
		// If the node has both children, replace it with it's successor(smallest element in the right subtree / largest element in the left subtree)
		// Here, we replace with the smallest element in the right subtree
		y := tree.minVal(node.right)
		originalColour = y.colour
		x = y.right

		// If y isn't a leaf node
		if node.right != y {
			tree.transplant(y, y.right)
			// Giving node's right child to y after transplant
			y.right = node.right
			y.right.parent = y
		} else {
			// In-case x is tree.sentinel, ie, y has no right child
			// Since x moves into y’s original position, the attribute x.parent must be set correctly
			x.parent = y
		}

		tree.transplant(node, y)
		// Give left child of the node to y, which had no left child(as it's the minimun in the right-subtree of the node)
		y.left = node.left
		y.left.parent = y
		y.colour = node.colour
		// In this case, we are actually replacing the node to be deleted with y(by changing it's colour to be the node's colour)
		// This is the same as just changing the node's value to y's value, but I wanted to actually delete the node
		// Technically both have the same effect
		// The actual node being "deleted" is y, as it gets replaced by it's right child, and that's why x=y.right
	}

	// Only if the original colour of the node was black will there be a problem in the properties of red-black tree
	if originalColour == BLACK {
		tree.balanceTreeDeletion(x)
	}
}

// 4 Cases
// x’s sibling w is red
// x’s sibling w is black, and both of w’s children are black
// x’s sibling w is black, w’s left child is red, and w’s right child is black
// x’s sibling w is black, and w’s right child is red
func (tree *RBTree) balanceTreeDeletion(x *RBNode) {
	for x != tree.root && x.colour == BLACK {
		if x == x.parent.left {
			sibling := x.parent.right
			// Case 1
			if sibling.colour == RED {
				sibling.colour = BLACK
				x.parent.colour = RED
				tree.leftRotate(x.parent)
				sibling = x.parent.right
			}

			// Case 2
			if sibling.left.colour == BLACK && sibling.right.colour == BLACK {
				sibling.colour = RED
				x = x.parent
			} else {
				// Case 3
				if sibling.right.colour == BLACK {
					sibling.left.colour = BLACK
					sibling.colour = RED
					tree.rightRotate(sibling)
					sibling = x.parent.right
				}
				// Case 4
				sibling.colour = x.parent.colour
				x.parent.colour = BLACK
				sibling.right.colour = BLACK
				tree.leftRotate(x.parent)
				x = tree.root
			}
		} else {
			// Mirror of the above cases
			sibling := x.parent.left
			if sibling.colour == RED {
				sibling.colour = BLACK
				x.parent.colour = RED
				tree.rightRotate(x.parent)
				sibling = x.parent.left
			}

			if sibling.right.colour == BLACK && sibling.left.colour == BLACK {
				sibling.colour = RED
				x = x.parent
			} else {
				if sibling.left.colour == BLACK {
					sibling.right.colour = BLACK
					sibling.colour = RED
					tree.leftRotate(sibling)
					sibling = x.parent.left
				}
				sibling.colour = x.parent.colour
				x.parent.colour = BLACK
				sibling.left.colour = BLACK
				tree.rightRotate(x.parent)
				x = tree.root
			}
		}
	}
	x.colour = BLACK
}

// Basic inOrder Traversal
func (tree *RBTree) inOrderTraversal(node *RBNode) []int {
	if node == tree.sentinel {
		return make([]int, 0)
	}
	left := tree.inOrderTraversal(node.left)
	right := tree.inOrderTraversal(node.right)
	lst := make([]int, len(left)+len(right)+1)
	i := 0
	for ; i < len(left); i++ {
		lst[i] = left[i]
	}
	lst[i] = node.val
	i++
	for j := 0; j < len(right); j++ {
		lst[i+j] = right[j]
	}
	return lst
}

func main() {
	// Initializing
	rbTree := RBTree{
		sentinel: &RBNode{
			colour: BLACK,
		},
	}
	rbTree.root = rbTree.sentinel

	// Insertion
	rbTree.insertNode(10)
	rbTree.insertNode(20)
	rbTree.insertNode(15)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
	// Insert with Balancing
	rbTree.insertNode(25)
	rbTree.insertNode(5)
	rbTree.insertNode(1)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
	// Insert Duplicate
	rbTree.insertNode(15)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))

	// Deletion
	// Leaf
	rbTree.deleteNode(1)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
	// One Child
	rbTree.deleteNode(25)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
	// 2 Children
	rbTree.deleteNode(10)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
	// Deleting the root node
	rbTree.deleteNode(rbTree.root.val)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
	// Delete All Nodes
	rbTree.deleteNode(5)
	rbTree.deleteNode(20)
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))

	// Scaling to 100 elements
	// Insertion
	for i := 1; i <= 100; i++ {
		rbTree.insertNode(i)
	}
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))

	// Deletion
	for i := 1; i <= 100; i++ {
		rbTree.deleteNode(i)
	}
	fmt.Println(rbTree.inOrderTraversal(rbTree.root))
}
