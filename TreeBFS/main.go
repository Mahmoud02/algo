package main

import (
	"fmt"
	"math"
	"time"
)

/*
This pattern is based on the Breadth First Search (BFS) technique to traverse a tree
and uses a queue to keep track of all the nodes of a level before jumping onto the next level.
#Any problem involving the traversal of a tree in a level-by-level order can be efficiently solved using this approach.
How to identify the Tree BFS pattern:
If you’re asked to traverse a tree in a level-by-level fashion (or level order traversal)
1-Binary Tree Level Order Traversal (easy)
2-Zigzag Traversal (medium)
*/

type Node struct {
	Val   int
	left  *Node
	right *Node
}
type BinaryTree struct {
	root *Node
}

func main() {
	/*tree := &BinaryTree{}
	tree.insert(8).
		insert(4).
		insert(10).
		insert(2).
		insert(1).
		insert(3).
		insert(6).
		insert(5).
		insert(7).
		insert(9).
		insert(11)
	inOrderTraverse(tree.root)
	println()
	printPreOrder(tree.root)
	println()
	printPostOrder(tree.root)
	println()
	println(tree.Min())
	println(tree.Max())
	println(tree.Search(10))*/
	tree := &BinaryTree{}
	tree.insert(50).
		insert(30).
		insert(20).
		insert(40).
		insert(70).
		insert(60).
		insert(80)
	/*println(tree.root.Val)
	x := tree.root
	printPreOrder(x)
	println()
	LevelWiseTraversal(tree.root)
	println()
	tree.root = remove(tree.root, 20)
	println()
	inOrderTraverse(tree.root)*/
	//levelOrder(tree.root)
	//println(findHeight(tree.root))
	levelOrderAndReturnResultInArray(tree.root)
}

/*tree operation*/
func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{Val: data, left: nil, right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

func (head *Node) insert(data int) {
	if head == nil {
		return
	} else if data <= head.Val {
		if head.left == nil {
			head.left = &Node{Val: data, left: nil, right: nil}
		} else {
			head.left.insert(data)
		}
	} else {
		if head.right == nil {
			head.right = &Node{Val: data, left: nil, right: nil}
		} else {
			head.right.insert(data)
		}
	}
}

/*
we visit all nodes by following the smallest link until we find the left-most leaf,
then processes the leaf and moves to other nodes by going into the next smallest key value linked to the current one.
In the above example: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11
print left then head then Right (recursive)
*/
// internal recursive function to traverse in order
func inOrderTraverse(n *Node) {
	if n != nil {
		inOrderTraverse(n.left)
		fmt.Printf("%d ", n.Val)
		inOrderTraverse(n.right)
	}
}

/*
before visiting the children, this approach visits the node first.
In the above example: 8 -> 4 -> 2 -> 1 -> 3 -> 6 -> 5 -> 7 -> 10 -> 9 -> 11
print head then left then right (recursive)
*/
func printPreOrder(n *Node) {
	if n == nil {
		return
	} else {
		fmt.Printf("%d ", n.Val)
		printPreOrder(n.left)
		printPreOrder(n.right)
	}
}

/*
we find the smallest leaf (the left-most one), then processes the sibling
and then the parent node, then goes to the next subtree,
and navigates up the parent.
1 -> 3 -> 2 -> 5 -> 7 -> 6 -> 4 -> 9 -> 11 -> 10 -> 8
print left then right then head (recursive)
*/
func printPostOrder(n *Node) {
	if n == nil {
		return
	} else {
		printPostOrder(n.left)
		printPostOrder(n.right)
		fmt.Printf("%d ", n.Val)
	}
}

// Min returns the Item with min value stored in the tree
func (t *BinaryTree) Min() int {
	n := t.root
	if n == nil {
		return math.MinInt
	}
	for {
		if n.left == nil {
			return n.Val
		}
		n = n.left
	}
}

// Max returns the Item with max value stored in the tree
func (t *BinaryTree) Max() int {

	n := t.root
	if n == nil {
		return math.MaxInt
	}
	for {
		if n.right == nil {
			return n.Val
		}
		n = n.right
	}
}

// Search returns true if the Item t exists in the tree
func (t *BinaryTree) Search(Val int) bool {
	return search(t.root, Val)
}

// internal recursive function to search an item in the tree
func search(n *Node, Val int) bool {
	if n == nil {
		return false
	}
	if Val < n.Val {
		return search(n.left, Val)
	}
	if Val > n.Val {
		return search(n.right, Val)
	}
	return true
}

//

// Remove removes the Item with key `key` from the tree

//Delete node from binary search tree
func Delete(node *Node, val int) *Node {
	if nil == node {
		return node
	}
	if node.Val > val {
		node.left = Delete(node.left, val)
	}
	if node.Val < val {
		node.right = Delete(node.right, val)
	}
	if node.Val == val {
		if node.left == nil && node.right == nil {
			node = nil
			return node
		}
		if node.left == nil && node.right != nil {
			temp := node.right
			node = nil
			node = temp
			return node
		}
		if node.left != nil && node.right == nil {
			temp := node.left
			node = nil
			node = temp
			return node
		}

		tempNode := minValued(node.right)
		node.Val = tempNode.Val
		node.right = Delete(node.right, tempNode.Val)
	}
	return node
}
func minValued(root *Node) *Node {
	temp := root
	for nil != temp && temp.left != nil {
		temp = temp.left
	}
	return temp
}
func (t *BinaryTree) Remove(Val int) {
	remove(t.root, Val)
}

// internal recursive function to remove an item
func remove(node *Node, Val int) *Node {
	if node == nil {
		return nil
	}
	if Val < node.Val {
		node.left = remove(node.left, Val)
		return node
	}
	if Val > node.Val {
		node.right = remove(node.right, Val)
		return node
	}
	// key == node.key
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}
	if node.left == nil {
		node = node.right
		return node
	}
	if node.right == nil {
		node = node.left
		return node
	}
	leftmostrightside := node.right
	for {
		//find the smallest value on the right side
		if leftmostrightside != nil && leftmostrightside.left != nil {
			leftmostrightside = leftmostrightside.left
		} else {
			break
		}
	}
	node.Val = leftmostrightside.Val
	node.right = remove(node.right, node.Val)
	return node
}

//Find findHeight of binary tree
func findHeight(root *Node) int {
	if root == nil {
		return 0
	}
	return findMax(findHeight(root.left), findHeight(root.right)) + 1
}
func findMax(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}

// TRAVERS Tree BY LEVEL

func LevelWiseTraversal(head *Node) {
	nodeList := make([]*Node, 0)
	if nil == head {
		return
	}
	temp := &Node{}
	nodeList = append(nodeList, head)
	for len(nodeList) > 0 {
		temp, nodeList = nodeList[0], nodeList[1:]
		if nil != temp.left {
			//enqueue in the list
			nodeList = append(nodeList, temp.left)
		}
		if nil != temp.right {
			nodeList = append(nodeList, temp.right)
		}
		fmt.Printf("%d\t", temp.Val)
		time.Sleep(1 * time.Second)
	}
	fmt.Println()
}

//Binary Tree Level Order Traversal
/*
Given the root of a binary tree, return the level order traversal of its nodes' values.
(i.e., from left to right, level by level).
link:https://leetcode.com/problems/binary-tree-level-order-traversal/
*/

func enqueue(queue []*Node, element *Node) []*Node {
	queue = append(queue, element) // Simply append to enqueue.
	return queue
}

func dequeue(queue []*Node) ([]*Node, *Node) {
	element := queue[0]      // The first element is the one to be dequeued.
	var newQueue = queue[1:] // Slice off the element once it is dequeued.
	return newQueue, element
}
func isQueueEmpty(queue []*Node) bool {
	if len(queue) >= 1 {
		return false
	}
	return true
}
func levelOrder(head *Node) {
	var queue = make([]*Node, 0)
	queue = enqueue(queue, head)
	var currentNode *Node = nil
	for isQueueEmpty(queue) == false {
		queue, currentNode = dequeue(queue)
		processCurrentNode(currentNode)
		if currentNode.left != nil {
			queue = enqueue(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = enqueue(queue, currentNode.right)
		}
	}
}
func processCurrentNode(node *Node) {
	println(node.Val)
}
func levelOrderAndReturnResultInArray(head *Node) {
	var queue = make([]*Node, 0)
	queue = enqueue(queue, head)
	var currentNode *Node = nil
	levels := findHeight(head)
	arr := make([][]int, levels)
	for i := range arr {
		x := i
		println(x)

		arr[i] = make([]int, 2)
	}
	arr[0][0] = head.Val
	currentLevel := 1
	for isQueueEmpty(queue) == false {
		queue, currentNode = dequeue(queue)
		processCurrentNode(currentNode)
		if currentNode.left != nil {
			queue = enqueue(queue, currentNode.left)
			arr[currentLevel][0] = currentNode.left.Val
		}
		if currentNode.right != nil {
			queue = enqueue(queue, currentNode.right)
			arr[currentLevel][1] = currentNode.right.Val
		}
	}
	for _, row := range arr {
		for _, val := range row {
			fmt.Println(val)
		}
	}
}
