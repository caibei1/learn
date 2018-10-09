package main
import "fmt"


//go语言实现二叉堆


//上浮  最小堆
func up(arr []int,childIndex int)  {
	//childIndex := len(arr)-1
	parentIndex := childIndex/2

	for childIndex > 0 && arr[childIndex] < arr[parentIndex]{

		//交换
		temp := arr[parentIndex]
		arr[parentIndex] = arr[childIndex]
		arr[childIndex] = temp

		childIndex = parentIndex
		parentIndex = parentIndex/2
	}
}



//下沉  最小堆
func down(arr []int,parentIndex,length int)  {

	childIndex := 2*parentIndex + 1

	//当子节点的索引小于数组长度，说明存在子节点，当下沉时，只要存在子节点，就还需要判断他是否需要继续下沉
	for childIndex < length{
		// 如果有右孩子，且右孩子小于左孩子的值，则定位到右孩子
		if childIndex + 1 < length && arr[childIndex+1] < arr[childIndex]{
			childIndex++
		}
		// 如果父节点小于任何一个孩子的值，直接跳出
		if arr[parentIndex] <= arr[childIndex]{
			break
		}

		//交换父节点和子节点的值
		temp := arr[parentIndex]
		arr[parentIndex] = arr[childIndex]
		arr[childIndex] = temp

		//从子节点开始，继续下沉
		parentIndex = childIndex
		childIndex = 2 * childIndex + 1
	}

}

func downbuild(arr []int)  {

	// 从最后一个非叶子节点开始，依次下沉调整
	for i := len(arr)/2;i >= 0;i--{
		down(arr,i,len(arr))
	}
}

func upbuild(arr []int)  {
	//除根节点，每个节点都上浮
	for i := len(arr)-1;i > 1; i--{
		up(arr,i)
	}
}


//验证是否是最小堆
func ifMinHeap(arr []int)bool  {
	for parentIndex := len(arr)/2;parentIndex > 0;parentIndex--{
		if (2*parentIndex+1 < len(arr) && arr[parentIndex] > arr[2*parentIndex+1]) || (2*parentIndex+2 < len(arr) &&arr[parentIndex] > arr[2*parentIndex+2]){
			fmt.Println(parentIndex)
			return false
		}
	}

	return true
}

func main()  {

	var array1 = []int{3,5,4,22,9,7,13,11,0,6}
	downbuild(array1)
	fmt.Println(ifMinHeap(array1))
	fmt.Println(array1)

	var array2 = []int{3,5,4,8,9,7,1,11,0,6}
	upbuild(array2)
	fmt.Println(ifMinHeap(array2))
	fmt.Println(array2)

}

