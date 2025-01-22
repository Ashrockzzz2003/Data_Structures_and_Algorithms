use std::collections::VecDeque;

/*
    Binary tree is a linked data structure where one node that has some value is linked to 0, 1, 2 other nodes 
    this is the base or the boiler plat data structure from which several changes and additions have been made to make some of the most complex and important data  
    strcutres like the quad tree, B+ tree, and so on... 

    A simple binary tree has a root node which is then connected 0,1,2 nodes 
    and the connected nodes are either the left child of the node or the right child of the node 
    in case a tree has no root node, then the supposedly "tree" is not a tree data stucture 

    in this implementaiton we have node struct which then we will point to or point from other nodes to make a simple bianry tree
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
struct BinaryTree {
    root_node: Option<Box<Node>>
}

impl From<Node> for Option<Box<Node>> {
    fn from(node: Node) -> Self {
        Some(Box::new(node))
    }
}

impl BinaryTree {
    fn new() -> Self {
        BinaryTree {
            root_node: None
        }
    }

    /*
        Here the insert_node sub routine simply adds or inserts a node to the binary tree but the additions are made in such a way that they are added from left 
        to right, and since we are inserting nodes and these insertions can either take place n times or the single insertion itself can take n time 
        the amortized time complexity for this operation is O(n) where n is the number of nodes 
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
            let right = &mut node.right_child;
            match right {
                None => {
                    *right = Node::new(value).into();
                    return;
                },
                Some(n) => {
                    aux_queue.push_back(n);
                }
            }
        }
    } 

    /*
        inorder traversal => time complexity is O(n) where n is the number of nodes
        reason: since we have to walk to every single node in the tree 
        path direction: left -> parent -> right
    */

    fn inorder_traversal(&self, node: Option<&Box<Node>>) {
        if let Some(node) = node {
            self.inorder_traversal(node.left_child.as_ref());
            println!("{}", node.val);
            self.inorder_traversal(node.right_child.as_ref());
        } 
    }

    /*
        pre order traversal => Time complextiy is O(n) where n is the number of nodes 
        reason: same as in inorder traversal
        path direction: parent -> left -> right
    */

    fn preorder_traversal(&self, node: Option<&Box<Node>>){
        if let Some(node) = node {
            println!("{}", node.val);
            self.preorder_traversal(node.left_child.as_ref());
            self.preorder_traversal(node.right_child.as_ref());
        }
    }

    /*
        post order traversal => time complexity is O(n) where n is the number of nodes 
        reason: same as in the inorder traversal 
        path direction => left -> right -> parent
    */

    fn postorder_traversal(&self, node: Option<&Box<Node>>){
        if let Some(node) = node {
            self.postorder_traversal(node.left_child.as_ref());
            self.postorder_traversal(node.right_child.as_ref());
            println!("{}", node.val);
        }
    }

    /*
        level order traversal => time complexity is O(n) where n is the number of ndoes 
        reason: same as in the inorder traversal 
        path direction: here we go or walk through nodes in level order, meaning from level 0 where the root node is and then till level n where half of the nodes in the tree is present
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
    let mut bt = BinaryTree::new();
    bt.insert_node(10);
    bt.insert_node(5);
    bt.insert_node(15);
    bt.insert_node(69);

    println!("{:#?}", bt);
    
    println!("Printing the node values as we do inorder traversal ...");
    bt.inorder_traversal(bt.root_node.as_ref());

    println!("Printing the node values as we do the post order traversal...");
    bt.postorder_traversal(bt.root_node.as_ref());

    println!("Printing the node values as we do the pre order traversal...");
    bt.preorder_traversal(bt.root_node.as_ref());

    println!("Printing the node values as we do the level order traversal...");
    bt.level_order_traversal(bt.root_node.as_ref());
}
    


