package main

//go语言实现二叉排序树及其前序遍历
import (
	"fmt"
	"math/rand"
)

type AVL struct{
	left,right *AVL
	value int
}

func (a *AVL)getLeft()(*AVL)  {
	if a.left != nil{
		return a.left
	}
	return nil
}


func (a *AVL)getRight()(*AVL)  {
	if a.right != nil{
		return a.right
	}
	return nil
}


func NewAVLNode(value int,left,right *AVL) *AVL {
	node := &AVL{value:value,left:left,right:right}
	return node
}

func (root *AVL)Find(value int) (*AVL) {
	var result *AVL
	for true{
		if value > root.value {
			root = root.right
			if root == nil{
				fmt.Println("不存在")
				break
			}
		}else if value < root.value{
			root = root.left
			if root == nil{
				fmt.Println("不存在")
				break
			}
		}else {
			result = root
			return result
		}
	}
	return result
}


func (root *AVL)Insert(a *AVL)  {
	for true{
		if a.value > root.value {
			if root.right == nil{
				root.right = a
				return
			}
			root = root.right

		}else if a.value < root.value{
			if root.left == nil{
				root.left = a
				//root = root1
				return
			}
			root = root.left
		}else {
			fmt.Println("元素已存在")
			return
		}
	}
}

func (a *AVL)FindPrint(i int)  {
	result := a.Find(i)
	if result != nil{
		fmt.Println(result.value)
	}

}

//先序遍历
func (a *AVL)LeftPrint()  {
	if a.left != nil {
		l := a.left
		l.LeftPrint()
	}
	fmt.Print("  ")
	fmt.Print(a.value)
	if a.right != nil {
		r := a.right
		r.LeftPrint()
	}
}


//先序遍历
func LeftPrint(node *AVL)  {
	if node != nil {
		LeftPrint(node.left)
		fmt.Print("  ")
		fmt.Print(node.value)
		LeftPrint(node.right)
	}
}


func (root *AVL)FindBefore(value int) (*AVL) {
	var result *AVL
	if value == root.value{
		fmt.Println("不能移除根节点")
		return nil
	}
	for true{
		if (root.right != nil && value == root.right.value )||(root.left != nil && value == root.left.value) {
			return root
		} else if value > root.value {
			root = root.right
			if root == nil{
				fmt.Println("不存在")
				break
			}
		}else if value < root.value{
			root = root.left
			if root == nil{
				fmt.Println("不存在")
				break
			}
		}else if root.left == nil && root .right == nil {
			fmt.Println("不存在")
			break
		}
	}
	return result
}

//删除元素

func (a *AVL)Remove1(v int)  {
	rm := a.Find(v)
	be := a.FindBefore(v)
	if rm != nil && be != nil {
		//要删除的元素的左右节点都为空
		if rm.left == nil && rm.right == nil{
			if rm.value > be.value{
				be.right = nil
			}else {
				be.left = nil
			}
		}else if rm.left == nil && rm.right != nil{
			//左节点为空 右节点不为空
			if rm.value > be.value{
				be.right = rm.right
			}else {
				be.left = rm.right
			}

		}else if rm.left != nil && rm.right == nil{
			//左节点不为空 右节点为空
			if rm.value > be.value{
				be.right = rm.left
			}else {
				be.left = rm.left
			}


		}else {
			//左节点不为空 右节点不为空

			//把前驱节点的指针断开
			if rm.value > be.value{
				be.right = nil
			}else {
				be.left = nil
			}
			//把rm.left及其下面的节点 以 rm.left 为代表作为一个整体重新插入
			a.Insert(rm.left)
			a.Insert(rm.right)
		}
	}

}



func main()  {
	//var root *AVL
	var root  = &AVL{nil,nil,50}
	//rand.Seed(time.Now().UnixNano())
	for i := 0;i<20;i++{
		v := rand.Intn(100)
		fmt.Println(v)
		node := NewAVLNode(v,nil,nil)
		root.Insert(node)
	}
	root.LeftPrint()
	fmt.Println("")
	LeftPrint(root)
	root.Remove1(47)
	fmt.Println("")
	root.LeftPrint()
	fmt.Println("")
	LeftPrint(root)

}
