package swapNode

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n) // read the number of the node
	// only 大写 开头函数可以被外部类访问 go 语言中

	var nums []int
	
	head := buildList(nums)
	newHead := swapPairs(head)
	printList(newHead)
}