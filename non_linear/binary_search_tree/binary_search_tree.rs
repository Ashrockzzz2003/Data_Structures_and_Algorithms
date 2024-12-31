use std::collections::VecDeque;

/*

Binary search trees, for any node x, the keys in the left subtree of x are at most x.key and the keys
at the right subtree are at least x.key

different binary search trees can represent the same set of values 

the worst case running time complexity for most tree operations are proportional to the height of the tree


main property: Let x be a node in a binary search tree, if y is a node in the left subtree of x, then y.key <= x.key
and if y is a node in the right subtree of x, then y.key >= x.key 

*/


/*

We recreate a simple binary search tree by creating a node , where the node has the following parameters:
- left child 
- right child 
- value of that particular node 

we create a binary search tree in this implementation iteratively by doing a level order traversal 
*/

#[derive(Debug)]
struct Node {
    val: i32,
    left_child: Option<Box<Node>>,
    right_child: Option<Box<Node>>
}

impl Node {
    fn new(value: i32) -> Self {
        Node {
            val: value, 
            left_child: None,
            right_child: None
        }
    }
}


#[derive(Debug)]
struct BinarySearchTree {
    root_node: Option<Box<Node>>
}

impl From<Node> for Option<Box<Node>> {
    fn from(node: Node) -> Self {
        Some(Box::new(node))
    }
}

impl BinarySearchTree {
    fn new() -> Self {
        BinarySearchTree {
            root_node: None
        }
    }
    
    /*
        insert_node sub-routine time complexity => amortized O(n) and worst case => O(h) where h is the height of the tree 
        Reason: Since we are trying to create a binary tree either out of a list or from a set of user inputs, and suppose the size or the number of inputs is "n"
        then the  insert_node will be called "n" times and worst case, the subroutine might travel till the end of the tree 
        And why isnt the worst O(log n) where log n is the height of the tree 
        Since we are not performing balancing , there are chances for the tree to be incredibly skewed either left or right side
        So we represnt the worst case O(h) given the chances for there being a skewness to the tree 
    */

    fn insert_node(&mut self, value:i32){
        if self.root_node.is_none() {
            self.root_node = Node::new(value).into();
            return;
        }

        let mut aux_queue: VecDeque<&mut Box<Node>> = VecDeque::new();
        let root = self.root_node.as_mut().unwrap();
        aux_queue.push_back(root);

        while let Some(node) = aux_queue.pop_front() {
            if value > node.val {
                let right = &mut  node.right_child;
                match right {
                    None => {
                        *right = Node::new(value).into();
                        return;
                    },
                    Some(n) => {
                        aux_queue.push_back(n);
                    }
                }
            } else if value < node.val {
                let left = &mut node.left_child;
                match left {
                    None => {
                        *left = Node::new(value).into();
                        return;
                    },
                    Some(n) => {
                        aux_queue.push_back(n);
                    }
                }
            }
        }
    } 

    /*
        Inorder traversal => Time complexity is O(n) 
        Reason: Most assume the time complexity is O(log n) as they forget the fact that we traverse not only left or right subtree but both and since we are traversing both left
        and right subtree, the time complexity becomes O(n). Another way to think is , we are simply trying to going and walk across all the nodes just in a different way
        we traverse from left -> parent -> right
    */

    fn inorder_traversal(&self, node: Option<&Box<Node>>) {
        if let Some(node) = node {
            self.inorder_traversal(node.left_child.as_ref());
            println!("{}", node.val);
            self.inorder_traversal(node.right_child.as_ref());
        } 
    }

    /*
        Pre order traversal => time complexity  is O(n)
        reasoning is the same as inorder traversal except for the walk as this traversal takes the walk in a different direction
        parent -> left -> right
    */

    fn preorder_traversal(&self, node: Option<&Box<Node>>){
        if let Some(node) = node {
            println!("{}", node.val);
            self.preorder_traversal(node.left_child.as_ref());
            self.preorder_traversal(node.right_child.as_ref());
        }
    }

    /*
        Post order traversal => Time complexity is O(n) 
        reason is the same as inorder traversal except for the walk as this traversal takes the walk in a different direction
        left -> right -> parent
    */

    fn postorder_traversal(&self, node: Option<&Box<Node>>){
        if let Some(node) = node {
            self.postorder_traversal(node.left_child.as_ref());
            self.postorder_traversal(node.right_child.as_ref());
            println!("{}", node.val);
        }
    }

    /*
        Level order traversal => Time complexity is O(n) 
        reason: here we simply try to visit each and every node that is present but in a level wise, so we go from level 0 that is root to level n where we have half of the elements in a tree  
    */

    fn level_order_traversal(&self, node: Option<&Box<Node>>){
        if let Some(root) = node {
            let mut aux_queue: VecDeque<&Box<Node>> = VecDeque::new();
            aux_queue.push_back(root);

            while let Some(node) = aux_queue.pop_front(){
                println!("{}", node.val);
                if let Some(left) = node.left_child.as_ref() {
                    aux_queue.push_back(left);
                }
                if let Some(right) = node.right_child.as_ref() {
                    aux_queue.push_back(right);
                }
            }
        }
    }
}

fn main() {
    let mut bst = BinarySearchTree::new();
    bst.insert_node(10);
    bst.insert_node(5);
    bst.insert_node(15);
    bst.insert_node(69);

    println!("{:#?}", bst);
    
    println!("Printing the node values as we do inorder traversal ...");
    bst.inorder_traversal(bst.root_node.as_ref());

    println!("Printing the node values as we do the post order traversal...");
    bst.postorder_traversal(bst.root_node.as_ref());

    println!("Printing the node values as we do the pre order traversal...");
    bst.preorder_traversal(bst.root_node.as_ref());

    println!("Printing the node values as we do the level order traversal...");
    bst.level_order_traversal(bst.root_node.as_ref());
}
    

